// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/chanzuckerberg/happy/cli/pkg/backend/aws/interfaces (interfaces: ECRAPI)

// Package interfaces is a generated GoMock package.
package interfaces

import (
	context "context"
	reflect "reflect"

	ecr "github.com/aws/aws-sdk-go-v2/service/ecr"
	gomock "github.com/golang/mock/gomock"
)

// MockECRAPI is a mock of ECRAPI interface.
type MockECRAPI struct {
	ctrl     *gomock.Controller
	recorder *MockECRAPIMockRecorder
}

// MockECRAPIMockRecorder is the mock recorder for MockECRAPI.
type MockECRAPIMockRecorder struct {
	mock *MockECRAPI
}

// NewMockECRAPI creates a new mock instance.
func NewMockECRAPI(ctrl *gomock.Controller) *MockECRAPI {
	mock := &MockECRAPI{ctrl: ctrl}
	mock.recorder = &MockECRAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockECRAPI) EXPECT() *MockECRAPIMockRecorder {
	return m.recorder
}

// BatchGetImage mocks base method.
func (m *MockECRAPI) BatchGetImage(arg0 context.Context, arg1 *ecr.BatchGetImageInput, arg2 ...func(*ecr.Options)) (*ecr.BatchGetImageOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchGetImage", varargs...)
	ret0, _ := ret[0].(*ecr.BatchGetImageOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchGetImage indicates an expected call of BatchGetImage.
func (mr *MockECRAPIMockRecorder) BatchGetImage(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchGetImage", reflect.TypeOf((*MockECRAPI)(nil).BatchGetImage), varargs...)
}

// GetAuthorizationToken mocks base method.
func (m *MockECRAPI) GetAuthorizationToken(arg0 context.Context, arg1 *ecr.GetAuthorizationTokenInput, arg2 ...func(*ecr.Options)) (*ecr.GetAuthorizationTokenOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAuthorizationToken", varargs...)
	ret0, _ := ret[0].(*ecr.GetAuthorizationTokenOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuthorizationToken indicates an expected call of GetAuthorizationToken.
func (mr *MockECRAPIMockRecorder) GetAuthorizationToken(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthorizationToken", reflect.TypeOf((*MockECRAPI)(nil).GetAuthorizationToken), varargs...)
}

// PutImage mocks base method.
func (m *MockECRAPI) PutImage(arg0 context.Context, arg1 *ecr.PutImageInput, arg2 ...func(*ecr.Options)) (*ecr.PutImageOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutImage", varargs...)
	ret0, _ := ret[0].(*ecr.PutImageOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutImage indicates an expected call of PutImage.
func (mr *MockECRAPIMockRecorder) PutImage(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutImage", reflect.TypeOf((*MockECRAPI)(nil).PutImage), varargs...)
}
