// Code generated by MockGen. DO NOT EDIT.
// Source: preflight.go

// Package preflight is a generated GoMock package.
package preflight

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta1 "github.com/rh-ecosystem-edge/kernel-module-management/api/v1beta1"
)

// MockPreflightAPI is a mock of PreflightAPI interface.
type MockPreflightAPI struct {
	ctrl     *gomock.Controller
	recorder *MockPreflightAPIMockRecorder
}

// MockPreflightAPIMockRecorder is the mock recorder for MockPreflightAPI.
type MockPreflightAPIMockRecorder struct {
	mock *MockPreflightAPI
}

// NewMockPreflightAPI creates a new mock instance.
func NewMockPreflightAPI(ctrl *gomock.Controller) *MockPreflightAPI {
	mock := &MockPreflightAPI{ctrl: ctrl}
	mock.recorder = &MockPreflightAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPreflightAPI) EXPECT() *MockPreflightAPIMockRecorder {
	return m.recorder
}

// PreflightUpgradeCheck mocks base method.
func (m *MockPreflightAPI) PreflightUpgradeCheck(ctx context.Context, pv *v1beta1.PreflightValidation, mod *v1beta1.Module) (bool, string) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PreflightUpgradeCheck", ctx, pv, mod)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(string)
	return ret0, ret1
}

// PreflightUpgradeCheck indicates an expected call of PreflightUpgradeCheck.
func (mr *MockPreflightAPIMockRecorder) PreflightUpgradeCheck(ctx, pv, mod interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PreflightUpgradeCheck", reflect.TypeOf((*MockPreflightAPI)(nil).PreflightUpgradeCheck), ctx, pv, mod)
}

// MockpreflightHelperAPI is a mock of preflightHelperAPI interface.
type MockpreflightHelperAPI struct {
	ctrl     *gomock.Controller
	recorder *MockpreflightHelperAPIMockRecorder
}

// MockpreflightHelperAPIMockRecorder is the mock recorder for MockpreflightHelperAPI.
type MockpreflightHelperAPIMockRecorder struct {
	mock *MockpreflightHelperAPI
}

// NewMockpreflightHelperAPI creates a new mock instance.
func NewMockpreflightHelperAPI(ctrl *gomock.Controller) *MockpreflightHelperAPI {
	mock := &MockpreflightHelperAPI{ctrl: ctrl}
	mock.recorder = &MockpreflightHelperAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockpreflightHelperAPI) EXPECT() *MockpreflightHelperAPIMockRecorder {
	return m.recorder
}

// verifyBuild mocks base method.
func (m *MockpreflightHelperAPI) verifyBuild(ctx context.Context, pv *v1beta1.PreflightValidation, mapping *v1beta1.KernelMapping, mod *v1beta1.Module) (bool, string) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "verifyBuild", ctx, pv, mapping, mod)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(string)
	return ret0, ret1
}

// verifyBuild indicates an expected call of verifyBuild.
func (mr *MockpreflightHelperAPIMockRecorder) verifyBuild(ctx, pv, mapping, mod interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "verifyBuild", reflect.TypeOf((*MockpreflightHelperAPI)(nil).verifyBuild), ctx, pv, mapping, mod)
}

// verifyImage mocks base method.
func (m *MockpreflightHelperAPI) verifyImage(ctx context.Context, mapping *v1beta1.KernelMapping, mod *v1beta1.Module, kernelVersion string) (bool, string) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "verifyImage", ctx, mapping, mod, kernelVersion)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(string)
	return ret0, ret1
}

// verifyImage indicates an expected call of verifyImage.
func (mr *MockpreflightHelperAPIMockRecorder) verifyImage(ctx, mapping, mod, kernelVersion interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "verifyImage", reflect.TypeOf((*MockpreflightHelperAPI)(nil).verifyImage), ctx, mapping, mod, kernelVersion)
}

// verifySign mocks base method.
func (m *MockpreflightHelperAPI) verifySign(ctx context.Context, pv *v1beta1.PreflightValidation, mapping *v1beta1.KernelMapping, mod *v1beta1.Module) (bool, string) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "verifySign", ctx, pv, mapping, mod)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(string)
	return ret0, ret1
}

// verifySign indicates an expected call of verifySign.
func (mr *MockpreflightHelperAPIMockRecorder) verifySign(ctx, pv, mapping, mod interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "verifySign", reflect.TypeOf((*MockpreflightHelperAPI)(nil).verifySign), ctx, pv, mapping, mod)
}
