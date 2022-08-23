// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/chanzuckerberg/happy/pkg/backend/aws/interfaces (interfaces: GetLogEventsAPIClient)

// Package interfaces is a generated GoMock package.
package interfaces

import (
	context "context"
	reflect "reflect"

	cloudwatchlogs "github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	gomock "github.com/golang/mock/gomock"
)

// MockGetLogEventsAPIClient is a mock of GetLogEventsAPIClient interface.
type MockGetLogEventsAPIClient struct {
	ctrl     *gomock.Controller
	recorder *MockGetLogEventsAPIClientMockRecorder
}

// MockGetLogEventsAPIClientMockRecorder is the mock recorder for MockGetLogEventsAPIClient.
type MockGetLogEventsAPIClientMockRecorder struct {
	mock *MockGetLogEventsAPIClient
}

// NewMockGetLogEventsAPIClient creates a new mock instance.
func NewMockGetLogEventsAPIClient(ctrl *gomock.Controller) *MockGetLogEventsAPIClient {
	mock := &MockGetLogEventsAPIClient{ctrl: ctrl}
	mock.recorder = &MockGetLogEventsAPIClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGetLogEventsAPIClient) EXPECT() *MockGetLogEventsAPIClientMockRecorder {
	return m.recorder
}

// DescribeLogStreams mocks base method.
func (m *MockGetLogEventsAPIClient) DescribeLogStreams(arg0 context.Context, arg1 *cloudwatchlogs.DescribeLogStreamsInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeLogStreamsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeLogStreams", varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.DescribeLogStreamsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeLogStreams indicates an expected call of DescribeLogStreams.
func (mr *MockGetLogEventsAPIClientMockRecorder) DescribeLogStreams(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeLogStreams", reflect.TypeOf((*MockGetLogEventsAPIClient)(nil).DescribeLogStreams), varargs...)
}

// GetLogEvents mocks base method.
func (m *MockGetLogEventsAPIClient) GetLogEvents(arg0 context.Context, arg1 *cloudwatchlogs.GetLogEventsInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.GetLogEventsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLogEvents", varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.GetLogEventsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLogEvents indicates an expected call of GetLogEvents.
func (mr *MockGetLogEventsAPIClientMockRecorder) GetLogEvents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogEvents", reflect.TypeOf((*MockGetLogEventsAPIClient)(nil).GetLogEvents), varargs...)
}