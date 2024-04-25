// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	corev1 "k8s.io/api/core/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	mock "github.com/stretchr/testify/mock"

	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

// Interface is an autogenerated mock type for the Interface type
type Interface struct {
	mock.Mock
}

// Events provides a mock function with given fields:
func (_m *Interface) Events() v1.EventInterface {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Events")
	}

	var r0 v1.EventInterface
	if rf, ok := ret.Get(0).(func() v1.EventInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.EventInterface)
		}
	}

	return r0
}

// GetAnnotationsOnPod provides a mock function with given fields: namespace, name
func (_m *Interface) GetAnnotationsOnPod(namespace string, name string) (map[string]string, error) {
	ret := _m.Called(namespace, name)

	if len(ret) == 0 {
		panic("no return value specified for GetAnnotationsOnPod")
	}

	var r0 map[string]string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (map[string]string, error)); ok {
		return rf(namespace, name)
	}
	if rf, ok := ret.Get(0).(func(string, string) map[string]string); ok {
		r0 = rf(namespace, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(namespace, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNamespaces provides a mock function with given fields: labelSelector
func (_m *Interface) GetNamespaces(labelSelector metav1.LabelSelector) ([]*corev1.Namespace, error) {
	ret := _m.Called(labelSelector)

	if len(ret) == 0 {
		panic("no return value specified for GetNamespaces")
	}

	var r0 []*corev1.Namespace
	var r1 error
	if rf, ok := ret.Get(0).(func(metav1.LabelSelector) ([]*corev1.Namespace, error)); ok {
		return rf(labelSelector)
	}
	if rf, ok := ret.Get(0).(func(metav1.LabelSelector) []*corev1.Namespace); ok {
		r0 = rf(labelSelector)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*corev1.Namespace)
		}
	}

	if rf, ok := ret.Get(1).(func(metav1.LabelSelector) error); ok {
		r1 = rf(labelSelector)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNode provides a mock function with given fields: name
func (_m *Interface) GetNode(name string) (*corev1.Node, error) {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for GetNode")
	}

	var r0 *corev1.Node
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*corev1.Node, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) *corev1.Node); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Node)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNodes provides a mock function with given fields:
func (_m *Interface) GetNodes() ([]*corev1.Node, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetNodes")
	}

	var r0 []*corev1.Node
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*corev1.Node, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*corev1.Node); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*corev1.Node)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPod provides a mock function with given fields: namespace, name
func (_m *Interface) GetPod(namespace string, name string) (*corev1.Pod, error) {
	ret := _m.Called(namespace, name)

	if len(ret) == 0 {
		panic("no return value specified for GetPod")
	}

	var r0 *corev1.Pod
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*corev1.Pod, error)); ok {
		return rf(namespace, name)
	}
	if rf, ok := ret.Get(0).(func(string, string) *corev1.Pod); ok {
		r0 = rf(namespace, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Pod)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(namespace, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPods provides a mock function with given fields: namespace, opts
func (_m *Interface) GetPods(namespace string, opts metav1.ListOptions) ([]*corev1.Pod, error) {
	ret := _m.Called(namespace, opts)

	if len(ret) == 0 {
		panic("no return value specified for GetPods")
	}

	var r0 []*corev1.Pod
	var r1 error
	if rf, ok := ret.Get(0).(func(string, metav1.ListOptions) ([]*corev1.Pod, error)); ok {
		return rf(namespace, opts)
	}
	if rf, ok := ret.Get(0).(func(string, metav1.ListOptions) []*corev1.Pod); ok {
		r0 = rf(namespace, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*corev1.Pod)
		}
	}

	if rf, ok := ret.Get(1).(func(string, metav1.ListOptions) error); ok {
		r1 = rf(namespace, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PatchNode provides a mock function with given fields: old, new
func (_m *Interface) PatchNode(old *corev1.Node, new *corev1.Node) error {
	ret := _m.Called(old, new)

	if len(ret) == 0 {
		panic("no return value specified for PatchNode")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*corev1.Node, *corev1.Node) error); ok {
		r0 = rf(old, new)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveTaintFromNode provides a mock function with given fields: nodeName, taint
func (_m *Interface) RemoveTaintFromNode(nodeName string, taint *corev1.Taint) error {
	ret := _m.Called(nodeName, taint)

	if len(ret) == 0 {
		panic("no return value specified for RemoveTaintFromNode")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *corev1.Taint) error); ok {
		r0 = rf(nodeName, taint)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetAnnotationsOnNamespace provides a mock function with given fields: namespaceName, annotations
func (_m *Interface) SetAnnotationsOnNamespace(namespaceName string, annotations map[string]interface{}) error {
	ret := _m.Called(namespaceName, annotations)

	if len(ret) == 0 {
		panic("no return value specified for SetAnnotationsOnNamespace")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, map[string]interface{}) error); ok {
		r0 = rf(namespaceName, annotations)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetAnnotationsOnNode provides a mock function with given fields: nodeName, annotations
func (_m *Interface) SetAnnotationsOnNode(nodeName string, annotations map[string]interface{}) error {
	ret := _m.Called(nodeName, annotations)

	if len(ret) == 0 {
		panic("no return value specified for SetAnnotationsOnNode")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, map[string]interface{}) error); ok {
		r0 = rf(nodeName, annotations)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetAnnotationsOnPod provides a mock function with given fields: namespace, podName, annotations
func (_m *Interface) SetAnnotationsOnPod(namespace string, podName string, annotations map[string]interface{}) error {
	ret := _m.Called(namespace, podName, annotations)

	if len(ret) == 0 {
		panic("no return value specified for SetAnnotationsOnPod")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, map[string]interface{}) error); ok {
		r0 = rf(namespace, podName, annotations)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetAnnotationsOnService provides a mock function with given fields: namespace, serviceName, annotations
func (_m *Interface) SetAnnotationsOnService(namespace string, serviceName string, annotations map[string]interface{}) error {
	ret := _m.Called(namespace, serviceName, annotations)

	if len(ret) == 0 {
		panic("no return value specified for SetAnnotationsOnService")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, map[string]interface{}) error); ok {
		r0 = rf(namespace, serviceName, annotations)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetLabelsOnNode provides a mock function with given fields: nodeName, labels
func (_m *Interface) SetLabelsOnNode(nodeName string, labels map[string]interface{}) error {
	ret := _m.Called(nodeName, labels)

	if len(ret) == 0 {
		panic("no return value specified for SetLabelsOnNode")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, map[string]interface{}) error); ok {
		r0 = rf(nodeName, labels)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetTaintOnNode provides a mock function with given fields: nodeName, taint
func (_m *Interface) SetTaintOnNode(nodeName string, taint *corev1.Taint) error {
	ret := _m.Called(nodeName, taint)

	if len(ret) == 0 {
		panic("no return value specified for SetTaintOnNode")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *corev1.Taint) error); ok {
		r0 = rf(nodeName, taint)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateNodeStatus provides a mock function with given fields: node
func (_m *Interface) UpdateNodeStatus(node *corev1.Node) error {
	ret := _m.Called(node)

	if len(ret) == 0 {
		panic("no return value specified for UpdateNodeStatus")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*corev1.Node) error); ok {
		r0 = rf(node)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdatePodStatus provides a mock function with given fields: pod
func (_m *Interface) UpdatePodStatus(pod *corev1.Pod) error {
	ret := _m.Called(pod)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePodStatus")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*corev1.Pod) error); ok {
		r0 = rf(pod)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewInterface creates a new instance of Interface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *Interface {
	mock := &Interface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
