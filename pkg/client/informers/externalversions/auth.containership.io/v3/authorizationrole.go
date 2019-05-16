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

// AuthorizationRoleInformer provides access to a shared informer and lister for
// AuthorizationRoles.
type AuthorizationRoleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v3.AuthorizationRoleLister
}

type authorizationRoleInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewAuthorizationRoleInformer constructs a new informer for AuthorizationRole type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAuthorizationRoleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAuthorizationRoleInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredAuthorizationRoleInformer constructs a new informer for AuthorizationRole type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAuthorizationRoleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ContainershipAuthV3().AuthorizationRoles(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ContainershipAuthV3().AuthorizationRoles(namespace).Watch(options)
			},
		},
		&authcontainershipiov3.AuthorizationRole{},
		resyncPeriod,
		indexers,
	)
}

func (f *authorizationRoleInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAuthorizationRoleInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *authorizationRoleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&authcontainershipiov3.AuthorizationRole{}, f.defaultInformer)
}

func (f *authorizationRoleInformer) Lister() v3.AuthorizationRoleLister {
	return v3.NewAuthorizationRoleLister(f.Informer().GetIndexer())
}