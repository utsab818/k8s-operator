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

package internalversion

import (
	"context"
	time "time"

	internalinterfaces "github.com/utsab818/kluster/pkg/client/informers/internalversion/internalinterfaces"
	internalversion "github.com/utsab818/kluster/pkg/client/listers/v1alpha1/internalversion"
	v1alpha1 "github.com/utsab818/kluster/pkg/apis/utsab.dev/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	// internalclientset "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset"
	internalclientset "github.com/utsab818/kluster/pkg/client/clientset/versioned"
)

// KlusterInformer provides access to a shared informer and lister for
// Klusters.
type KlusterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.KlusterLister
}

type klusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewKlusterInformer constructs a new informer for Kluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewKlusterInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredKlusterInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredKlusterInformer constructs a new informer for Kluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredKlusterInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.UtsabV1alpha1().Klusters(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.UtsabV1alpha1().Klusters(namespace).Watch(context.TODO(), options)
			},
		},
		&v1alpha1.Kluster{},
		resyncPeriod,
		indexers,
	)
}

func (f *klusterInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredKlusterInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *klusterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&v1alpha1.Kluster{}, f.defaultInformer)
}

func (f *klusterInformer) Lister() internalversion.KlusterLister {
	return internalversion.NewKlusterLister(f.Informer().GetIndexer())
}
