// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/gloo/projects/controller/cli/pkg/cmd/install (interfaces: HelmUninstallation)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	release "helm.sh/helm/v3/pkg/release"
)

// MockHelmUninstallation is a mock of HelmUninstallation interface.
type MockHelmUninstallation struct {
	ctrl     *gomock.Controller
	recorder *MockHelmUninstallationMockRecorder
}

// MockHelmUninstallationMockRecorder is the mock recorder for MockHelmUninstallation.
type MockHelmUninstallationMockRecorder struct {
	mock *MockHelmUninstallation
}

// NewMockHelmUninstallation creates a new mock instance.
func NewMockHelmUninstallation(ctrl *gomock.Controller) *MockHelmUninstallation {
	mock := &MockHelmUninstallation{ctrl: ctrl}
	mock.recorder = &MockHelmUninstallationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHelmUninstallation) EXPECT() *MockHelmUninstallationMockRecorder {
	return m.recorder
}

// Run mocks base method.
func (m *MockHelmUninstallation) Run(arg0 string) (*release.UninstallReleaseResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0)
	ret0, _ := ret[0].(*release.UninstallReleaseResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Run indicates an expected call of Run.
func (mr *MockHelmUninstallationMockRecorder) Run(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockHelmUninstallation)(nil).Run), arg0)
}
