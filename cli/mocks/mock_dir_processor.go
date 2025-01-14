// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/chanzuckerberg/happy/cli/pkg/util (interfaces: DirProcessor)

// Package mocks is a generated GoMock package.
package mocks

import (
	os "os"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDirProcessor is a mock of DirProcessor interface.
type MockDirProcessor struct {
	ctrl     *gomock.Controller
	recorder *MockDirProcessorMockRecorder
}

// MockDirProcessorMockRecorder is the mock recorder for MockDirProcessor.
type MockDirProcessorMockRecorder struct {
	mock *MockDirProcessor
}

// NewMockDirProcessor creates a new mock instance.
func NewMockDirProcessor(ctrl *gomock.Controller) *MockDirProcessor {
	mock := &MockDirProcessor{ctrl: ctrl}
	mock.recorder = &MockDirProcessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDirProcessor) EXPECT() *MockDirProcessorMockRecorder {
	return m.recorder
}

// Tarzip mocks base method.
func (m *MockDirProcessor) Tarzip(arg0 string, arg1 *os.File) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tarzip", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Tarzip indicates an expected call of Tarzip.
func (mr *MockDirProcessorMockRecorder) Tarzip(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tarzip", reflect.TypeOf((*MockDirProcessor)(nil).Tarzip), arg0, arg1)
}
