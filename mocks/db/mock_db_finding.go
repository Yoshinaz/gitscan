// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gitscan/internal/database (interfaces: FindingInterface)

// Package dbMocks is a generated GoMock package.
package dbMocks

import (
	reflect "reflect"

	database "github.com/gitscan/internal/database"
	gomock "github.com/golang/mock/gomock"
)

// MockFindingInterface is a mock of FindingInterface interface.
type MockFindingInterface struct {
	ctrl     *gomock.Controller
	recorder *MockFindingInterfaceMockRecorder
}

// MockFindingInterfaceMockRecorder is the mock recorder for MockFindingInterface.
type MockFindingInterfaceMockRecorder struct {
	mock *MockFindingInterface
}

// NewMockFindingInterface creates a new mock instance.
func NewMockFindingInterface(ctrl *gomock.Controller) *MockFindingInterface {
	mock := &MockFindingInterface{ctrl: ctrl}
	mock.recorder = &MockFindingInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFindingInterface) EXPECT() *MockFindingInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockFindingInterface) Create(arg0 database.Finding) (database.Finding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(database.Finding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockFindingInterfaceMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockFindingInterface)(nil).Create), arg0)
}

// Creates mocks base method.
func (m *MockFindingInterface) Creates(arg0 []database.Finding) ([]database.Finding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Creates", arg0)
	ret0, _ := ret[0].([]database.Finding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Creates indicates an expected call of Creates.
func (mr *MockFindingInterfaceMockRecorder) Creates(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Creates", reflect.TypeOf((*MockFindingInterface)(nil).Creates), arg0)
}

// FindByInfoID mocks base method.
func (m *MockFindingInterface) FindByInfoID(arg0 string) ([]database.Finding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByInfoID", arg0)
	ret0, _ := ret[0].([]database.Finding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByInfoID indicates an expected call of FindByInfoID.
func (mr *MockFindingInterfaceMockRecorder) FindByInfoID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByInfoID", reflect.TypeOf((*MockFindingInterface)(nil).FindByInfoID), arg0)
}

// FindByInfoIDAndCommit mocks base method.
func (m *MockFindingInterface) FindByInfoIDAndCommit(arg0, arg1 string) (database.Finding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByInfoIDAndCommit", arg0, arg1)
	ret0, _ := ret[0].(database.Finding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByInfoIDAndCommit indicates an expected call of FindByInfoIDAndCommit.
func (mr *MockFindingInterfaceMockRecorder) FindByInfoIDAndCommit(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByInfoIDAndCommit", reflect.TypeOf((*MockFindingInterface)(nil).FindByInfoIDAndCommit), arg0, arg1)
}