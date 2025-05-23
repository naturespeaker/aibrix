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
	v1alpha1 "github.com/aibrix/aibrix/api/orchestration/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RayClusterReplicaSetLister helps list RayClusterReplicaSets.
// All objects returned here must be treated as read-only.
type RayClusterReplicaSetLister interface {
	// List lists all RayClusterReplicaSets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.RayClusterReplicaSet, err error)
	// RayClusterReplicaSets returns an object that can list and get RayClusterReplicaSets.
	RayClusterReplicaSets(namespace string) RayClusterReplicaSetNamespaceLister
	RayClusterReplicaSetListerExpansion
}

// rayClusterReplicaSetLister implements the RayClusterReplicaSetLister interface.
type rayClusterReplicaSetLister struct {
	indexer cache.Indexer
}

// NewRayClusterReplicaSetLister returns a new RayClusterReplicaSetLister.
func NewRayClusterReplicaSetLister(indexer cache.Indexer) RayClusterReplicaSetLister {
	return &rayClusterReplicaSetLister{indexer: indexer}
}

// List lists all RayClusterReplicaSets in the indexer.
func (s *rayClusterReplicaSetLister) List(selector labels.Selector) (ret []*v1alpha1.RayClusterReplicaSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RayClusterReplicaSet))
	})
	return ret, err
}

// RayClusterReplicaSets returns an object that can list and get RayClusterReplicaSets.
func (s *rayClusterReplicaSetLister) RayClusterReplicaSets(namespace string) RayClusterReplicaSetNamespaceLister {
	return rayClusterReplicaSetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RayClusterReplicaSetNamespaceLister helps list and get RayClusterReplicaSets.
// All objects returned here must be treated as read-only.
type RayClusterReplicaSetNamespaceLister interface {
	// List lists all RayClusterReplicaSets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.RayClusterReplicaSet, err error)
	// Get retrieves the RayClusterReplicaSet from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.RayClusterReplicaSet, error)
	RayClusterReplicaSetNamespaceListerExpansion
}

// rayClusterReplicaSetNamespaceLister implements the RayClusterReplicaSetNamespaceLister
// interface.
type rayClusterReplicaSetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RayClusterReplicaSets in the indexer for a given namespace.
func (s rayClusterReplicaSetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.RayClusterReplicaSet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RayClusterReplicaSet))
	})
	return ret, err
}

// Get retrieves the RayClusterReplicaSet from the indexer for a given namespace and name.
func (s rayClusterReplicaSetNamespaceLister) Get(name string) (*v1alpha1.RayClusterReplicaSet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("rayclusterreplicaset"), name)
	}
	return obj.(*v1alpha1.RayClusterReplicaSet), nil
}
