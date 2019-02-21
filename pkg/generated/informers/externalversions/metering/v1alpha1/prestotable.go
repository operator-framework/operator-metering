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

// PrestoTableInformer provides access to a shared informer and lister for
// PrestoTables.
type PrestoTableInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.PrestoTableLister
}

type prestoTableInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPrestoTableInformer constructs a new informer for PrestoTable type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPrestoTableInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPrestoTableInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPrestoTableInformer constructs a new informer for PrestoTable type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPrestoTableInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MeteringV1alpha1().PrestoTables(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MeteringV1alpha1().PrestoTables(namespace).Watch(options)
			},
		},
		&meteringv1alpha1.PrestoTable{},
		resyncPeriod,
		indexers,
	)
}

func (f *prestoTableInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPrestoTableInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *prestoTableInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&meteringv1alpha1.PrestoTable{}, f.defaultInformer)
}

func (f *prestoTableInformer) Lister() v1alpha1.PrestoTableLister {
	return v1alpha1.NewPrestoTableLister(f.Informer().GetIndexer())
}
