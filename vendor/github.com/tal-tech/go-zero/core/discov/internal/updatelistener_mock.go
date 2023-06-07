// Code generated by MockGen. DO NOT EDIT.
// Source: updatelistener.go

// Package internal is a generated GoMock package.
package internal

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUpdateListener is a mock of UpdateListener interface
type MockUpdateListener struct {
	ctrl     *gomock.Controller
	recorder *MockUpdateListenerMockRecorder
}

// MockUpdateListenerMockRecorder is the mock recorder for MockUpdateListener
type MockUpdateListenerMockRecorder struct {
	mock *MockUpdateListener
}

// NewMockUpdateListener creates a new mock instance
func NewMockUpdateListener(ctrl *gomock.Controller) *MockUpdateListener {
	mock := &MockUpdateListener{ctrl: ctrl}
	mock.recorder = &MockUpdateListenerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUpdateListener) EXPECT() *MockUpdateListenerMockRecorder {
	return m.recorder
}

// OnAdd mocks base method
func (m *MockUpdateListener) OnAdd(kv KV) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnAdd", kv)
}

// OnAdd indicates an expected call of OnAdd
func (mr *MockUpdateListenerMockRecorder) OnAdd(kv interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnAdd", reflect.TypeOf((*MockUpdateListener)(nil).OnAdd), kv)
}

// OnDelete mocks base method
func (m *MockUpdateListener) OnDelete(kv KV) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnDelete", kv)
}

// OnDelete indicates an expected call of OnDelete
func (mr *MockUpdateListenerMockRecorder) OnDelete(kv interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnDelete", reflect.TypeOf((*MockUpdateListener)(nil).OnDelete), kv)
}
