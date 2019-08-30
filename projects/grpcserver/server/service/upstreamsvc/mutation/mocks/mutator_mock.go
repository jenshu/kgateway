// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/solo-projects/projects/grpcserver/server/service/upstreamsvc/mutation (interfaces: Mutator)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	mutation "github.com/solo-io/solo-projects/projects/grpcserver/server/service/upstreamsvc/mutation"
)

// MockMutator is a mock of Mutator interface
type MockMutator struct {
	ctrl     *gomock.Controller
	recorder *MockMutatorMockRecorder
}

// MockMutatorMockRecorder is the mock recorder for MockMutator
type MockMutatorMockRecorder struct {
	mock *MockMutator
}

// NewMockMutator creates a new mock instance
func NewMockMutator(ctrl *gomock.Controller) *MockMutator {
	mock := &MockMutator{ctrl: ctrl}
	mock.recorder = &MockMutatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMutator) EXPECT() *MockMutatorMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockMutator) Create(arg0 context.Context, arg1 *core.ResourceRef, arg2 mutation.Mutation) (*v1.Upstream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Upstream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockMutatorMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMutator)(nil).Create), arg0, arg1, arg2)
}

// CreateUpstream mocks base method
func (m *MockMutator) CreateUpstream(arg0 context.Context, arg1 *v1.Upstream) (*v1.Upstream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUpstream", arg0, arg1)
	ret0, _ := ret[0].(*v1.Upstream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUpstream indicates an expected call of CreateUpstream
func (mr *MockMutatorMockRecorder) CreateUpstream(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUpstream", reflect.TypeOf((*MockMutator)(nil).CreateUpstream), arg0, arg1)
}

// Update mocks base method
func (m *MockMutator) Update(arg0 context.Context, arg1 *core.ResourceRef, arg2 mutation.Mutation) (*v1.Upstream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Upstream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockMutatorMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMutator)(nil).Update), arg0, arg1, arg2)
}

// UpdateUpstream mocks base method
func (m *MockMutator) UpdateUpstream(arg0 context.Context, arg1 *v1.Upstream) (*v1.Upstream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUpstream", arg0, arg1)
	ret0, _ := ret[0].(*v1.Upstream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUpstream indicates an expected call of UpdateUpstream
func (mr *MockMutatorMockRecorder) UpdateUpstream(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUpstream", reflect.TypeOf((*MockMutator)(nil).UpdateUpstream), arg0, arg1)
}
