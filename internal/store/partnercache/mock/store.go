// Code generated by MockGen. DO NOT EDIT.
// Source: internal/store/partnercache/store.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPartnerCacheStoreMethod is a mock of PartnerCacheStoreMethod interface.
type MockPartnerCacheStoreMethod struct {
	ctrl     *gomock.Controller
	recorder *MockPartnerCacheStoreMethodMockRecorder
}

// MockPartnerCacheStoreMethodMockRecorder is the mock recorder for MockPartnerCacheStoreMethod.
type MockPartnerCacheStoreMethodMockRecorder struct {
	mock *MockPartnerCacheStoreMethod
}

// NewMockPartnerCacheStoreMethod creates a new mock instance.
func NewMockPartnerCacheStoreMethod(ctrl *gomock.Controller) *MockPartnerCacheStoreMethod {
	mock := &MockPartnerCacheStoreMethod{ctrl: ctrl}
	mock.recorder = &MockPartnerCacheStoreMethodMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPartnerCacheStoreMethod) EXPECT() *MockPartnerCacheStoreMethodMockRecorder {
	return m.recorder
}

// GetCurentPartnerState mocks base method.
func (m *MockPartnerCacheStoreMethod) GetCurentPartnerState(userID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurentPartnerState", userID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurentPartnerState indicates an expected call of GetCurentPartnerState.
func (mr *MockPartnerCacheStoreMethodMockRecorder) GetCurentPartnerState(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurentPartnerState", reflect.TypeOf((*MockPartnerCacheStoreMethod)(nil).GetCurentPartnerState), userID)
}

// GetViewedPartnerHistory mocks base method.
func (m *MockPartnerCacheStoreMethod) GetViewedPartnerHistory(userID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetViewedPartnerHistory", userID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetViewedPartnerHistory indicates an expected call of GetViewedPartnerHistory.
func (mr *MockPartnerCacheStoreMethodMockRecorder) GetViewedPartnerHistory(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetViewedPartnerHistory", reflect.TypeOf((*MockPartnerCacheStoreMethod)(nil).GetViewedPartnerHistory), userID)
}

// GetViewedUserCounter mocks base method.
func (m *MockPartnerCacheStoreMethod) GetViewedUserCounter(userID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetViewedUserCounter", userID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetViewedUserCounter indicates an expected call of GetViewedUserCounter.
func (mr *MockPartnerCacheStoreMethodMockRecorder) GetViewedUserCounter(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetViewedUserCounter", reflect.TypeOf((*MockPartnerCacheStoreMethod)(nil).GetViewedUserCounter), userID)
}

// SetCurentPartnerState mocks base method.
func (m *MockPartnerCacheStoreMethod) SetCurentPartnerState(userID, partnerID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCurentPartnerState", userID, partnerID)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetCurentPartnerState indicates an expected call of SetCurentPartnerState.
func (mr *MockPartnerCacheStoreMethodMockRecorder) SetCurentPartnerState(userID, partnerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCurentPartnerState", reflect.TypeOf((*MockPartnerCacheStoreMethod)(nil).SetCurentPartnerState), userID, partnerID)
}

// SetViewedPartnerHistory mocks base method.
func (m *MockPartnerCacheStoreMethod) SetViewedPartnerHistory(userID, value string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetViewedPartnerHistory", userID, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetViewedPartnerHistory indicates an expected call of SetViewedPartnerHistory.
func (mr *MockPartnerCacheStoreMethodMockRecorder) SetViewedPartnerHistory(userID, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetViewedPartnerHistory", reflect.TypeOf((*MockPartnerCacheStoreMethod)(nil).SetViewedPartnerHistory), userID, value)
}

// SetViewedUserCounter mocks base method.
func (m *MockPartnerCacheStoreMethod) SetViewedUserCounter(userID, value string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetViewedUserCounter", userID, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetViewedUserCounter indicates an expected call of SetViewedUserCounter.
func (mr *MockPartnerCacheStoreMethodMockRecorder) SetViewedUserCounter(userID, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetViewedUserCounter", reflect.TypeOf((*MockPartnerCacheStoreMethod)(nil).SetViewedUserCounter), userID, value)
}