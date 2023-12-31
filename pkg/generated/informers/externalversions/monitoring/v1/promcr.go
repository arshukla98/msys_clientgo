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

package v1

import (
	"context"
	time "time"

	monitoringv1 "github.com/arshukla98/msys_clientgo/pkg/apis/monitoring/v1"
	versioned "github.com/arshukla98/msys_clientgo/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/arshukla98/msys_clientgo/pkg/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/arshukla98/msys_clientgo/pkg/generated/listers/monitoring/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PromCRInformer provides access to a shared informer and lister for
// PromCRs.
type PromCRInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.PromCRLister
}

type promCRInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPromCRInformer constructs a new informer for PromCR type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPromCRInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPromCRInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPromCRInformer constructs a new informer for PromCR type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPromCRInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MonitoringV1().PromCRs(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MonitoringV1().PromCRs(namespace).Watch(context.TODO(), options)
			},
		},
		&monitoringv1.PromCR{},
		resyncPeriod,
		indexers,
	)
}

func (f *promCRInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPromCRInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *promCRInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&monitoringv1.PromCR{}, f.defaultInformer)
}

func (f *promCRInformer) Lister() v1.PromCRLister {
	return v1.NewPromCRLister(f.Informer().GetIndexer())
}
