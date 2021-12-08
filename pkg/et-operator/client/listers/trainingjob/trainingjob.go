package trainingjob

import (
	v1alpha1 "github.com/AlanFokCo/et-operator-extension/pkg/et-operator/apis/v1alpha1"
	v1 "github.com/kube-queue/tf-operator-extension/pkg/tf-operator/apis/tensorflow/v1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

)

type TrainingJobLister interface {
	List(selector labels.Selector) (ret []*v1alpha1.TrainingJob, err error)
	TrainingJobs(namespace string) TrainingJobNamespaceLister
	TrainingJobListerExpansion
}

// tFJobLister implements the TFJobLister interface.
type tFJobLister struct {
	indexer cache.Indexer
}

// NewTFJobLister returns a new TFJobLister.
func NewTrainingJobLister(indexer cache.Indexer) TFJobLister {
	return &tFJobLister{indexer: indexer}
}

// List lists all TFJobs in the indexer.
func (s *trainingJobLister) List(selector labels.Selector) (ret []*v1alpha1.TrainingJob, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.TFJob))
	})
	return ret, err
}

// TFJobs returns an object that can list and get TFJobs.
func (s *tFJobLister) TFJobs(namespace string) TFJobNamespaceLister {
	return tFJobNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TFJobNamespaceLister helps list and get TFJobs.
type TFJobNamespaceLister interface {
	// List lists all TFJobs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.TrainingJob, err error)
	// Get retrieves the TFJob from the indexer for a given namespace and name.
	Get(name string) (*v1.TFJob, error)
	TFJobNamespaceListerExpansion
}

// tFJobNamespaceLister implements the TFJobNamespaceLister
// interface.
type tFJobNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TFJobs in the indexer for a given namespace.
func (s tFJobNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.TrainingJob, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TrainingJob))
	})
	return ret, err
}

// Get retrieves the TFJob from the indexer for a given namespace and name.
func (s tFJobNamespaceLister) Get(name string) (*v1alpha1.TrainingJob, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("tfjob"), name)
	}
	return obj.(*v1alpha1.TrainingJob), nil
}