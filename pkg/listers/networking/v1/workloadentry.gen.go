// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "istio.io/client-go/pkg/apis/networking/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// WorkloadEntryLister helps list WorkloadEntries.
// All objects returned here must be treated as read-only.
type WorkloadEntryLister interface {
	// List lists all WorkloadEntries in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.WorkloadEntry, err error)
	// WorkloadEntries returns an object that can list and get WorkloadEntries.
	WorkloadEntries(namespace string) WorkloadEntryNamespaceLister
	WorkloadEntryListerExpansion
}

// workloadEntryLister implements the WorkloadEntryLister interface.
type workloadEntryLister struct {
	indexer cache.Indexer
}

// NewWorkloadEntryLister returns a new WorkloadEntryLister.
func NewWorkloadEntryLister(indexer cache.Indexer) WorkloadEntryLister {
	return &workloadEntryLister{indexer: indexer}
}

// List lists all WorkloadEntries in the indexer.
func (s *workloadEntryLister) List(selector labels.Selector) (ret []*v1.WorkloadEntry, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.WorkloadEntry))
	})
	return ret, err
}

// WorkloadEntries returns an object that can list and get WorkloadEntries.
func (s *workloadEntryLister) WorkloadEntries(namespace string) WorkloadEntryNamespaceLister {
	return workloadEntryNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// WorkloadEntryNamespaceLister helps list and get WorkloadEntries.
// All objects returned here must be treated as read-only.
type WorkloadEntryNamespaceLister interface {
	// List lists all WorkloadEntries in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.WorkloadEntry, err error)
	// Get retrieves the WorkloadEntry from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.WorkloadEntry, error)
	WorkloadEntryNamespaceListerExpansion
}

// workloadEntryNamespaceLister implements the WorkloadEntryNamespaceLister
// interface.
type workloadEntryNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all WorkloadEntries in the indexer for a given namespace.
func (s workloadEntryNamespaceLister) List(selector labels.Selector) (ret []*v1.WorkloadEntry, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.WorkloadEntry))
	})
	return ret, err
}

// Get retrieves the WorkloadEntry from the indexer for a given namespace and name.
func (s workloadEntryNamespaceLister) Get(name string) (*v1.WorkloadEntry, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("workloadentry"), name)
	}
	return obj.(*v1.WorkloadEntry), nil
}
