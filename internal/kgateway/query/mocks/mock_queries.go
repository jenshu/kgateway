// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kgateway-dev/kgateway/internal/kgateway/query (interfaces: GatewayQueries)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	krt "istio.io/istio/pkg/kube/krt"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	v1 "sigs.k8s.io/gateway-api/apis/v1"

	ir "github.com/kgateway-dev/kgateway/v2/internal/kgateway/ir"
	query "github.com/kgateway-dev/kgateway/v2/internal/kgateway/query"
)

// MockGatewayQueries is a mock of GatewayQueries interface.
type MockGatewayQueries struct {
	ctrl     *gomock.Controller
	recorder *MockGatewayQueriesMockRecorder
}

// MockGatewayQueriesMockRecorder is the mock recorder for MockGatewayQueries.
type MockGatewayQueriesMockRecorder struct {
	mock *MockGatewayQueries
}

// NewMockGatewayQueries creates a new mock instance.
func NewMockGatewayQueries(ctrl *gomock.Controller) *MockGatewayQueries {
	mock := &MockGatewayQueries{ctrl: ctrl}
	mock.recorder = &MockGatewayQueriesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGatewayQueries) EXPECT() *MockGatewayQueriesMockRecorder {
	return m.recorder
}

// GetRouteChain mocks base method.
func (m *MockGatewayQueries) GetRouteChain(arg0 krt.HandlerContext, arg1 context.Context, arg2 ir.Route, arg3 []string, arg4 v1.ParentReference) *query.RouteInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRouteChain", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*query.RouteInfo)
	return ret0
}

// GetRouteChain indicates an expected call of GetRouteChain.
func (mr *MockGatewayQueriesMockRecorder) GetRouteChain(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRouteChain", reflect.TypeOf((*MockGatewayQueries)(nil).GetRouteChain), arg0, arg1, arg2, arg3, arg4)
}

// GetRoutesForGateway mocks base method.
func (m *MockGatewayQueries) GetRoutesForGateway(arg0 krt.HandlerContext, arg1 context.Context, arg2 *v1.Gateway) (*query.RoutesForGwResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoutesForGateway", arg0, arg1, arg2)
	ret0, _ := ret[0].(*query.RoutesForGwResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoutesForGateway indicates an expected call of GetRoutesForGateway.
func (mr *MockGatewayQueriesMockRecorder) GetRoutesForGateway(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoutesForGateway", reflect.TypeOf((*MockGatewayQueries)(nil).GetRoutesForGateway), arg0, arg1, arg2)
}

// GetSecretForRef mocks base method.
func (m *MockGatewayQueries) GetSecretForRef(arg0 krt.HandlerContext, arg1 context.Context, arg2 schema.GroupKind, arg3 string, arg4 v1.SecretObjectReference) (*ir.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretForRef", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*ir.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecretForRef indicates an expected call of GetSecretForRef.
func (mr *MockGatewayQueriesMockRecorder) GetSecretForRef(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretForRef", reflect.TypeOf((*MockGatewayQueries)(nil).GetSecretForRef), arg0, arg1, arg2, arg3, arg4)
}
