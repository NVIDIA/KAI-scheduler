/*
Copyright 2025 NVIDIA CORPORATION
SPDX-License-Identifier: Apache-2.0
*/
// Code generated by informer-gen. DO NOT EDIT.

package v2alpha2

import (
	context "context"
	time "time"

	versioned "github.com/NVIDIA/KAI-scheduler/pkg/apis/client/clientset/versioned"
	internalinterfaces "github.com/NVIDIA/KAI-scheduler/pkg/apis/client/informers/externalversions/internalinterfaces"
	schedulingv2alpha2 "github.com/NVIDIA/KAI-scheduler/pkg/apis/client/listers/scheduling/v2alpha2"
	apisschedulingv2alpha2 "github.com/NVIDIA/KAI-scheduler/pkg/apis/scheduling/v2alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PodGroupInformer provides access to a shared informer and lister for
// PodGroups.
type PodGroupInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() schedulingv2alpha2.PodGroupLister
}

type podGroupInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPodGroupInformer constructs a new informer for PodGroup type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPodGroupInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPodGroupInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPodGroupInformer constructs a new informer for PodGroup type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPodGroupInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SchedulingV2alpha2().PodGroups(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SchedulingV2alpha2().PodGroups(namespace).Watch(context.TODO(), options)
			},
		},
		&apisschedulingv2alpha2.PodGroup{},
		resyncPeriod,
		indexers,
	)
}

func (f *podGroupInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPodGroupInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *podGroupInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apisschedulingv2alpha2.PodGroup{}, f.defaultInformer)
}

func (f *podGroupInformer) Lister() schedulingv2alpha2.PodGroupLister {
	return schedulingv2alpha2.NewPodGroupLister(f.Informer().GetIndexer())
}
