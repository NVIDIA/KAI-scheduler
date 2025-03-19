// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/scheduler/k8s_utils/k8s_utils.go
//
// Generated by this command:
//
//	mockgen -source=pkg/scheduler/k8s_utils/k8s_utils.go -destination=pkg/scheduler/k8s_utils/k8s_utils_mock.go -package=k8s_utils
//

// Package k8s_utils is a generated GoMock package.
package k8s_utils

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	v1 "k8s.io/api/core/v1"
	kubernetes "k8s.io/client-go/kubernetes"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
	isgomock struct{}
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// PatchPodAnnotationsAndLabelsInterface mocks base method.
func (m *MockInterface) PatchPodAnnotationsAndLabelsInterface(kubeclient kubernetes.Interface, pod *v1.Pod, annotations, labels map[string]any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchPodAnnotationsAndLabelsInterface", kubeclient, pod, annotations, labels)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchPodAnnotationsAndLabelsInterface indicates an expected call of PatchPodAnnotationsAndLabelsInterface.
func (mr *MockInterfaceMockRecorder) PatchPodAnnotationsAndLabelsInterface(kubeclient, pod, annotations, labels any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchPodAnnotationsAndLabelsInterface", reflect.TypeOf((*MockInterface)(nil).PatchPodAnnotationsAndLabelsInterface), kubeclient, pod, annotations, labels)
}
