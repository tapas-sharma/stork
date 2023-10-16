// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/portworx/sched-ops/k8s/kubevirt-dynamic (interfaces: Ops)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	kubevirtdynamic "github.com/portworx/sched-ops/k8s/kubevirt-dynamic"
	rest "k8s.io/client-go/rest"
)

// MockOps is a mock of Ops interface.
type MockOps struct {
	ctrl     *gomock.Controller
	recorder *MockOpsMockRecorder
}

// MockOpsMockRecorder is the mock recorder for MockOps.
type MockOpsMockRecorder struct {
	mock *MockOps
}

// NewMockOps creates a new mock instance.
func NewMockOps(ctrl *gomock.Controller) *MockOps {
	mock := &MockOps{ctrl: ctrl}
	mock.recorder = &MockOpsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOps) EXPECT() *MockOpsMockRecorder {
	return m.recorder
}

// CreateVirtualMachineInstanceMigration mocks base method.
func (m *MockOps) CreateVirtualMachineInstanceMigration(arg0 context.Context, arg1, arg2 string) (*kubevirtdynamic.VirtualMachineInstanceMigration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVirtualMachineInstanceMigration", arg0, arg1, arg2)
	ret0, _ := ret[0].(*kubevirtdynamic.VirtualMachineInstanceMigration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVirtualMachineInstanceMigration indicates an expected call of CreateVirtualMachineInstanceMigration.
func (mr *MockOpsMockRecorder) CreateVirtualMachineInstanceMigration(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVirtualMachineInstanceMigration", reflect.TypeOf((*MockOps)(nil).CreateVirtualMachineInstanceMigration), arg0, arg1, arg2)
}

// GetVirtualMachineInstance mocks base method.
func (m *MockOps) GetVirtualMachineInstance(arg0 context.Context, arg1, arg2 string) (*kubevirtdynamic.VirtualMachineInstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVirtualMachineInstance", arg0, arg1, arg2)
	ret0, _ := ret[0].(*kubevirtdynamic.VirtualMachineInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVirtualMachineInstance indicates an expected call of GetVirtualMachineInstance.
func (mr *MockOpsMockRecorder) GetVirtualMachineInstance(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVirtualMachineInstance", reflect.TypeOf((*MockOps)(nil).GetVirtualMachineInstance), arg0, arg1, arg2)
}

// GetVirtualMachineInstanceMigration mocks base method.
func (m *MockOps) GetVirtualMachineInstanceMigration(arg0 context.Context, arg1, arg2 string) (*kubevirtdynamic.VirtualMachineInstanceMigration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVirtualMachineInstanceMigration", arg0, arg1, arg2)
	ret0, _ := ret[0].(*kubevirtdynamic.VirtualMachineInstanceMigration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVirtualMachineInstanceMigration indicates an expected call of GetVirtualMachineInstanceMigration.
func (mr *MockOpsMockRecorder) GetVirtualMachineInstanceMigration(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVirtualMachineInstanceMigration", reflect.TypeOf((*MockOps)(nil).GetVirtualMachineInstanceMigration), arg0, arg1, arg2)
}

// ListVirtualMachineInstanceMigrations mocks base method.
func (m *MockOps) ListVirtualMachineInstanceMigrations(arg0 context.Context, arg1 string) ([]*kubevirtdynamic.VirtualMachineInstanceMigration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVirtualMachineInstanceMigrations", arg0, arg1)
	ret0, _ := ret[0].([]*kubevirtdynamic.VirtualMachineInstanceMigration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVirtualMachineInstanceMigrations indicates an expected call of ListVirtualMachineInstanceMigrations.
func (mr *MockOpsMockRecorder) ListVirtualMachineInstanceMigrations(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVirtualMachineInstanceMigrations", reflect.TypeOf((*MockOps)(nil).ListVirtualMachineInstanceMigrations), arg0, arg1)
}

// SetConfig mocks base method.
func (m *MockOps) SetConfig(arg0 *rest.Config) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetConfig", arg0)
}

// SetConfig indicates an expected call of SetConfig.
func (mr *MockOpsMockRecorder) SetConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConfig", reflect.TypeOf((*MockOps)(nil).SetConfig), arg0)
}
