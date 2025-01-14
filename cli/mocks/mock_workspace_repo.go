// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/chanzuckerberg/happy/cli/pkg/workspace_repo (interfaces: WorkspaceRepoIface)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	workspace_repo "github.com/chanzuckerberg/happy/cli/pkg/workspace_repo"
	gomock "github.com/golang/mock/gomock"
)

// MockWorkspaceRepoIface is a mock of WorkspaceRepoIface interface.
type MockWorkspaceRepoIface struct {
	ctrl     *gomock.Controller
	recorder *MockWorkspaceRepoIfaceMockRecorder
}

// MockWorkspaceRepoIfaceMockRecorder is the mock recorder for MockWorkspaceRepoIface.
type MockWorkspaceRepoIfaceMockRecorder struct {
	mock *MockWorkspaceRepoIface
}

// NewMockWorkspaceRepoIface creates a new mock instance.
func NewMockWorkspaceRepoIface(ctrl *gomock.Controller) *MockWorkspaceRepoIface {
	mock := &MockWorkspaceRepoIface{ctrl: ctrl}
	mock.recorder = &MockWorkspaceRepoIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorkspaceRepoIface) EXPECT() *MockWorkspaceRepoIfaceMockRecorder {
	return m.recorder
}

// EstimateBacklogSize mocks base method.
func (m *MockWorkspaceRepoIface) EstimateBacklogSize(arg0 context.Context) (int, map[string]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EstimateBacklogSize", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(map[string]int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EstimateBacklogSize indicates an expected call of EstimateBacklogSize.
func (mr *MockWorkspaceRepoIfaceMockRecorder) EstimateBacklogSize(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EstimateBacklogSize", reflect.TypeOf((*MockWorkspaceRepoIface)(nil).EstimateBacklogSize), arg0)
}

// GetWorkspace mocks base method.
func (m *MockWorkspaceRepoIface) GetWorkspace(arg0 context.Context, arg1 string) (workspace_repo.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkspace", arg0, arg1)
	ret0, _ := ret[0].(workspace_repo.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkspace indicates an expected call of GetWorkspace.
func (mr *MockWorkspaceRepoIfaceMockRecorder) GetWorkspace(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkspace", reflect.TypeOf((*MockWorkspaceRepoIface)(nil).GetWorkspace), arg0, arg1)
}
