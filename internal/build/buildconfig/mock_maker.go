// Code generated by MockGen. DO NOT EDIT.
// Source: maker.go

// Package buildconfig is a generated GoMock package.
package buildconfig

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/openshift/api/build/v1"
	v1beta1 "github.com/rh-ecosystem-edge/kernel-module-management/api/v1beta1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MockMaker is a mock of Maker interface.
type MockMaker struct {
	ctrl     *gomock.Controller
	recorder *MockMakerMockRecorder
}

// MockMakerMockRecorder is the mock recorder for MockMaker.
type MockMakerMockRecorder struct {
	mock *MockMaker
}

// NewMockMaker creates a new mock instance.
func NewMockMaker(ctrl *gomock.Controller) *MockMaker {
	mock := &MockMaker{ctrl: ctrl}
	mock.recorder = &MockMakerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMaker) EXPECT() *MockMakerMockRecorder {
	return m.recorder
}

// MakeBuildTemplate mocks base method.
func (m *MockMaker) MakeBuildTemplate(ctx context.Context, mod v1beta1.Module, mapping v1beta1.KernelMapping, targetKernel string, pushImage bool, owner v10.Object) (*v1.Build, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeBuildTemplate", ctx, mod, mapping, targetKernel, pushImage, owner)
	ret0, _ := ret[0].(*v1.Build)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MakeBuildTemplate indicates an expected call of MakeBuildTemplate.
func (mr *MockMakerMockRecorder) MakeBuildTemplate(ctx, mod, mapping, targetKernel, pushImage, owner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeBuildTemplate", reflect.TypeOf((*MockMaker)(nil).MakeBuildTemplate), ctx, mod, mapping, targetKernel, pushImage, owner)
}
