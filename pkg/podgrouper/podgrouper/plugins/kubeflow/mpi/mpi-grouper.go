// Copyright 2025 NVIDIA CORPORATION
// SPDX-License-Identifier: Apache-2.0

package mpi

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	"github.com/NVIDIA/KAI-scheduler/pkg/podgrouper/podgroup"
	"github.com/NVIDIA/KAI-scheduler/pkg/podgrouper/podgrouper/plugins/kubeflow"
)

const (
	ReplicaSpecName = "mpiReplicaSpecs"
	MasterName      = "Launcher"
	WorkerName      = "Worker"
)

func GetPodGroupMetadata(topOwner *unstructured.Unstructured, pod *v1.Pod, _ ...*metav1.PartialObjectMetadata) (*podgroup.Metadata, error) {
	return kubeflow.GetPodGroupMetadata(topOwner, pod,
		ReplicaSpecName, []string{MasterName, WorkerName})
}
