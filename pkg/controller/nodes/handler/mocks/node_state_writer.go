// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import handler "github.com/lyft/flytepropeller/pkg/controller/nodes/handler"
import mock "github.com/stretchr/testify/mock"

// NodeStateWriter is an autogenerated mock type for the NodeStateWriter type
type NodeStateWriter struct {
	mock.Mock
}

// PutBranchNode provides a mock function with given fields: s
func (_m *NodeStateWriter) PutBranchNode(s handler.BranchNodeState) error {
	ret := _m.Called(s)

	var r0 error
	if rf, ok := ret.Get(0).(func(handler.BranchNodeState) error); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PutDynamicNodeState provides a mock function with given fields: s
func (_m *NodeStateWriter) PutDynamicNodeState(s handler.DynamicNodeState) error {
	ret := _m.Called(s)

	var r0 error
	if rf, ok := ret.Get(0).(func(handler.DynamicNodeState) error); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PutTaskNodeState provides a mock function with given fields: s
func (_m *NodeStateWriter) PutTaskNodeState(s handler.TaskNodeState) error {
	ret := _m.Called(s)

	var r0 error
	if rf, ok := ret.Get(0).(func(handler.TaskNodeState) error); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
