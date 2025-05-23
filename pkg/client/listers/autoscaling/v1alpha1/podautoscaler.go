/*
Copyright 2024 The Aibrix Team.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/aibrix/aibrix/api/autoscaling/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PodAutoscalerLister helps list PodAutoscalers.
// All objects returned here must be treated as read-only.
type PodAutoscalerLister interface {
	// List lists all PodAutoscalers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PodAutoscaler, err error)
	// PodAutoscalers returns an object that can list and get PodAutoscalers.
	PodAutoscalers(namespace string) PodAutoscalerNamespaceLister
	PodAutoscalerListerExpansion
}

// podAutoscalerLister implements the PodAutoscalerLister interface.
type podAutoscalerLister struct {
	indexer cache.Indexer
}

// NewPodAutoscalerLister returns a new PodAutoscalerLister.
func NewPodAutoscalerLister(indexer cache.Indexer) PodAutoscalerLister {
	return &podAutoscalerLister{indexer: indexer}
}

// List lists all PodAutoscalers in the indexer.
func (s *podAutoscalerLister) List(selector labels.Selector) (ret []*v1alpha1.PodAutoscaler, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PodAutoscaler))
	})
	return ret, err
}

// PodAutoscalers returns an object that can list and get PodAutoscalers.
func (s *podAutoscalerLister) PodAutoscalers(namespace string) PodAutoscalerNamespaceLister {
	return podAutoscalerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PodAutoscalerNamespaceLister helps list and get PodAutoscalers.
// All objects returned here must be treated as read-only.
type PodAutoscalerNamespaceLister interface {
	// List lists all PodAutoscalers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PodAutoscaler, err error)
	// Get retrieves the PodAutoscaler from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.PodAutoscaler, error)
	PodAutoscalerNamespaceListerExpansion
}

// podAutoscalerNamespaceLister implements the PodAutoscalerNamespaceLister
// interface.
type podAutoscalerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PodAutoscalers in the indexer for a given namespace.
func (s podAutoscalerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PodAutoscaler, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PodAutoscaler))
	})
	return ret, err
}

// Get retrieves the PodAutoscaler from the indexer for a given namespace and name.
func (s podAutoscalerNamespaceLister) Get(name string) (*v1alpha1.PodAutoscaler, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("podautoscaler"), name)
	}
	return obj.(*v1alpha1.PodAutoscaler), nil
}
