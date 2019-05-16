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

// Code generated by informer-gen. DO NOT EDIT.

package v3

import (
	time "time"

	authcontainershipiov3 "github.com/containership/cluster-manager/pkg/apis/auth.containership.io/v3"
	versioned "github.com/containership/cluster-manager/pkg/client/clientset/versioned"
	internalinterfaces "github.com/containership/cluster-manager/pkg/client/informers/externalversions/internalinterfaces"
	v3 "github.com/containership/cluster-manager/pkg/client/listers/auth.containership.io/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// AuthorizationRoleBindingInformer provides access to a shared informer and lister for
// AuthorizationRoleBindings.
type AuthorizationRoleBindingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v3.AuthorizationRoleBindingLister
}

type authorizationRoleBindingInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewAuthorizationRoleBindingInformer constructs a new informer for AuthorizationRoleBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAuthorizationRoleBindingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAuthorizationRoleBindingInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredAuthorizationRoleBindingInformer constructs a new informer for AuthorizationRoleBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAuthorizationRoleBindingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ContainershipAuthV3().AuthorizationRoleBindings(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ContainershipAuthV3().AuthorizationRoleBindings(namespace).Watch(options)
			},
		},
		&authcontainershipiov3.AuthorizationRoleBinding{},
		resyncPeriod,
		indexers,
	)
}

func (f *authorizationRoleBindingInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAuthorizationRoleBindingInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *authorizationRoleBindingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&authcontainershipiov3.AuthorizationRoleBinding{}, f.defaultInformer)
}

func (f *authorizationRoleBindingInformer) Lister() v3.AuthorizationRoleBindingLister {
	return v3.NewAuthorizationRoleBindingLister(f.Informer().GetIndexer())
}