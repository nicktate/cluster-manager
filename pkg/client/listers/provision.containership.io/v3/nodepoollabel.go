/*
Copyright The Kubernetes Authors.

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

package v3

import (
	v3 "github.com/containership/cluster-manager/pkg/apis/provision.containership.io/v3"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NodePoolLabelLister helps list NodePoolLabels.
type NodePoolLabelLister interface {
	// List lists all NodePoolLabels in the indexer.
	List(selector labels.Selector) (ret []*v3.NodePoolLabel, err error)
	// NodePoolLabels returns an object that can list and get NodePoolLabels.
	NodePoolLabels(namespace string) NodePoolLabelNamespaceLister
	NodePoolLabelListerExpansion
}

// nodePoolLabelLister implements the NodePoolLabelLister interface.
type nodePoolLabelLister struct {
	indexer cache.Indexer
}

// NewNodePoolLabelLister returns a new NodePoolLabelLister.
func NewNodePoolLabelLister(indexer cache.Indexer) NodePoolLabelLister {
	return &nodePoolLabelLister{indexer: indexer}
}

// List lists all NodePoolLabels in the indexer.
func (s *nodePoolLabelLister) List(selector labels.Selector) (ret []*v3.NodePoolLabel, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.NodePoolLabel))
	})
	return ret, err
}

// NodePoolLabels returns an object that can list and get NodePoolLabels.
func (s *nodePoolLabelLister) NodePoolLabels(namespace string) NodePoolLabelNamespaceLister {
	return nodePoolLabelNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NodePoolLabelNamespaceLister helps list and get NodePoolLabels.
type NodePoolLabelNamespaceLister interface {
	// List lists all NodePoolLabels in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v3.NodePoolLabel, err error)
	// Get retrieves the NodePoolLabel from the indexer for a given namespace and name.
	Get(name string) (*v3.NodePoolLabel, error)
	NodePoolLabelNamespaceListerExpansion
}

// nodePoolLabelNamespaceLister implements the NodePoolLabelNamespaceLister
// interface.
type nodePoolLabelNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NodePoolLabels in the indexer for a given namespace.
func (s nodePoolLabelNamespaceLister) List(selector labels.Selector) (ret []*v3.NodePoolLabel, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.NodePoolLabel))
	})
	return ret, err
}

// Get retrieves the NodePoolLabel from the indexer for a given namespace and name.
func (s nodePoolLabelNamespaceLister) Get(name string) (*v3.NodePoolLabel, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v3.Resource("nodepoollabel"), name)
	}
	return obj.(*v3.NodePoolLabel), nil
}
