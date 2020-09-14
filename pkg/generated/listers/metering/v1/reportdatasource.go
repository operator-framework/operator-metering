// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/kube-reporting/metering-operator/pkg/apis/metering/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ReportDataSourceLister helps list ReportDataSources.
type ReportDataSourceLister interface {
	// List lists all ReportDataSources in the indexer.
	List(selector labels.Selector) (ret []*v1.ReportDataSource, err error)
	// ReportDataSources returns an object that can list and get ReportDataSources.
	ReportDataSources(namespace string) ReportDataSourceNamespaceLister
	ReportDataSourceListerExpansion
}

// reportDataSourceLister implements the ReportDataSourceLister interface.
type reportDataSourceLister struct {
	indexer cache.Indexer
}

// NewReportDataSourceLister returns a new ReportDataSourceLister.
func NewReportDataSourceLister(indexer cache.Indexer) ReportDataSourceLister {
	return &reportDataSourceLister{indexer: indexer}
}

// List lists all ReportDataSources in the indexer.
func (s *reportDataSourceLister) List(selector labels.Selector) (ret []*v1.ReportDataSource, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ReportDataSource))
	})
	return ret, err
}

// ReportDataSources returns an object that can list and get ReportDataSources.
func (s *reportDataSourceLister) ReportDataSources(namespace string) ReportDataSourceNamespaceLister {
	return reportDataSourceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ReportDataSourceNamespaceLister helps list and get ReportDataSources.
type ReportDataSourceNamespaceLister interface {
	// List lists all ReportDataSources in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.ReportDataSource, err error)
	// Get retrieves the ReportDataSource from the indexer for a given namespace and name.
	Get(name string) (*v1.ReportDataSource, error)
	ReportDataSourceNamespaceListerExpansion
}

// reportDataSourceNamespaceLister implements the ReportDataSourceNamespaceLister
// interface.
type reportDataSourceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ReportDataSources in the indexer for a given namespace.
func (s reportDataSourceNamespaceLister) List(selector labels.Selector) (ret []*v1.ReportDataSource, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ReportDataSource))
	})
	return ret, err
}

// Get retrieves the ReportDataSource from the indexer for a given namespace and name.
func (s reportDataSourceNamespaceLister) Get(name string) (*v1.ReportDataSource, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("reportdatasource"), name)
	}
	return obj.(*v1.ReportDataSource), nil
}
