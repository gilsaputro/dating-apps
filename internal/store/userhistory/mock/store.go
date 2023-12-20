// Code generated by MockGen. DO NOT EDIT.
// Source: internal/store/userhistory/store.go

// Package mock is a generated GoMock package.
package mock

import (
	models "gilsaputro/dating-apps/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserHistoryStoreMethod is a mock of UserHistoryStoreMethod interface.
type MockUserHistoryStoreMethod struct {
	ctrl     *gomock.Controller
	recorder *MockUserHistoryStoreMethodMockRecorder
}

// MockUserHistoryStoreMethodMockRecorder is the mock recorder for MockUserHistoryStoreMethod.
type MockUserHistoryStoreMethodMockRecorder struct {
	mock *MockUserHistoryStoreMethod
}

// NewMockUserHistoryStoreMethod creates a new mock instance.
func NewMockUserHistoryStoreMethod(ctrl *gomock.Controller) *MockUserHistoryStoreMethod {
	mock := &MockUserHistoryStoreMethod{ctrl: ctrl}
	mock.recorder = &MockUserHistoryStoreMethodMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserHistoryStoreMethod) EXPECT() *MockUserHistoryStoreMethodMockRecorder {
	return m.recorder
}

// CountByUserIDAndPartnerID mocks base method.
func (m *MockUserHistoryStoreMethod) CountByUserIDAndPartnerID(userID, partnerID int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountByUserIDAndPartnerID", userID, partnerID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountByUserIDAndPartnerID indicates an expected call of CountByUserIDAndPartnerID.
func (mr *MockUserHistoryStoreMethodMockRecorder) CountByUserIDAndPartnerID(userID, partnerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountByUserIDAndPartnerID", reflect.TypeOf((*MockUserHistoryStoreMethod)(nil).CountByUserIDAndPartnerID), userID, partnerID)
}

// CreateUserHistory mocks base method.
func (m *MockUserHistoryStoreMethod) CreateUserHistory(hist models.UserMatchHistory) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserHistory", hist)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUserHistory indicates an expected call of CreateUserHistory.
func (mr *MockUserHistoryStoreMethodMockRecorder) CreateUserHistory(hist interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserHistory", reflect.TypeOf((*MockUserHistoryStoreMethod)(nil).CreateUserHistory), hist)
}

// GetUserHistoryListByUserID mocks base method.
func (m *MockUserHistoryStoreMethod) GetUserHistoryListByUserID(hist models.UserMatchHistory) ([]models.UserMatchHistory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserHistoryListByUserID", hist)
	ret0, _ := ret[0].([]models.UserMatchHistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserHistoryListByUserID indicates an expected call of GetUserHistoryListByUserID.
func (mr *MockUserHistoryStoreMethodMockRecorder) GetUserHistoryListByUserID(hist interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserHistoryListByUserID", reflect.TypeOf((*MockUserHistoryStoreMethod)(nil).GetUserHistoryListByUserID), hist)
}

// UpdatePartnerStatus mocks base method.
func (m *MockUserHistoryStoreMethod) UpdatePartnerStatus(history models.UserMatchHistory) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePartnerStatus", history)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePartnerStatus indicates an expected call of UpdatePartnerStatus.
func (mr *MockUserHistoryStoreMethodMockRecorder) UpdatePartnerStatus(history interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePartnerStatus", reflect.TypeOf((*MockUserHistoryStoreMethod)(nil).UpdatePartnerStatus), history)
}