/*
Copyright 2025 NVIDIA CORPORATION
SPDX-License-Identifier: Apache-2.0
*/
// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	fmt "fmt"

	v1alpha2 "github.com/NVIDIA/KAI-scheduler/pkg/apis/scheduling/v1alpha2"
	v2 "github.com/NVIDIA/KAI-scheduler/pkg/apis/scheduling/v2"
	v2alpha2 "github.com/NVIDIA/KAI-scheduler/pkg/apis/scheduling/v2alpha2"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=scheduling.run.ai, Version=v1alpha2
	case v1alpha2.SchemeGroupVersion.WithResource("bindrequests"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Scheduling().V1alpha2().BindRequests().Informer()}, nil

		// Group=scheduling.run.ai, Version=v2
	case v2.SchemeGroupVersion.WithResource("queues"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Scheduling().V2().Queues().Informer()}, nil

		// Group=scheduling.run.ai, Version=v2alpha2
	case v2alpha2.SchemeGroupVersion.WithResource("podgroups"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Scheduling().V2alpha2().PodGroups().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
