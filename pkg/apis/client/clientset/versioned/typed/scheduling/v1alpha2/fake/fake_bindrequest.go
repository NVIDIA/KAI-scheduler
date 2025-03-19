/*
Copyright 2025 NVIDIA CORPORATION
SPDX-License-Identifier: Apache-2.0
*/
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	schedulingv1alpha2 "github.com/NVIDIA/KAI-scheduler/pkg/apis/client/clientset/versioned/typed/scheduling/v1alpha2"
	v1alpha2 "github.com/NVIDIA/KAI-scheduler/pkg/apis/scheduling/v1alpha2"
	gentype "k8s.io/client-go/gentype"
)

// fakeBindRequests implements BindRequestInterface
type fakeBindRequests struct {
	*gentype.FakeClientWithList[*v1alpha2.BindRequest, *v1alpha2.BindRequestList]
	Fake *FakeSchedulingV1alpha2
}

func newFakeBindRequests(fake *FakeSchedulingV1alpha2, namespace string) schedulingv1alpha2.BindRequestInterface {
	return &fakeBindRequests{
		gentype.NewFakeClientWithList[*v1alpha2.BindRequest, *v1alpha2.BindRequestList](
			fake.Fake,
			namespace,
			v1alpha2.SchemeGroupVersion.WithResource("bindrequests"),
			v1alpha2.SchemeGroupVersion.WithKind("BindRequest"),
			func() *v1alpha2.BindRequest { return &v1alpha2.BindRequest{} },
			func() *v1alpha2.BindRequestList { return &v1alpha2.BindRequestList{} },
			func(dst, src *v1alpha2.BindRequestList) { dst.ListMeta = src.ListMeta },
			func(list *v1alpha2.BindRequestList) []*v1alpha2.BindRequest {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1alpha2.BindRequestList, items []*v1alpha2.BindRequest) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
