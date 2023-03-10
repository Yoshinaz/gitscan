// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gitscan/rules (interfaces: Interface)

// Package ruleMocks is a generated GoMock package.
package ruleMocks

import (
	reflect "reflect"

	report "github.com/gitscan/internal/report"
	rules "github.com/gitscan/rules"
	gomock "github.com/golang/mock/gomock"
	object "gopkg.in/src-d/go-git.v4/plumbing/object"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
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

// Add mocks base method.
func (m *MockInterface) Add(arg0 rules.RuleInfoInterface) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Add", arg0)
}

// Add indicates an expected call of Add.
func (mr *MockInterfaceMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockInterface)(nil).Add), arg0)
}

// GetMetaData mocks base method.
func (m *MockInterface) GetMetaData(arg0 string) report.Metadata {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetaData", arg0)
	ret0, _ := ret[0].(report.Metadata)
	return ret0
}

// GetMetaData indicates an expected call of GetMetaData.
func (mr *MockInterfaceMockRecorder) GetMetaData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetaData", reflect.TypeOf((*MockInterface)(nil).GetMetaData), arg0)
}

// Process mocks base method.
func (m *MockInterface) Process(arg0 *object.Commit) ([]report.Finding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Process", arg0)
	ret0, _ := ret[0].([]report.Finding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Process indicates an expected call of Process.
func (mr *MockInterfaceMockRecorder) Process(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Process", reflect.TypeOf((*MockInterface)(nil).Process), arg0)
}

// RuleSet mocks base method.
func (m *MockInterface) RuleSet() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RuleSet")
	ret0, _ := ret[0].(string)
	return ret0
}

// RuleSet indicates an expected call of RuleSet.
func (mr *MockInterfaceMockRecorder) RuleSet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RuleSet", reflect.TypeOf((*MockInterface)(nil).RuleSet))
}
