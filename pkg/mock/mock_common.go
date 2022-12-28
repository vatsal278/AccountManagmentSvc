// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/vatsal278/AccountManagmentSvc/internal/handler (interfaces: Commoner,HealthChecker)

// Package mock is a generated GoMock package.
package mock

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCommoner is a mock of Commoner interface.
type MockCommoner struct {
	ctrl     *gomock.Controller
	recorder *MockCommonerMockRecorder
}

// MockCommonerMockRecorder is the mock recorder for MockCommoner.
type MockCommonerMockRecorder struct {
	mock *MockCommoner
}

// NewMockCommoner creates a new mock instance.
func NewMockCommoner(ctrl *gomock.Controller) *MockCommoner {
	mock := &MockCommoner{ctrl: ctrl}
	mock.recorder = &MockCommonerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommoner) EXPECT() *MockCommonerMockRecorder {
	return m.recorder
}

// HealthCheck mocks base method.
func (m *MockCommoner) HealthCheck(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HealthCheck", arg0, arg1)
}

// HealthCheck indicates an expected call of HealthCheck.
func (mr *MockCommonerMockRecorder) HealthCheck(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockCommoner)(nil).HealthCheck), arg0, arg1)
}

// MethodNotAllowed mocks base method.
func (m *MockCommoner) MethodNotAllowed(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "MethodNotAllowed", arg0, arg1)
}

// MethodNotAllowed indicates an expected call of MethodNotAllowed.
func (mr *MockCommonerMockRecorder) MethodNotAllowed(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MethodNotAllowed", reflect.TypeOf((*MockCommoner)(nil).MethodNotAllowed), arg0, arg1)
}

// RouteNotFound mocks base method.
func (m *MockCommoner) RouteNotFound(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RouteNotFound", arg0, arg1)
}

// RouteNotFound indicates an expected call of RouteNotFound.
func (mr *MockCommonerMockRecorder) RouteNotFound(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RouteNotFound", reflect.TypeOf((*MockCommoner)(nil).RouteNotFound), arg0, arg1)
}

// MockHealthChecker is a mock of HealthChecker interface.
type MockHealthChecker struct {
	ctrl     *gomock.Controller
	recorder *MockHealthCheckerMockRecorder
}

// MockHealthCheckerMockRecorder is the mock recorder for MockHealthChecker.
type MockHealthCheckerMockRecorder struct {
	mock *MockHealthChecker
}

// NewMockHealthChecker creates a new mock instance.
func NewMockHealthChecker(ctrl *gomock.Controller) *MockHealthChecker {
	mock := &MockHealthChecker{ctrl: ctrl}
	mock.recorder = &MockHealthCheckerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHealthChecker) EXPECT() *MockHealthCheckerMockRecorder {
	return m.recorder
}

// HealthCheck mocks base method.
func (m *MockHealthChecker) HealthCheck() (string, string, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HealthCheck")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(bool)
	return ret0, ret1, ret2
}

// HealthCheck indicates an expected call of HealthCheck.
func (mr *MockHealthCheckerMockRecorder) HealthCheck() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockHealthChecker)(nil).HealthCheck))
}
