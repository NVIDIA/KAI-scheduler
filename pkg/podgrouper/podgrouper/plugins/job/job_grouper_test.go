// Copyright 2025 NVIDIA CORPORATION
// SPDX-License-Identifier: Apache-2.0

package job

import (
	schedulingv2 "github.com/NVIDIA/KAI-scheduler/pkg/apis/scheduling/v2alpha2"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"testing"

	"github.com/stretchr/testify/assert"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestGetPodGroupMetadata_Hpo(t *testing.T) {
	owner := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "test_kind",
			"apiVersion": "test_version",
			"metadata": map[string]interface{}{
				"name":      "test_name",
				"namespace": "test_namespace",
				"uid":       "1234-5678",
				"labels": map[string]interface{}{
					"test_label": "test_value",
				},
				"annotations": map[string]interface{}{
					"test_annotation": "test_value",
				},
			},
			"spec": map[string]interface{}{
				"parallelism": int64(2),
			},
		},
	}
	pod := &v1.Pod{
		ObjectMeta: v12.ObjectMeta{
			Name: "test_name-4kgrb",
			Labels: map[string]string{
				"random": "abcd",
			},
		},
	}
	pod2 := &v1.Pod{
		ObjectMeta: v12.ObjectMeta{
			Name: "test_name-77t4m",
			Labels: map[string]string{
				"random": "abcd",
			},
		},
	}

	scheme := runtime.NewScheme()
	err := schedulingv2.AddToScheme(scheme)
	if err != nil {
		t.Fail()
	}

	client := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects().Build()

	runaiJobGrouper := NewK8sJobGrouper(client, false)

	podGroupMetadata, err := runaiJobGrouper.GetPodGroupMetadata(owner, pod)
	podGroupMetadata2, err2 := runaiJobGrouper.GetPodGroupMetadata(owner, pod2)

	assert.Nil(t, err)
	assert.Nil(t, err2)
	assert.Equal(t, "pg-test_name-4kgrb-1234-5678", podGroupMetadata.Name)
	assert.NotEqual(t, podGroupMetadata.Name, podGroupMetadata2.Name)
	assert.Equal(t, "pg-test_name-77t4m-1234-5678", podGroupMetadata2.Name)

	assert.Equal(t, "test_kind", podGroupMetadata.Owner.Kind)
	assert.Equal(t, "test_version", podGroupMetadata.Owner.APIVersion)
	assert.Equal(t, "1234-5678", string(podGroupMetadata.Owner.UID))
	assert.Equal(t, "test_name", podGroupMetadata.Owner.Name)
	assert.Equal(t, 3, len(podGroupMetadata.Annotations))
	assert.Equal(t, 1, len(podGroupMetadata.Labels))
	assert.Equal(t, "default-queue", podGroupMetadata.Queue)
	assert.Equal(t, "train", podGroupMetadata.PriorityClassName)
}

func TestGetPodGroupMetadata_LegacyPodGroup(t *testing.T) {
	owner := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "test_kind",
			"apiVersion": "test_version",
			"metadata": map[string]interface{}{
				"name":      "test_name",
				"namespace": "test_namespace",
				"uid":       "1234-5678",
				"labels": map[string]interface{}{
					"test_label": "test_value",
				},
				"annotations": map[string]interface{}{
					"test_annotation": "test_value",
				},
			},
			"spec": map[string]interface{}{},
		},
	}
	pod := &v1.Pod{
		ObjectMeta: v12.ObjectMeta{
			Name:      "test_name-4kgrb",
			Namespace: "test_namespace",
			Labels: map[string]string{
				"random": "abcd",
			},
		},
	}

	var runaiTestResources = []runtime.Object{
		&schedulingv2.PodGroup{
			TypeMeta: metav1.TypeMeta{
				Kind:       "PodGroup",
				APIVersion: "scheduling.run.ai/v2alpha2",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:        "pg-test_name-4kgrb-1234-5678",
				Namespace:   owner.GetNamespace(),
				UID:         owner.GetUID(),
				Annotations: owner.GetAnnotations(),
				Labels:      owner.GetLabels(),
			},
			Spec:   schedulingv2.PodGroupSpec{},
			Status: schedulingv2.PodGroupStatus{},
		},
	}

	scheme := runtime.NewScheme()
	err := schedulingv2.AddToScheme(scheme)
	if err != nil {
		t.Fail()
	}

	client := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(runaiTestResources...).Build()

	runaiJobGrouper := NewK8sJobGrouper(client, true)

	podGroupMetadata, err := runaiJobGrouper.GetPodGroupMetadata(owner, pod)

	assert.Nil(t, err)
	assert.Equal(t, "pg-test_name-4kgrb-1234-5678", podGroupMetadata.Name)

	assert.Equal(t, "test_kind", podGroupMetadata.Owner.Kind)
	assert.Equal(t, "test_version", podGroupMetadata.Owner.APIVersion)
	assert.Equal(t, "1234-5678", string(podGroupMetadata.Owner.UID))
	assert.Equal(t, "test_name", podGroupMetadata.Owner.Name)
	assert.Equal(t, 3, len(podGroupMetadata.Annotations))
	assert.Equal(t, 1, len(podGroupMetadata.Labels))
	assert.Equal(t, "default-queue", podGroupMetadata.Queue)
	assert.Equal(t, "train", podGroupMetadata.PriorityClassName)
}

func TestGetPodGroupMetadata_LegacyDisabledPodGroup(t *testing.T) {
	owner := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "test_kind",
			"apiVersion": "test_version",
			"metadata": map[string]interface{}{
				"name":      "test_name",
				"namespace": "test_namespace",
				"uid":       "1234-5678",
				"labels": map[string]interface{}{
					"test_label": "test_value",
				},
				"annotations": map[string]interface{}{
					"test_annotation": "test_value",
				},
			},
			"spec": map[string]interface{}{},
		},
	}
	pod := &v1.Pod{
		ObjectMeta: v12.ObjectMeta{
			Name:      "test_name-4kgrb",
			Namespace: "test_namespace",
			Labels: map[string]string{
				"random": "abcd",
			},
		},
	}

	var runaiTestResources = []runtime.Object{
		&schedulingv2.PodGroup{
			TypeMeta: metav1.TypeMeta{
				Kind:       "PodGroup",
				APIVersion: "scheduling.run.ai/v2alpha2",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:        "pg-test_name-4kgrb-1234-5678",
				Namespace:   owner.GetNamespace(),
				UID:         owner.GetUID(),
				Annotations: owner.GetAnnotations(),
				Labels:      owner.GetLabels(),
			},
			Spec:   schedulingv2.PodGroupSpec{},
			Status: schedulingv2.PodGroupStatus{},
		},
	}

	scheme := runtime.NewScheme()
	err := schedulingv2.AddToScheme(scheme)
	if err != nil {
		t.Fail()
	}

	client := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(runaiTestResources...).Build()

	runaiJobGrouper := NewK8sJobGrouper(client, false)

	podGroupMetadata, err := runaiJobGrouper.GetPodGroupMetadata(owner, pod)

	assert.Nil(t, err)
	assert.Equal(t, "pg-test_name-4kgrb-1234-5678", podGroupMetadata.Name)

	assert.Equal(t, "test_kind", podGroupMetadata.Owner.Kind)
	assert.Equal(t, "test_version", podGroupMetadata.Owner.APIVersion)
	assert.Equal(t, "1234-5678", string(podGroupMetadata.Owner.UID))
	assert.Equal(t, "test_name", podGroupMetadata.Owner.Name)
	assert.Equal(t, 3, len(podGroupMetadata.Annotations))
	assert.Equal(t, 1, len(podGroupMetadata.Labels))
	assert.Equal(t, "default-queue", podGroupMetadata.Queue)
	assert.Equal(t, "train", podGroupMetadata.PriorityClassName)
}

func TestGetPodGroupMetadata_LegacyNotFound(t *testing.T) {
	owner := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "test_kind",
			"apiVersion": "test_version",
			"metadata": map[string]interface{}{
				"name":      "test_name",
				"namespace": "test_namespace",
				"uid":       "1234-5678",
				"labels": map[string]interface{}{
					"test_label": "test_value",
				},
				"annotations": map[string]interface{}{
					"test_annotation": "test_value",
				},
			},
			"spec": map[string]interface{}{},
		},
	}
	pod := &v1.Pod{
		ObjectMeta: v12.ObjectMeta{
			Name:      "test_name-4kgrb",
			Namespace: "test_namespace",
			Labels: map[string]string{
				"random": "abcd",
			},
		},
	}

	var runaiTestResources = []runtime.Object{}

	scheme := runtime.NewScheme()
	err := schedulingv2.AddToScheme(scheme)
	if err != nil {
		t.Fail()
	}

	client := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(runaiTestResources...).Build()

	runaiJobGrouper := NewK8sJobGrouper(client, true)

	podGroupMetadata, err := runaiJobGrouper.GetPodGroupMetadata(owner, pod)

	assert.Nil(t, err)
	assert.Equal(t, "pg-test_name-4kgrb-1234-5678", podGroupMetadata.Name)

	assert.Equal(t, "test_kind", podGroupMetadata.Owner.Kind)
	assert.Equal(t, "test_version", podGroupMetadata.Owner.APIVersion)
	assert.Equal(t, "1234-5678", string(podGroupMetadata.Owner.UID))
	assert.Equal(t, "test_name", podGroupMetadata.Owner.Name)
	assert.Equal(t, 3, len(podGroupMetadata.Annotations))
	assert.Equal(t, 1, len(podGroupMetadata.Labels))
	assert.Equal(t, "default-queue", podGroupMetadata.Queue)
	assert.Equal(t, "train", podGroupMetadata.PriorityClassName)
}

func TestGetPodGroupMetadata_RegularPodGroup(t *testing.T) {
	owner := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "test_kind",
			"apiVersion": "test_version",
			"metadata": map[string]interface{}{
				"name":      "test_name",
				"namespace": "test_namespace",
				"uid":       "1234-5678",
				"labels": map[string]interface{}{
					"test_label": "test_value",
				},
				"annotations": map[string]interface{}{
					"test_annotation": "test_value",
				},
			},
			"spec": map[string]interface{}{},
		},
	}
	pod := &v1.Pod{
		ObjectMeta: v12.ObjectMeta{
			Name:      "test_name-4kgrb",
			Namespace: "test_namespace",
			Labels: map[string]string{
				"random": "abcd",
			},
		},
	}

	scheme := runtime.NewScheme()
	err := schedulingv2.AddToScheme(scheme)
	if err != nil {
		t.Fail()
	}

	client := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects().Build()

	runaiJobGrouper := NewK8sJobGrouper(client, false)

	podGroupMetadata, err := runaiJobGrouper.GetPodGroupMetadata(owner, pod)

	assert.Nil(t, err)
	assert.Equal(t, "pg-test_name-4kgrb-1234-5678", podGroupMetadata.Name)

	assert.Equal(t, "test_kind", podGroupMetadata.Owner.Kind)
	assert.Equal(t, "test_version", podGroupMetadata.Owner.APIVersion)
	assert.Equal(t, "1234-5678", string(podGroupMetadata.Owner.UID))
	assert.Equal(t, "test_name", podGroupMetadata.Owner.Name)
	assert.Equal(t, 3, len(podGroupMetadata.Annotations))
	assert.Equal(t, 1, len(podGroupMetadata.Labels))
	assert.Equal(t, "default-queue", podGroupMetadata.Queue)
	assert.Equal(t, "train", podGroupMetadata.PriorityClassName)
}
