// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/chanzuckerberg/happy/cli/pkg/artifact_builder (interfaces: ArtifactBuilderIface)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	artifact_builder "github.com/chanzuckerberg/happy/cli/pkg/artifact_builder"
	aws "github.com/chanzuckerberg/happy/cli/pkg/backend/aws"
	config "github.com/chanzuckerberg/happy/cli/pkg/config"
	gomock "github.com/golang/mock/gomock"
)

// MockArtifactBuilderIface is a mock of ArtifactBuilderIface interface.
type MockArtifactBuilderIface struct {
	ctrl     *gomock.Controller
	recorder *MockArtifactBuilderIfaceMockRecorder
}

// MockArtifactBuilderIfaceMockRecorder is the mock recorder for MockArtifactBuilderIface.
type MockArtifactBuilderIfaceMockRecorder struct {
	mock *MockArtifactBuilderIface
}

// NewMockArtifactBuilderIface creates a new mock instance.
func NewMockArtifactBuilderIface(ctrl *gomock.Controller) *MockArtifactBuilderIface {
	mock := &MockArtifactBuilderIface{ctrl: ctrl}
	mock.recorder = &MockArtifactBuilderIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArtifactBuilderIface) EXPECT() *MockArtifactBuilderIfaceMockRecorder {
	return m.recorder
}

// Build mocks base method.
func (m *MockArtifactBuilderIface) Build(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Build", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Build indicates an expected call of Build.
func (mr *MockArtifactBuilderIfaceMockRecorder) Build(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockArtifactBuilderIface)(nil).Build), arg0)
}

// BuildAndPush mocks base method.
func (m *MockArtifactBuilderIface) BuildAndPush(arg0 context.Context, arg1 ...artifact_builder.ArtifactBuilderBuildOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BuildAndPush", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// BuildAndPush indicates an expected call of BuildAndPush.
func (mr *MockArtifactBuilderIfaceMockRecorder) BuildAndPush(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildAndPush", reflect.TypeOf((*MockArtifactBuilderIface)(nil).BuildAndPush), varargs...)
}

// CheckImageExists mocks base method.
func (m *MockArtifactBuilderIface) CheckImageExists(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckImageExists", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckImageExists indicates an expected call of CheckImageExists.
func (mr *MockArtifactBuilderIfaceMockRecorder) CheckImageExists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckImageExists", reflect.TypeOf((*MockArtifactBuilderIface)(nil).CheckImageExists), arg0, arg1)
}

// Push mocks base method.
func (m *MockArtifactBuilderIface) Push(arg0 context.Context, arg1 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Push", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Push indicates an expected call of Push.
func (mr *MockArtifactBuilderIfaceMockRecorder) Push(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Push", reflect.TypeOf((*MockArtifactBuilderIface)(nil).Push), arg0, arg1)
}

// RegistryLogin mocks base method.
func (m *MockArtifactBuilderIface) RegistryLogin(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegistryLogin", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegistryLogin indicates an expected call of RegistryLogin.
func (mr *MockArtifactBuilderIfaceMockRecorder) RegistryLogin(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegistryLogin", reflect.TypeOf((*MockArtifactBuilderIface)(nil).RegistryLogin), arg0)
}

// RetagImages mocks base method.
func (m *MockArtifactBuilderIface) RetagImages(arg0 context.Context, arg1 map[string]*config.RegistryConfig, arg2 string, arg3, arg4 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetagImages", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// RetagImages indicates an expected call of RetagImages.
func (mr *MockArtifactBuilderIfaceMockRecorder) RetagImages(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetagImages", reflect.TypeOf((*MockArtifactBuilderIface)(nil).RetagImages), arg0, arg1, arg2, arg3, arg4)
}

// WithBackend mocks base method.
func (m *MockArtifactBuilderIface) WithBackend(arg0 *aws.Backend) artifact_builder.ArtifactBuilderIface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithBackend", arg0)
	ret0, _ := ret[0].(artifact_builder.ArtifactBuilderIface)
	return ret0
}

// WithBackend indicates an expected call of WithBackend.
func (mr *MockArtifactBuilderIfaceMockRecorder) WithBackend(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithBackend", reflect.TypeOf((*MockArtifactBuilderIface)(nil).WithBackend), arg0)
}

// WithConfig mocks base method.
func (m *MockArtifactBuilderIface) WithConfig(arg0 *artifact_builder.BuilderConfig) artifact_builder.ArtifactBuilderIface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithConfig", arg0)
	ret0, _ := ret[0].(artifact_builder.ArtifactBuilderIface)
	return ret0
}

// WithConfig indicates an expected call of WithConfig.
func (mr *MockArtifactBuilderIfaceMockRecorder) WithConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithConfig", reflect.TypeOf((*MockArtifactBuilderIface)(nil).WithConfig), arg0)
}

// WithTags mocks base method.
func (m *MockArtifactBuilderIface) WithTags(arg0 []string) artifact_builder.ArtifactBuilderIface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithTags", arg0)
	ret0, _ := ret[0].(artifact_builder.ArtifactBuilderIface)
	return ret0
}

// WithTags indicates an expected call of WithTags.
func (mr *MockArtifactBuilderIfaceMockRecorder) WithTags(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithTags", reflect.TypeOf((*MockArtifactBuilderIface)(nil).WithTags), arg0)
}
