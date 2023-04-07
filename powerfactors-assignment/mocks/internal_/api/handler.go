// Code generated by MockGen. DO NOT EDIT.
// Source: internal/api/handler.go

// Package mock_api is a generated GoMock package.
package mock_api

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTimestampHandlerInt is a mock of TimestampHandlerInt interface.
type MockTimestampHandlerInt struct {
	ctrl     *gomock.Controller
	recorder *MockTimestampHandlerIntMockRecorder
}

// MockTimestampHandlerIntMockRecorder is the mock recorder for MockTimestampHandlerInt.
type MockTimestampHandlerIntMockRecorder struct {
	mock *MockTimestampHandlerInt
}

// NewMockTimestampHandlerInt creates a new mock instance.
func NewMockTimestampHandlerInt(ctrl *gomock.Controller) *MockTimestampHandlerInt {
	mock := &MockTimestampHandlerInt{ctrl: ctrl}
	mock.recorder = &MockTimestampHandlerIntMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTimestampHandlerInt) EXPECT() *MockTimestampHandlerIntMockRecorder {
	return m.recorder
}

// GetTimestamp mocks base method.
func (m *MockTimestampHandlerInt) GetTimestamp(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetTimestamp", w, r)
}

// GetTimestamp indicates an expected call of GetTimestamp.
func (mr *MockTimestampHandlerIntMockRecorder) GetTimestamp(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTimestamp", reflect.TypeOf((*MockTimestampHandlerInt)(nil).GetTimestamp), w, r)
}
