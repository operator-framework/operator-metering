package operator

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/rob***REMOVED***g/cron"
	log "github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	cbTypes "github.com/operator-framework/operator-metering/pkg/apis/metering/v1alpha1"
	cbutil "github.com/operator-framework/operator-metering/pkg/apis/metering/v1alpha1/util"
	"github.com/operator-framework/operator-metering/pkg/hive"
	"github.com/operator-framework/operator-metering/pkg/operator/reporting"
	"github.com/operator-framework/operator-metering/pkg/operator/reportingutil"
	"github.com/operator-framework/operator-metering/pkg/util/slice"
)

const (
	reportFinalizer = cbTypes.GroupName + "/report"
)

var (
	reportPrometheusMetricLabels = []string{"report", "namespace", "reportgenerationquery", "table_name"}

	generateReportTotalCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: prometheusMetricNamespace,
			Name:      "generate_reports_total",
			Help:      "Duration to generate a Report.",
		},
		reportPrometheusMetricLabels,
	)

	generateReportFailedCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: prometheusMetricNamespace,
			Name:      "generate_reports_failed_total",
			Help:      "Duration to generate a Report.",
		},
		reportPrometheusMetricLabels,
	)

	generateReportDurationHistogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: prometheusMetricNamespace,
			Name:      "generate_report_duration_seconds",
			Help:      "Duration to generate a Report.",
			Buckets:   []float64{60.0, 300.0, 600.0},
		},
		reportPrometheusMetricLabels,
	)
)

func init() {
	prometheus.MustRegister(generateReportFailedCounter)
	prometheus.MustRegister(generateReportTotalCounter)
	prometheus.MustRegister(generateReportDurationHistogram)
}

func (op *Reporting) runReportWorker() {
	logger := op.logger.WithField("component", "reportWorker")
	logger.Infof("Report worker started")
	const maxRequeues = 5
	for op.processResource(logger, op.syncReport, "Report", op.reportQueue, maxRequeues) {
	}
}

func (op *Reporting) syncReport(logger log.FieldLogger, key string) error {
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		logger.WithError(err).Errorf("invalid resource key :%s", key)
		return nil
	}

	logger = logger.WithFields(log.Fields{"report": name, "namespace": namespace})
	report, err := op.reportLister.Reports(namespace).Get(name)
	if err != nil {
		if apierrors.IsNotFound(err) {
			logger.Infof("report %s/%s does not exist anymore, stopping and removing any running jobs for Report", namespace, name)
			return nil
		}
		return err
	}
	sr := report.DeepCopy()

	if report.DeletionTimestamp != nil {
		_, err = op.removeReportFinalizer(sr)
		return err
	}

	return op.handleReport(logger, sr)
}

type reportSchedule interface {
	// Return the next activation time, later than the given time.
	// Next is invoked initially, and then each time the job runs..
	Next(time.Time) time.Time
}

func getSchedule(reportSched *cbTypes.ReportSchedule) (reportSchedule, error) {
	var cronSpec string
	switch reportSched.Period {
	case cbTypes.ReportPeriodCron:
		if reportSched.Cron == nil || reportSched.Cron.Expression == "" {
			return nil, fmt.Errorf("spec.schedule.cron.expression must be speci***REMOVED***ed")
		}
		return cron.ParseStandard(reportSched.Cron.Expression)
	case cbTypes.ReportPeriodHourly:
		sched := reportSched.Hourly
		if sched == nil {
			sched = &cbTypes.ReportScheduleHourly{}
		}
		if err := validateMinute(sched.Minute); err != nil {
			return nil, err
		}
		if err := validateSecond(sched.Second); err != nil {
			return nil, err
		}
		cronSpec = fmt.Sprintf("%d %d * * * *", sched.Second, sched.Minute)
	case cbTypes.ReportPeriodDaily:
		sched := reportSched.Daily
		if sched == nil {
			sched = &cbTypes.ReportScheduleDaily{}
		}
		if err := validateHour(sched.Hour); err != nil {
			return nil, err
		}
		if err := validateMinute(sched.Minute); err != nil {
			return nil, err
		}
		if err := validateSecond(sched.Second); err != nil {
			return nil, err
		}
		cronSpec = fmt.Sprintf("%d %d %d * * *", sched.Second, sched.Minute, sched.Hour)
	case cbTypes.ReportPeriodWeekly:
		sched := reportSched.Weekly
		if sched == nil {
			sched = &cbTypes.ReportScheduleWeekly{}
		}
		dow := 0
		if sched.DayOfWeek != nil {
			var err error
			dow, err = convertDayOfWeek(*sched.DayOfWeek)
			if err != nil {
				return nil, err
			}
		}
		if err := validateHour(sched.Hour); err != nil {
			return nil, err
		}
		if err := validateMinute(sched.Minute); err != nil {
			return nil, err
		}
		if err := validateSecond(sched.Second); err != nil {
			return nil, err
		}
		cronSpec = fmt.Sprintf("%d %d %d * * %d", sched.Second, sched.Minute, sched.Hour, dow)
	case cbTypes.ReportPeriodMonthly:
		sched := reportSched.Monthly
		if sched == nil {
			sched = &cbTypes.ReportScheduleMonthly{}
		}
		dom := int64(1)
		if sched.DayOfMonth != nil {
			dom = *sched.DayOfMonth
		}
		if err := validateDayOfMonth(dom); err != nil {
			return nil, err
		}
		if err := validateHour(sched.Hour); err != nil {
			return nil, err
		}
		if err := validateMinute(sched.Minute); err != nil {
			return nil, err
		}
		if err := validateSecond(sched.Second); err != nil {
			return nil, err
		}
		cronSpec = fmt.Sprintf("%d %d %d %d * *", sched.Second, sched.Minute, sched.Hour, dom)
	default:
		return nil, fmt.Errorf("invalid Report.spec.schedule.period: %s", reportSched.Period)
	}
	return cron.Parse(cronSpec)
}

func (op *Reporting) handleReport(logger log.FieldLogger, report *cbTypes.Report) error {
	if op.cfg.EnableFinalizers && reportNeedsFinalizer(report) {
		var err error
		report, err = op.addReportFinalizer(report)
		if err != nil {
			return err
		}
	}

	return op.runReport(logger, report)
}

type reportPeriod struct {
	periodEnd   time.Time
	periodStart time.Time
}

// isReportFinished checks the running condition of the report parameter and returns true if the report has previously run
func isReportFinished(logger log.FieldLogger, report *cbTypes.Report) bool {
	// check if this report was previously ***REMOVED***nished
	runningCond := cbutil.GetReportCondition(report.Status, cbTypes.ReportRunning)

	if runningCond == nil {
		logger.Infof("new report, validating report")
	} ***REMOVED*** if runningCond.Reason == cbutil.ReportFinishedReason && runningCond.Status != v1.ConditionTrue {
		// Found an already ***REMOVED***nished runOnce report. Log that we're not
		// re-processing runOnce reports after they're previously ***REMOVED***nished
		if report.Spec.Schedule == nil {
			logger.Infof("Report %s is a previously ***REMOVED***nished run-once report, not re-processing", report.Name)
			return true
		}
		// log some messages to indicate we're processing what was a previously ***REMOVED***nished report

		// if the report's reportingEnd is unset or after the lastReportTime
		// then the report was updated since it last ***REMOVED***nished and we should
		// consider it something to be reprocessed
		if report.Spec.ReportingEnd == nil {
			logger.Infof("previously ***REMOVED***nished report's spec.reportingEnd is unset: beginning processing of report")
		} ***REMOVED*** if report.Status.LastReportTime != nil && report.Spec.ReportingEnd.Time.After(report.Status.LastReportTime.Time) {
			logger.Infof("previously ***REMOVED***nished report's spec.reportingEnd (%s) is now after lastReportTime (%s): beginning processing of report", report.Spec.ReportingEnd.Time, report.Status.LastReportTime.Time)
		} ***REMOVED*** {
			// return without processing because the report is complete
			logger.Infof("Report %s is already ***REMOVED***nished: %s", report.Name, runningCond.Message)
			return true
		}
	}

	return false
}

// validateReport takes a Report structure and checks if it contains valid ***REMOVED***elds
func validateReport(
	report *cbTypes.Report,
	queryGetter reporting.ReportGenerationQueryGetter,
	depResolver DependencyResolver,
	handler *reporting.UninitialiedDependendenciesHandler,
) (*cbTypes.ReportGenerationQuery, *reporting.DependencyResolutionResult, error) {
	// Validate the ReportGenerationQuery is set
	if report.Spec.GenerationQueryName == "" {
		return nil, nil, errors.New("must set spec.generationQuery")
	}

	// Validate the reportingStart and reportingEnd make sense and are set when
	// required
	if report.Spec.ReportingStart != nil && report.Spec.ReportingEnd != nil && (report.Spec.ReportingStart.Time.After(report.Spec.ReportingEnd.Time) || report.Spec.ReportingStart.Time.Equal(report.Spec.ReportingEnd.Time)) {
		return nil, nil, fmt.Errorf("spec.reportingEnd (%s) must be after spec.reportingStart (%s)", report.Spec.ReportingEnd.Time, report.Spec.ReportingStart.Time)
	}
	if report.Spec.ReportingEnd == nil && report.Spec.RunImmediately {
		return nil, nil, errors.New("spec.reportingEnd must be set if report.spec.runImmediately is true")
	}

	// Validate the ReportGenerationQuery that the Report used exists
	genQuery, err := GetReportGenerationQueryForReport(report, queryGetter)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return nil, nil, fmt.Errorf("ReportGenerationQuery (%s) does not exist", report.Spec.GenerationQueryName)
		}
		return nil, nil, fmt.Errorf("failed to get report generation query")
	}

	// Validate the dependencies of this Report's query exist
	dependencyResult, err := depResolver.ResolveDependencies(
		genQuery.Namespace,
		genQuery.Spec.Inputs,
		report.Spec.Inputs)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to resolve ReportGenerationQuery dependencies %s: %v", genQuery.Name, err)
	}
	err = reporting.ValidateGenerationQueryDependencies(dependencyResult.Dependencies, handler)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to validate ReportGenerationQuery dependencies %s: %v", genQuery.Name, err)
	}

	return genQuery, dependencyResult, nil
}

// getReportPeriod determines a Report's reporting period based off the report parameter's ***REMOVED***elds.
// Returns a pointer to a reportPeriod structure if no error was encountered, ***REMOVED*** panic or return an error.
func getReportPeriod(now time.Time, logger log.FieldLogger, report *cbTypes.Report) (*reportPeriod, error) {
	var reportPeriod *reportPeriod

	// check if the report's schedule spec is set
	if report.Spec.Schedule != nil {
		reportSchedule, err := getSchedule(report.Spec.Schedule)
		if err != nil {
			return nil, err
		}

		if report.Status.LastReportTime != nil {
			reportPeriod = getNextReportPeriod(reportSchedule, report.Spec.Schedule.Period, report.Status.LastReportTime.Time)
		} ***REMOVED*** {
			if report.Spec.ReportingStart != nil {
				logger.Infof("no last report time for report, using spec.reportingStart %s as starting point", report.Spec.ReportingStart.Time)
				reportPeriod = getNextReportPeriod(reportSchedule, report.Spec.Schedule.Period, report.Spec.ReportingStart.Time)
			} ***REMOVED*** if report.Status.NextReportTime != nil {
				logger.Infof("no last report time for report, using status.nextReportTime %s as starting point", report.Status.NextReportTime.Time)
				reportPeriod = getNextReportPeriod(reportSchedule, report.Spec.Schedule.Period, report.Status.NextReportTime.Time)
			} ***REMOVED*** {
				// the current period, [now, nextScheduledTime]
				currentPeriod := getNextReportPeriod(reportSchedule, report.Spec.Schedule.Period, now)
				// the next full report period from [nextScheduledTime, nextScheduledTime+1]
				reportPeriod = getNextReportPeriod(reportSchedule, report.Spec.Schedule.Period, currentPeriod.periodEnd)
				report.Status.NextReportTime = &metav1.Time{Time: reportPeriod.periodStart}
			}
		}
	} ***REMOVED*** {
		var err error
		// if there's the Spec.Schedule ***REMOVED***eld is unset, then the report must be a run-once report
		reportPeriod, err = getRunOnceReportPeriod(report)
		if err != nil {
			return nil, err
		}
	}

	if reportPeriod.periodStart.After(reportPeriod.periodEnd) {
		panic("periodStart should never come after periodEnd")
	}

	if report.Spec.ReportingEnd != nil && reportPeriod.periodEnd.After(report.Spec.ReportingEnd.Time) {
		logger.Debugf("calculated Report periodEnd %s goes beyond spec.reportingEnd %s, setting periodEnd to reportingEnd", reportPeriod.periodEnd, report.Spec.ReportingEnd.Time)
		// we need to truncate the reportPeriod to align with the reportingEnd
		reportPeriod.periodEnd = report.Spec.ReportingEnd.Time
	}

	return reportPeriod, nil
}

// runReport takes a report, and generates reporting data
// according the report's schedule. If the next scheduled reporting period
// hasn't elapsed, runReport will requeue the resource for a time when
// the period has elapsed.
func (op *Reporting) runReport(logger log.FieldLogger, report *cbTypes.Report) error {
	// check if the report was previously ***REMOVED***nished; store result in bool
	if reportFinished := isReportFinished(logger, report); reportFinished {
		return nil
	}

	runningCond := cbutil.GetReportCondition(report.Status, cbTypes.ReportRunning)
	queryGetter := reporting.NewReportGenerationQueryListerGetter(op.reportGenerationQueryLister)

	// validate that Report contains valid Spec ***REMOVED***elds
	genQuery, dependencyResult, err := validateReport(report, queryGetter, op.dependencyResolver, op.uninitialiedDependendenciesHandler())
	if err != nil {
		return op.setReportStatusInvalidReport(report, err.Error())
	}

	now := op.clock.Now().UTC()

	// get the report's reporting period
	reportPeriod, err := getReportPeriod(now, logger, report)
	if err != nil {
		return err
	}

	logger = logger.WithFields(log.Fields{
		"periodStart":       reportPeriod.periodStart,
		"periodEnd":         reportPeriod.periodEnd,
		"overwriteExisting": report.Spec.OverwriteExistingData,
	})

	// create the table before we check to see if the report has dependencies
	// that are missing data
	var prestoTable *cbTypes.PrestoTable
	// if tableName isn't set, this report is still new and we should make sure
	// no tables exist already in case of a previously failed cleanup.
	if report.Status.TableRef.Name != "" {
		prestoTable, err = op.prestoTableLister.PrestoTables(report.Namespace).Get(report.Status.TableRef.Name)
		if err != nil {
			return fmt.Errorf("unable to get PrestoTable %s for Report %s, %s", report.Status.TableRef, report.Name, err)
		}
		logger.Infof("Report %s table already exists, tableName: %s", report.Name, prestoTable.Status.TableName)
	} ***REMOVED*** {
		tableName := reportingutil.ReportTableName(report.Namespace, report.Name)
		hiveStorage, err := op.getHiveStorage(report.Spec.Output, report.Namespace)
		if err != nil {
			return fmt.Errorf("storage incorrectly con***REMOVED***gured for Report %s, err: %v", report.Name, err)
		}
		if hiveStorage.Status.Hive.DatabaseName == "" {
			return fmt.Errorf("StorageLocation %s Hive database %s does not exist yet", hiveStorage.Name, hiveStorage.Spec.Hive.DatabaseName)
		}

		cols, err := reportingutil.PrestoColumnsToHiveColumns(reportingutil.GeneratePrestoColumns(genQuery))
		if err != nil {
			return fmt.Errorf("unable to convert Presto columns to Hive columns: %s", err)
		}

		params := hive.TableParameters{
			Database: hiveStorage.Status.Hive.DatabaseName,
			Name:     tableName,
			Columns:  cols,
		}
		if hiveStorage.Spec.Hive.DefaultTableProperties != nil {
			params.SerdeFormat = hiveStorage.Spec.Hive.DefaultTableProperties.SerdeFormat
			params.FileFormat = hiveStorage.Spec.Hive.DefaultTableProperties.FileFormat
			params.SerdeRowProperties = hiveStorage.Spec.Hive.DefaultTableProperties.SerdeRowProperties
		}

		logger.Infof("creating table %s", tableName)
		hiveTable, err := op.createHiveTableCR(report, cbTypes.ReportGVK, params, false, nil)
		if err != nil {
			return fmt.Errorf("error creating table for Report %s: %s", report.Name, err)
		}
		hiveTable, err = op.waitForHiveTable(hiveTable.Namespace, hiveTable.Name, time.Second, 20*time.Second)
		if err != nil {
			return fmt.Errorf("error creating table for Report %s: %s", report.Name, err)
		}
		prestoTable, err = op.waitForPrestoTable(hiveTable.Namespace, hiveTable.Name, time.Second, 20*time.Second)
		if err != nil {
			return fmt.Errorf("error creating table for Report %s: %s", report.Name, err)
		}

		logger.Infof("created table %s", tableName)
		dataSourceName := fmt.Sprintf("report-%s", report.Name)
		logger.Infof("creating PrestoTable ReportDataSource %s pointing at report table %s", dataSourceName, prestoTable.Status.TableName)
		ownerRef := metav1.NewControllerRef(prestoTable, cbTypes.PrestoTableGVK)
		newReportDataSource := &cbTypes.ReportDataSource{
			TypeMeta: metav1.TypeMeta{
				Kind:       "ReportDataSource",
				APIVersion: cbTypes.ReportDataSourceGVK.GroupVersion().String(),
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      dataSourceName,
				Namespace: prestoTable.Namespace,
				Labels:    prestoTable.ObjectMeta.Labels,
				OwnerReferences: []metav1.OwnerReference{
					*ownerRef,
				},
			},
			Spec: cbTypes.ReportDataSourceSpec{
				PrestoTable: &cbTypes.PrestoTableDataSource{
					TableRef: v1.LocalObjectReference{
						Name: prestoTable.Name,
					},
				},
			},
		}
		_, err = op.meteringClient.MeteringV1alpha1().ReportDataSources(report.Namespace).Create(newReportDataSource)
		if err != nil {
			if apierrors.IsAlreadyExists(err) {
				logger.Infof("ReportDataSource %s already exists", dataSourceName)
			} ***REMOVED*** {
				return fmt.Errorf("error creating PrestoTable ReportDataSource %s: %s", dataSourceName, err)
			}
		}
		logger.Infof("created PrestoTable ReportDataSource %s", dataSourceName)

		report.Status.TableRef = v1.LocalObjectReference{Name: hiveTable.Name}
		report, err = op.meteringClient.MeteringV1alpha1().Reports(report.Namespace).Update(report)
		if err != nil {
			logger.WithError(err).Errorf("unable to update Report status with tableName")
			return err
		}
	}

	var runningMsg, runningReason string
	if report.Spec.RunImmediately {
		runningReason = cbutil.RunImmediatelyReason
		runningMsg = fmt.Sprintf("Report %s scheduled: runImmediately=true bypassing reporting period [%s to %s].", report.Name, reportPeriod.periodStart, reportPeriod.periodEnd)
	} ***REMOVED*** {
		// Check if it's time to generate the report
		if reportPeriod.periodEnd.After(now) {
			waitTime := reportPeriod.periodEnd.Sub(now)
			waitMsg := fmt.Sprintf("Next scheduled report period is [%s to %s]. next run time is %s.", reportPeriod.periodStart, reportPeriod.periodEnd, reportPeriod.periodEnd)
			logger.Infof(waitMsg+". waiting %s", waitTime)

			if runningCond := cbutil.GetReportCondition(report.Status, cbTypes.ReportRunning); runningCond != nil && runningCond.Status == v1.ConditionTrue && runningCond.Reason == cbutil.ReportingPeriodWaitingReason {
				op.enqueueReportAfter(report, waitTime)
				return nil
			}

			var err error
			report, err = op.updateReportStatus(report, cbutil.NewReportCondition(cbTypes.ReportRunning, v1.ConditionFalse, cbutil.ReportingPeriodWaitingReason, waitMsg))
			if err != nil {
				return err
			}

			// we requeue this for later when the period we need to report on next
			// has elapsed
			op.enqueueReportAfter(report, waitTime)
			return nil
		}

		runningReason = cbutil.ScheduledReason
		runningMsg = fmt.Sprintf("Report %s scheduled: reached end of reporting period [%s to %s].", report.Name, reportPeriod.periodStart, reportPeriod.periodEnd)

		var unmetDataStartDataSourceDependendencies, unmetDataEndDataSourceDependendencies, unstartedDataSourceDependencies []string
		// Validate all ReportDataSources that the Report depends on have indicated
		// they have data available that covers the current reportPeriod.
		for _, dataSource := range dependencyResult.Dependencies.ReportDataSources {
			if dataSource.Spec.Promsum != nil {
				// queue the dataSource and store the list of reports so we can
				// add information to the Report's status on what's currently
				// not ready
				queue := false
				if dataSource.Status.PrometheusMetricImportStatus == nil {
					unstartedDataSourceDependencies = append(unmetDataStartDataSourceDependendencies, dataSource.Name)
					queue = true
				} ***REMOVED*** {
					// reportPeriod lower bound not covered
					if dataSource.Status.PrometheusMetricImportStatus.ImportDataStartTime == nil || reportPeriod.periodStart.Before(dataSource.Status.PrometheusMetricImportStatus.ImportDataStartTime.Time) {
						queue = true
						unmetDataStartDataSourceDependendencies = append(unmetDataStartDataSourceDependendencies, dataSource.Name)
					}
					// reportPeriod upper bound is not covered
					if dataSource.Status.PrometheusMetricImportStatus.ImportDataEndTime == nil || reportPeriod.periodEnd.After(dataSource.Status.PrometheusMetricImportStatus.ImportDataEndTime.Time) {
						queue = true
						unmetDataEndDataSourceDependendencies = append(unmetDataEndDataSourceDependendencies, dataSource.Name)
					}
				}
				if queue {
					op.enqueueReportDataSource(dataSource)
				}
			}
		}

		// Validate all sub-reports that the Report depends on have reported on the
		// current reportPeriod
		var unmetReportDependendencies []string
		for _, subReport := range dependencyResult.Dependencies.Reports {
			if subReport.Status.LastReportTime != nil && subReport.Status.LastReportTime.Time.Before(reportPeriod.periodEnd) {
				op.enqueueReport(subReport)
				unmetReportDependendencies = append(unmetReportDependendencies, subReport.Name)
			}
		}

		if len(unstartedDataSourceDependencies) != 0 || len(unmetDataStartDataSourceDependendencies) != 0 || len(unmetDataEndDataSourceDependendencies) != 0 || len(unmetReportDependendencies) != 0 {
			unmetMsg := "The following Report dependencies do not have data currently available for the current reportPeriod being processed:"
			if len(unstartedDataSourceDependencies) != 0 || len(unmetDataStartDataSourceDependendencies) != 0 || len(unmetDataEndDataSourceDependendencies) != 0 {
				var msgs []string
				if len(unstartedDataSourceDependencies) != 0 {
					// sort so the message is reproducible
					sort.Strings(unstartedDataSourceDependencies)
					msgs = append(msgs, fmt.Sprintf("no data: [%s]", strings.Join(unstartedDataSourceDependencies, ", ")))
				}
				if len(unmetDataStartDataSourceDependendencies) != 0 {
					// sort so the message is reproducible
					sort.Strings(unmetDataStartDataSourceDependendencies)
					msgs = append(msgs, fmt.Sprintf("periodStart %s is before importDataStartTime of [%s]", reportPeriod.periodStart, strings.Join(unmetDataStartDataSourceDependendencies, ", ")))
				}
				if len(unmetDataEndDataSourceDependendencies) != 0 {
					// sort so the message is reproducible
					sort.Strings(unmetDataEndDataSourceDependendencies)
					msgs = append(msgs, fmt.Sprintf("periodEnd %s is after importDataEndTime of [%s]", reportPeriod.periodEnd, strings.Join(unmetDataEndDataSourceDependendencies, ", ")))
				}
				unmetMsg += fmt.Sprintf(" ReportDataSources: %s", strings.Join(msgs, ", "))
			}
			if len(unmetReportDependendencies) != 0 {
				// sort so the message is reproducible
				sort.Strings(unmetReportDependendencies)
				unmetMsg += fmt.Sprintf(" Reports: lastReportTime not prior to periodEnd %s: [%s]", reportPeriod.periodEnd, strings.Join(unmetReportDependendencies, ", "))
			}

			// If the previous condition is unmet dependencies, check if the
			// message changes, and only update if it does
			if runningCond != nil && runningCond.Status == v1.ConditionFalse && runningCond.Reason == cbutil.ReportingPeriodUnmetDependenciesReason && runningCond.Message == unmetMsg {
				logger.Debugf("Report %s already has Running condition=false with reason=%s and unchanged message, skipping update", report.Name, cbutil.ReportingPeriodUnmetDependenciesReason)
				return nil
			}
			logger.Warnf(unmetMsg)
			_, err := op.updateReportStatus(report, cbutil.NewReportCondition(cbTypes.ReportRunning, v1.ConditionFalse, cbutil.ReportingPeriodUnmetDependenciesReason, unmetMsg))
			return err
		}
	}
	logger.Infof(runningMsg + " Running now.")

	report, err = op.updateReportStatus(report, cbutil.NewReportCondition(cbTypes.ReportRunning, v1.ConditionTrue, runningReason, runningMsg))
	if err != nil {
		return err
	}

	prestoTables, err := op.prestoTableLister.PrestoTables(report.Namespace).List(labels.Everything())
	if err != nil {
		return err
	}

	reports, err := op.reportLister.Reports(report.Namespace).List(labels.Everything())
	if err != nil {
		return err
	}

	datasources, err := op.reportDataSourceLister.ReportDataSources(report.Namespace).List(labels.Everything())
	if err != nil {
		return err
	}

	queries, err := op.reportGenerationQueryLister.ReportGenerationQueries(report.Namespace).List(labels.Everything())
	if err != nil {
		return err
	}

	queryCtx := &reporting.ReportQueryTemplateContext{
		Namespace:               report.Namespace,
		ReportQuery:             genQuery,
		Reports:                 reports,
		ReportGenerationQueries: queries,
		ReportDataSources:       datasources,
		PrestoTables:            prestoTables,
	}
	tmplCtx := reporting.TemplateContext{
		Report: reporting.ReportTemplateInfo{
			ReportingStart: &reportPeriod.periodStart,
			ReportingEnd:   &reportPeriod.periodEnd,
			Inputs:         dependencyResult.InputValues,
		},
	}

	// Render the query template
	query, err := reporting.RenderQuery(queryCtx, tmplCtx)
	if err != nil {
		return err
	}

	metricLabels := prometheus.Labels{
		"report":                report.Name,
		"namespace":             report.Namespace,
		"reportgenerationquery": report.Spec.GenerationQueryName,
		"table_name":            prestoTable.Status.TableName,
	}

	genReportTotalCounter := generateReportTotalCounter.With(metricLabels)
	genReportFailedCounter := generateReportFailedCounter.With(metricLabels)
	genReportDurationObserver := generateReportDurationHistogram.With(metricLabels)

	logger.Infof("generating Report %s using query %s and periodStart: %s, periodEnd: %s", report.Name, genQuery.Name, reportPeriod.periodStart, reportPeriod.periodEnd)

	genReportTotalCounter.Inc()
	generateReportStart := op.clock.Now()
	tableName := reportingutil.FullyQuali***REMOVED***edTableName(prestoTable)
	err = op.reportGenerator.GenerateReport(tableName, query, report.Spec.OverwriteExistingData)
	generateReportDuration := op.clock.Since(generateReportStart)
	genReportDurationObserver.Observe(float64(generateReportDuration.Seconds()))
	if err != nil {
		genReportFailedCounter.Inc()
		// update the status to Failed with message containing the
		// error
		errMsg := fmt.Sprintf("error occurred while generating report: %s", err)
		_, updateErr := op.updateReportStatus(report, cbutil.NewReportCondition(cbTypes.ReportRunning, v1.ConditionFalse, cbutil.GenerateReportFailedReason, errMsg))
		if updateErr != nil {
			logger.WithError(updateErr).Errorf("unable to update Report status")
			return updateErr
		}
		return fmt.Errorf("failed to generateReport for Report %s, err: %v", report.Name, err)
	}

	logger.Infof("successfully generated Report %s using query %s and periodStart: %s, periodEnd: %s", report.Name, genQuery.Name, reportPeriod.periodStart, reportPeriod.periodEnd)

	// Update the LastReportTime on the report status
	report.Status.LastReportTime = &metav1.Time{Time: reportPeriod.periodEnd}

	// check if we've reached the con***REMOVED***gured ReportingEnd, and if so, update
	// the status to indicate the report has ***REMOVED***nished
	if report.Spec.ReportingEnd != nil && report.Status.LastReportTime.Time.Equal(report.Spec.ReportingEnd.Time) {
		msg := fmt.Sprintf("Report has ***REMOVED***nished reporting. Report has reached the con***REMOVED***gured spec.reportingEnd: %s", report.Spec.ReportingEnd.Time)
		runningCond := cbutil.NewReportCondition(cbTypes.ReportRunning, v1.ConditionFalse, cbutil.ReportFinishedReason, msg)
		cbutil.SetReportCondition(&report.Status, *runningCond)
		logger.Infof(msg)
	} ***REMOVED*** if report.Spec.Schedule != nil {
		// determine the next reportTime, if it's not a run-once report and then
		// queue the report for that time
		reportSchedule, err := getSchedule(report.Spec.Schedule)
		if err != nil {
			return err
		}

		nextReportPeriod := getNextReportPeriod(reportSchedule, report.Spec.Schedule.Period, report.Status.LastReportTime.Time)

		// update the NextReportTime on the report status
		report.Status.NextReportTime = &metav1.Time{Time: nextReportPeriod.periodEnd}

		// calculate the time to reprocess after queuing
		now = op.clock.Now().UTC()
		nextRunTime := nextReportPeriod.periodEnd
		waitTime := nextRunTime.Sub(now)

		waitMsg := fmt.Sprintf("Next scheduled report period is [%s to %s]. next run time is %s.", reportPeriod.periodStart, reportPeriod.periodEnd, nextRunTime)
		runningCond := cbutil.NewReportCondition(cbTypes.ReportRunning, v1.ConditionFalse, cbutil.ReportingPeriodWaitingReason, waitMsg)
		cbutil.SetReportCondition(&report.Status, *runningCond)
		logger.Infof(waitMsg+". waiting %s", waitTime)
		op.enqueueReportAfter(report, waitTime)
	}

	// Update the status
	report, err = op.meteringClient.MeteringV1alpha1().Reports(report.Namespace).Update(report)
	if err != nil {
		logger.WithError(err).Errorf("unable to update Report status")
		return err
	}

	if err := op.queueDependentReportGenerationQueriesForReport(report); err != nil {
		logger.WithError(err).Errorf("error queuing ReportGenerationQuery dependents of Report %s", report.Name)
	}
	if err := op.queueDependentReportsForReport(report); err != nil {
		logger.WithError(err).Errorf("error queuing Report dependents of Report %s", report.Name)
	}
	return nil
}

func getRunOnceReportPeriod(report *cbTypes.Report) (*reportPeriod, error) {
	if report.Spec.ReportingEnd == nil || report.Spec.ReportingStart == nil {
		return nil, fmt.Errorf("run-once reports must have both ReportingEnd and ReportingStart")
	}
	reportPeriod := &reportPeriod{
		periodStart: report.Spec.ReportingStart.UTC(),
		periodEnd:   report.Spec.ReportingEnd.UTC(),
	}
	return reportPeriod, nil
}

func getNextReportPeriod(schedule reportSchedule, period cbTypes.ReportPeriod, lastScheduled time.Time) *reportPeriod {
	periodStart := lastScheduled.UTC()
	periodEnd := schedule.Next(periodStart)
	return &reportPeriod{
		periodStart: periodStart.Truncate(time.Millisecond).UTC(),
		periodEnd:   periodEnd.Truncate(time.Millisecond).UTC(),
	}
}

func convertDayOfWeek(dow string) (int, error) {
	switch strings.ToLower(dow) {
	case "sun", "sunday":
		return 0, nil
	case "mon", "monday":
		return 1, nil
	case "tue", "tues", "tuesday":
		return 2, nil
	case "wed", "weds", "wednesday":
		return 3, nil
	case "thur", "thurs", "thursday":
		return 4, nil
	case "fri", "friday":
		return 5, nil
	case "sat", "saturday":
		return 6, nil
	}
	return 0, fmt.Errorf("invalid day of week: %s", dow)
}

func (op *Reporting) addReportFinalizer(report *cbTypes.Report) (*cbTypes.Report, error) {
	report.Finalizers = append(report.Finalizers, reportFinalizer)
	newReport, err := op.meteringClient.MeteringV1alpha1().Reports(report.Namespace).Update(report)
	logger := op.logger.WithFields(log.Fields{"report": report.Name, "namespace": report.Namespace})
	if err != nil {
		logger.WithError(err).Errorf("error adding %s ***REMOVED***nalizer to Report: %s/%s", reportFinalizer, report.Namespace, report.Name)
		return nil, err
	}
	logger.Infof("added %s ***REMOVED***nalizer to Report: %s/%s", reportFinalizer, report.Namespace, report.Name)
	return newReport, nil
}

func (op *Reporting) removeReportFinalizer(report *cbTypes.Report) (*cbTypes.Report, error) {
	if !slice.ContainsString(report.ObjectMeta.Finalizers, reportFinalizer, nil) {
		return report, nil
	}
	report.Finalizers = slice.RemoveString(report.Finalizers, reportFinalizer, nil)
	newReport, err := op.meteringClient.MeteringV1alpha1().Reports(report.Namespace).Update(report)
	logger := op.logger.WithFields(log.Fields{"report": report.Name, "namespace": report.Namespace})
	if err != nil {
		logger.WithError(err).Errorf("error removing %s ***REMOVED***nalizer from Report: %s/%s", reportFinalizer, report.Namespace, report.Name)
		return nil, err
	}
	logger.Infof("removed %s ***REMOVED***nalizer from Report: %s/%s", reportFinalizer, report.Namespace, report.Name)
	return newReport, nil
}

func reportNeedsFinalizer(report *cbTypes.Report) bool {
	return report.ObjectMeta.DeletionTimestamp == nil && !slice.ContainsString(report.ObjectMeta.Finalizers, reportFinalizer, nil)
}

func (op *Reporting) updateReportStatus(report *cbTypes.Report, cond *cbTypes.ReportCondition) (*cbTypes.Report, error) {
	cbutil.SetReportCondition(&report.Status, *cond)
	return op.meteringClient.MeteringV1alpha1().Reports(report.Namespace).Update(report)
}

func (op *Reporting) setReportStatusInvalidReport(report *cbTypes.Report, msg string) error {
	logger := op.logger.WithFields(log.Fields{"report": report.Name, "namespace": report.Namespace})
	// don't update unless the validation error changes
	if runningCond := cbutil.GetReportCondition(report.Status, cbTypes.ReportRunning); runningCond != nil && runningCond.Status == v1.ConditionFalse && runningCond.Reason == cbutil.InvalidReportReason && runningCond.Message == msg {
		logger.Debugf("Report %s failed validation last reconcile, skipping updating status", report.Name)
		return nil
	}

	logger.Warnf("Report %s failed validation: %s", report.Name, msg)
	cond := cbutil.NewReportCondition(cbTypes.ReportRunning, v1.ConditionFalse, cbutil.InvalidReportReason, msg)
	_, err := op.updateReportStatus(report, cond)
	return err
}

// GetReportGenerationQueryForReport returns the ReportGenerationQuery that was used in the Report parameter
func GetReportGenerationQueryForReport(report *cbTypes.Report, queryGetter reporting.ReportGenerationQueryGetter) (*cbTypes.ReportGenerationQuery, error) {
	return queryGetter.GetReportGenerationQuery(report.Namespace, report.Spec.GenerationQueryName)
}

func (op *Reporting) getReportDependencies(report *cbTypes.Report) (*reporting.ReportGenerationQueryDependencies, error) {
	return op.getGenerationQueryDependencies(report.Namespace, report.Spec.GenerationQueryName, report.Spec.Inputs)
}

func (op *Reporting) queueDependentReportsForReport(report *cbTypes.Report) error {
	// Look for all reports in the namespace
	reports, err := op.reportLister.Reports(report.Namespace).List(labels.Everything())
	if err != nil {
		return err
	}

	// for each report in the namespace, ***REMOVED***nd ones that depend on the report
	// passed into the function.
	for _, otherReport := range reports {
		deps, err := op.getReportDependencies(otherReport)
		if err != nil {
			return err
		}
		// If this otherReport has a dependency on the passed in report, queue
		// it
		for _, dep := range deps.Reports {
			if dep.Name == report.Name {
				op.enqueueReport(otherReport)
				break
			}
		}
	}
	return nil
}

// queueDependentReportGenerationQueriesForReport will queue all
// ReportGenerationQueries in the namespace which have a dependency on the
// report
func (op *Reporting) queueDependentReportGenerationQueriesForReport(report *cbTypes.Report) error {
	queryLister := op.meteringClient.MeteringV1alpha1().ReportGenerationQueries(report.Namespace)
	queries, err := queryLister.List(metav1.ListOptions{})
	if err != nil {
		return err
	}

	for _, query := range queries.Items {
		// For every query in the namespace, lookup it's dependencies, and if
		// it has a dependency on the passed in Report, requeue it
		deps, err := op.getGenerationQueryDependencies(query.Namespace, query.Name, nil)
		if err != nil {
			return err
		}
		for _, dependency := range deps.Reports {
			if dependency.Name == report.Name {
				// this query depends on the Report passed in
				op.enqueueReportGenerationQuery(query)
				break
			}
		}
	}
	return nil
}
