// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/kube-reporting/metering-operator/pkg/apis/metering/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ReportQueryLister helps list ReportQueries.
// All objects returned here must be treated as read-only.
type ReportQueryLister interface {
	// List lists all ReportQueries in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ReportQuery, err error)
	// ReportQueries returns an object that can list and get ReportQueries.
	ReportQueries(namespace string) ReportQueryNamespaceLister
	ReportQueryListerExpansion
}

// reportQueryLister implements the ReportQueryLister interface.
type reportQueryLister struct {
	indexer cache.Indexer
}

// NewReportQueryLister returns a new ReportQueryLister.
func NewReportQueryLister(indexer cache.Indexer) ReportQueryLister {
	return &reportQueryLister{indexer: indexer}
}

// List lists all ReportQueries in the indexer.
func (s *reportQueryLister) List(selector labels.Selector) (ret []*v1.ReportQuery, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ReportQuery))
	})
	return ret, err
}

// ReportQueries returns an object that can list and get ReportQueries.
func (s *reportQueryLister) ReportQueries(namespace string) ReportQueryNamespaceLister {
	return reportQueryNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ReportQueryNamespaceLister helps list and get ReportQueries.
// All objects returned here must be treated as read-only.
type ReportQueryNamespaceLister interface {
	// List lists all ReportQueries in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ReportQuery, err error)
	// Get retrieves the ReportQuery from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ReportQuery, error)
	ReportQueryNamespaceListerExpansion
}

// reportQueryNamespaceLister implements the ReportQueryNamespaceLister
// interface.
type reportQueryNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ReportQueries in the indexer for a given namespace.
func (s reportQueryNamespaceLister) List(selector labels.Selector) (ret []*v1.ReportQuery, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ReportQuery))
	})
	return ret, err
}

// Get retrieves the ReportQuery from the indexer for a given namespace and name.
func (s reportQueryNamespaceLister) Get(name string) (*v1.ReportQuery, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("reportquery"), name)
	}
	return obj.(*v1.ReportQuery), nil
}
