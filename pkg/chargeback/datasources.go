package chargeback

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/tools/cache"

	cbTypes "github.com/coreos-inc/kube-chargeback/pkg/apis/chargeback/v1alpha1"
	"github.com/coreos-inc/kube-chargeback/pkg/aws"
	"github.com/coreos-inc/kube-chargeback/pkg/hive"
)

func (c *Chargeback) runReportDataSourceWorker() {
	logger := c.logger.WithField("component", "reportDataSourceWorker")
	logger.Infof("ReportDataSource worker started")
	for c.processReportDataSource(logger) {

	}
}

func (c *Chargeback) processReportDataSource(logger log.FieldLogger) bool {
	key, quit := c.queues.reportDataSourceQueue.Get()
	if quit {
		return false
	}
	defer c.queues.reportDataSourceQueue.Done(key)

	logger = logger.WithFields(c.newLogIdenti***REMOVED***er())
	err := c.syncReportDataSource(logger, key.(string))
	c.handleErr(logger, err, "ReportDataSource", key, c.queues.reportDataSourceQueue)
	return true
}

func (c *Chargeback) syncReportDataSource(logger log.FieldLogger, key string) error {
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		logger.WithError(err).Errorf("invalid resource key :%s", key)
		return nil
	}

	logger = logger.WithField("datasource", name)
	reportDataSource, err := c.informers.Chargeback().V1alpha1().ReportDataSources().Lister().ReportDataSources(namespace).Get(name)
	if err != nil {
		if apierrors.IsNotFound(err) {
			logger.Infof("ReportDataSource %s does not exist anymore", key)
			return nil
		}
		return err
	}

	logger.Infof("syncing reportDataSource %s", reportDataSource.GetName())
	err = c.handleReportDataSource(logger, reportDataSource)
	if err != nil {
		logger.WithError(err).Errorf("error syncing reportDataSource %s", reportDataSource.GetName())
		return err
	}
	logger.Infof("successfully synced reportDataSource %s", reportDataSource.GetName())
	return nil
}

func (c *Chargeback) handleReportDataSource(logger log.FieldLogger, dataSource *cbTypes.ReportDataSource) error {
	dataSource = dataSource.DeepCopy()
	if dataSource.TableName == "" {
		logger.Infof("new dataSource discovered")
	} ***REMOVED*** {
		logger.Infof("existing dataSource discovered, tableName: %s", dataSource.TableName)
		return nil
	}

	switch {
	case dataSource.Spec.Promsum != nil:
		return c.handlePromsumDataSource(logger, dataSource)
	case dataSource.Spec.AWSBilling != nil:
		return c.handleAWSBillingDataSource(logger, dataSource)
	default:
		return fmt.Errorf("datasource %s: improperly con***REMOVED***gured missing promsum or awsBilling con***REMOVED***guration", dataSource.Name)
	}
}

func (c *Chargeback) handlePromsumDataSource(logger log.FieldLogger, dataSource *cbTypes.ReportDataSource) error {
	storage := dataSource.Spec.Promsum.Storage
	tableName := dataSourceTableName(dataSource.Name)

	storageSpec, err := c.getStorageSpec(logger, storage, "datasource")
	if err != nil {
		return err
	}

	var createTableParams hive.CreateTableParameters
	if storageSpec.Local != nil {
		logger.Debugf("creating local table %s", tableName)
		createTableParams, err = hive.CreateLocalPromsumTable(c.hiveQueryer, tableName)
		if err != nil {
			return err
		}
	} ***REMOVED*** if storageSpec.S3 != nil {
		logger.Debugf("creating table %s backed by s3 bucket %s at pre***REMOVED***x %s", tableName, storageSpec.S3.Bucket, storageSpec.S3.Pre***REMOVED***x)
		createTableParams, err = hive.CreateS3PromsumTable(c.hiveQueryer, tableName, storageSpec.S3.Bucket, storageSpec.S3.Pre***REMOVED***x)
		if err != nil {
			return err
		}
	} ***REMOVED*** {
		return fmt.Errorf("storage incorrectly con***REMOVED***gured on datasource %s", dataSource.Name)
	}

	logger.Debugf("creating presto table CR for table %q", tableName)
	err = c.createPrestoTableCR(dataSource, cbTypes.GroupName, "datasource", createTableParams)
	if err != nil {
		logger.WithError(err).Errorf("failed to create PrestoTable CR %q", tableName)
		return err
	}

	logger.Debugf("successfully created table %s", tableName)

	return c.updateDataSourceTableName(logger, dataSource, tableName)
}

func (c *Chargeback) handleAWSBillingDataSource(logger log.FieldLogger, dataSource *cbTypes.ReportDataSource) error {
	source := dataSource.Spec.AWSBilling.Source
	if source == nil {
		return fmt.Errorf("datasource %q: improperly con***REMOVED***gured datasource, source is empty", dataSource.Name)
	}

	manifestRetriever, err := aws.NewManifestRetriever(source.Bucket, source.Pre***REMOVED***x)
	if err != nil {
		return err
	}

	manifests, err := manifestRetriever.RetrieveManifests()
	if err != nil {
		return err
	}

	if len(manifests) == 0 {
		logger.Warnf("datasource %q has no report manifests in it's bucket, the ***REMOVED***rst report has likely not been generated yet", dataSource.Name)
		return nil
	}

	tableName := dataSourceTableName(dataSource.Name)
	logger.Debugf("creating AWS Billing DataSource table %s pointing to s3 bucket %s at pre***REMOVED***x %s", tableName, source.Bucket, source.Pre***REMOVED***x)
	createTableParams, err := hive.CreateAWSUsageTable(c.hiveQueryer, tableName, source.Bucket, source.Pre***REMOVED***x, manifests)
	if err != nil {
		return err
	}

	logger.Debugf("creating presto table CR for table %q", tableName)
	err = c.createPrestoTableCR(dataSource, cbTypes.GroupName, "datasource", createTableParams)
	if err != nil {
		logger.WithError(err).Errorf("failed to create PrestoTable CR %q", tableName)
		return err
	}

	logger.Debugf("successfully created AWS Billing DataSource table %s pointing to s3 bucket %s at pre***REMOVED***x %s", tableName, source.Bucket, source.Pre***REMOVED***x)

	err = c.updateDataSourceTableName(logger, dataSource, tableName)
	if err != nil {
		return err
	}

	c.prestoTablePartitionQueue <- dataSource
	return nil
}

func (c *Chargeback) updateDataSourceTableName(logger log.FieldLogger, dataSource *cbTypes.ReportDataSource, tableName string) error {
	dataSource.TableName = tableName
	_, err := c.chargebackClient.ChargebackV1alpha1().ReportDataSources(dataSource.Namespace).Update(dataSource)
	if err != nil {
		logger.WithError(err).Errorf("failed to update ReportDataSource table name for %q", dataSource.Name)
		return err
	}
	return nil
}
