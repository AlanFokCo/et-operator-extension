package v1alpha1

import (
	"context"
	time "time"

	v1alpha1 "github.com/AlanFokCo/et-operator-extension/pkg/et-operator/apis/v1alpha1"
	versioned "github.com/AlanFokCo/et-operator-extension/pkg/et-operator/client/clientset/versioned"
	internalinterfaces "github.com/AlanFokCo/et-operator-extension/pkg/et-operator/client/informers/externalversions/internalinterfaces"
	v1 "github.com/AlanFokCo/et-operator-extension/pkg/et-operator/client/listers/tensorflow/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// TFJobInformer provides access to a shared informer and lister for
// TFJobs.
type TFJobInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.TFJobLister
}

type tFJobInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewTFJobInformer constructs a new informer for TFJob type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTFJobInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTFJobInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredTFJobInformer constructs a new informer for TFJob type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTFJobInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeflowV1().TFJobs(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeflowV1().TFJobs(namespace).Watch(context.TODO(), options)
			},
		},
		&tensorflowv1.TFJob{},
		resyncPeriod,
		indexers,
	)
}

func (f *tFJobInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTFJobInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *tFJobInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&tensorflowv1.TFJob{}, f.defaultInformer)
}

func (f *tFJobInformer) Lister() v1.TFJobLister {
	return v1.NewTFJobLister(f.Informer().GetIndexer())
}