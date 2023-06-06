// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

// Code generated by informer-gen. DO NOT EDIT.

package v2beta1

import (
	"context"
	time "time"

	spiderpoolspidernetiov2beta1 "github.com/spidernet-io/spiderpool/pkg/k8s/apis/spiderpool.spidernet.io/v2beta1"
	versioned "github.com/spidernet-io/spiderpool/pkg/k8s/client/clientset/versioned"
	internalinterfaces "github.com/spidernet-io/spiderpool/pkg/k8s/client/informers/externalversions/internalinterfaces"
	v2beta1 "github.com/spidernet-io/spiderpool/pkg/k8s/client/listers/spiderpool.spidernet.io/v2beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SpiderMultusConfigInformer provides access to a shared informer and lister for
// SpiderMultusConfigs.
type SpiderMultusConfigInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v2beta1.SpiderMultusConfigLister
}

type spiderMultusConfigInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSpiderMultusConfigInformer constructs a new informer for SpiderMultusConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSpiderMultusConfigInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSpiderMultusConfigInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSpiderMultusConfigInformer constructs a new informer for SpiderMultusConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSpiderMultusConfigInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SpiderpoolV2beta1().SpiderMultusConfigs(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SpiderpoolV2beta1().SpiderMultusConfigs(namespace).Watch(context.TODO(), options)
			},
		},
		&spiderpoolspidernetiov2beta1.SpiderMultusConfig{},
		resyncPeriod,
		indexers,
	)
}

func (f *spiderMultusConfigInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSpiderMultusConfigInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *spiderMultusConfigInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&spiderpoolspidernetiov2beta1.SpiderMultusConfig{}, f.defaultInformer)
}

func (f *spiderMultusConfigInformer) Lister() v2beta1.SpiderMultusConfigLister {
	return v2beta1.NewSpiderMultusConfigLister(f.Informer().GetIndexer())
}
