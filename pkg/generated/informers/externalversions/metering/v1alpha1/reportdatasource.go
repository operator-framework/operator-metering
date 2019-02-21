// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	meteringv1alpha1 "github.com/operator-framework/operator-metering/pkg/apis/metering/v1alpha1"
	versioned "github.com/operator-framework/operator-metering/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/operator-framework/operator-metering/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/operator-framework/operator-metering/pkg/generated/listers/metering/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ReportDataSourceInformer provides access to a shared informer and lister for
// ReportDataSources.
type ReportDataSourceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ReportDataSourceLister
}

type reportDataSourceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewReportDataSourceInformer constructs a new informer for ReportDataSource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewReportDataSourceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredReportDataSourceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredReportDataSourceInformer constructs a new informer for ReportDataSource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredReportDataSourceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MeteringV1alpha1().ReportDataSources(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MeteringV1alpha1().ReportDataSources(namespace).Watch(options)
			},
		},
		&meteringv1alpha1.ReportDataSource{},
		resyncPeriod,
		indexers,
	)
}

func (f *reportDataSourceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredReportDataSourceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *reportDataSourceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&meteringv1alpha1.ReportDataSource{}, f.defaultInformer)
}

func (f *reportDataSourceInformer) Lister() v1alpha1.ReportDataSourceLister {
	return v1alpha1.NewReportDataSourceLister(f.Informer().GetIndexer())
}
