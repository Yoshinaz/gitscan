// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gitscan/internal/service/repo (interfaces: Interface)

// Package repoMocks is a generated GoMock package.
package repoMocks

import (
	reflect "reflect"

	rules "github.com/gitscan/rules"
	gomock "github.com/golang/mock/gomock"
	git "gopkg.in/src-d/go-git.v4"
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

// Clone mocks base method.
func (m *MockInterface) Clone() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clone")
	ret0, _ := ret[0].(error)
	return ret0
}

// Clone indicates an expected call of Clone.
func (mr *MockInterfaceMockRecorder) Clone() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockInterface)(nil).Clone))
}

// GetHeadCommitHash mocks base method.
func (m *MockInterface) GetHeadCommitHash() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeadCommitHash")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHeadCommitHash indicates an expected call of GetHeadCommitHash.
func (mr *MockInterfaceMockRecorder) GetHeadCommitHash() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeadCommitHash", reflect.TypeOf((*MockInterface)(nil).GetHeadCommitHash))
}

// Name mocks base method.
func (m *MockInterface) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockInterfaceMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockInterface)(nil).Name))
}

// Repo mocks base method.
func (m *MockInterface) Repo() *git.Repository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Repo")
	ret0, _ := ret[0].(*git.Repository)
	return ret0
}

// Repo indicates an expected call of Repo.
func (mr *MockInterfaceMockRecorder) Repo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Repo", reflect.TypeOf((*MockInterface)(nil).Repo))
}

// Rules mocks base method.
func (m *MockInterface) Rules() rules.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rules")
	ret0, _ := ret[0].(rules.Interface)
	return ret0
}

// Rules indicates an expected call of Rules.
func (mr *MockInterfaceMockRecorder) Rules() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rules", reflect.TypeOf((*MockInterface)(nil).Rules))
}

// URL mocks base method.
func (m *MockInterface) URL() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "URL")
	ret0, _ := ret[0].(string)
	return ret0
}

// URL indicates an expected call of URL.
func (mr *MockInterfaceMockRecorder) URL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "URL", reflect.TypeOf((*MockInterface)(nil).URL))
}
