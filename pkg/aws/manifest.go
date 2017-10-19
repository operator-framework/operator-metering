package aws

import (
	"path/***REMOVED***lepath"
	"time"

	"github.com/sirupsen/logrus"
)

// Manifest is a representation of the ***REMOVED***le AWS provides with metadata for current usage information.
type Manifest struct {
	AssemblyID             string        `json:"assemblyId"`
	Account                string        `json:"account"`
	Columns                []Column      `json:"columns"`
	Charset                string        `json:"charset"`
	Compression            string        `json:"compression"`
	ContentType            string        `json:"contentType"`
	ReportID               string        `json:"reportId"`
	ReportName             string        `json:"reportName"`
	BillingPeriod          BillingPeriod `json:"billingPeriod"`
	Bucket                 string        `json:"bucket"`
	ReportKeys             []string      `json:"reportKeys"`
	AdditionalArtifactKeys []string      `json:"additionalArtifactKeys"`
}

type BillingPeriod struct {
	Start Time `json:"start"`
	End   Time `json:"end"`
}

// Column is a description of a ***REMOVED***eld from a AWS usage report manifest ***REMOVED***le.
type Column struct {
	Category string `json:"category"`
	Name     string `json:"name"`
}

// Paths returns the directories containing usage data. The result will be free of duplicates.
func (m Manifest) DataDirectory() string {
	var dirPath string
	pathMap := make(map[string]struct{})
	for _, key := range m.ReportKeys {
		dirPath = ***REMOVED***lepath.Dir(key)
		pathMap[dirPath] = struct{}{}
	}

	if len(pathMap) != 1 {
		logrus.Errorf("aws manifest %s has multiple directories containing usage data, expected 1, reportKeys: %v", m.AssemblyID, m.ReportKeys)
	}

	return dirPath
}

type Time struct {
	time.Time
}

const manifestTime = "20060102T000000.000Z"

func (t *Time) UnmarshalJSON(b []byte) error {
	// b contains quotes around the timestamp
	tt, err := time.Parse(manifestTime, string(b[1:len(b)-1]))
	if err == nil {
		*t = Time{tt}
	}
	return err
}

func (t *Time) String() string {
	return t.Format(manifestTime)
}
