// Code generated by MockGen. DO NOT EDIT.
// Source: ./translator.go

// Package mock_config is a generated GoMock package.
package mock_config

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "github.com/solo-io/ext-auth-plugins/api"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/extauth/v1"
)

// MockExtAuthConfigTranslator is a mock of ExtAuthConfigTranslator interface
type MockExtAuthConfigTranslator struct {
	ctrl     *gomock.Controller
	recorder *MockExtAuthConfigTranslatorMockRecorder
}

// MockExtAuthConfigTranslatorMockRecorder is the mock recorder for MockExtAuthConfigTranslator
type MockExtAuthConfigTranslatorMockRecorder struct {
	mock *MockExtAuthConfigTranslator
}

// NewMockExtAuthConfigTranslator creates a new mock instance
func NewMockExtAuthConfigTranslator(ctrl *gomock.Controller) *MockExtAuthConfigTranslator {
	mock := &MockExtAuthConfigTranslator{ctrl: ctrl}
	mock.recorder = &MockExtAuthConfigTranslatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockExtAuthConfigTranslator) EXPECT() *MockExtAuthConfigTranslatorMockRecorder {
	return m.recorder
}

// Translate mocks base method
func (m *MockExtAuthConfigTranslator) Translate(ctx context.Context, resource *v1.ExtAuthConfig) (api.AuthService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Translate", ctx, resource)
	ret0, _ := ret[0].(api.AuthService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Translate indicates an expected call of Translate
func (mr *MockExtAuthConfigTranslatorMockRecorder) Translate(ctx, resource interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Translate", reflect.TypeOf((*MockExtAuthConfigTranslator)(nil).Translate), ctx, resource)
}
