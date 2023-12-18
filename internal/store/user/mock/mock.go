// Code generated by MockGen. DO NOT EDIT.
// Source: C:\Users\gilsp\go\src\dating-apps\internal\store\user\user.go

// Package mock_user is a generated GoMock package.
package mock_user

import (
	user "gilsaputro/dating-apps/internal/store/user"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserStoreMethod is a mock of UserStoreMethod interface.
type MockUserStoreMethod struct {
	ctrl     *gomock.Controller
	recorder *MockUserStoreMethodMockRecorder
}

// MockUserStoreMethodMockRecorder is the mock recorder for MockUserStoreMethod.
type MockUserStoreMethodMockRecorder struct {
	mock *MockUserStoreMethod
}

// NewMockUserStoreMethod creates a new mock instance.
func NewMockUserStoreMethod(ctrl *gomock.Controller) *MockUserStoreMethod {
	mock := &MockUserStoreMethod{ctrl: ctrl}
	mock.recorder = &MockUserStoreMethodMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserStoreMethod) EXPECT() *MockUserStoreMethodMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserStoreMethod) CreateUser(userinfo user.UserStoreInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", userinfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserStoreMethodMockRecorder) CreateUser(userinfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserStoreMethod)(nil).CreateUser), userinfo)
}

// DeleteUser mocks base method.
func (m *MockUserStoreMethod) DeleteUser(userid int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", userid)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserStoreMethodMockRecorder) DeleteUser(userid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserStoreMethod)(nil).DeleteUser), userid)
}

// GetAllUserInfoWithPagging mocks base method.
func (m *MockUserStoreMethod) GetAllUserInfoWithPagging(size, cursor int) ([]user.UserStoreInfo, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUserInfoWithPagging", size, cursor)
	ret0, _ := ret[0].([]user.UserStoreInfo)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAllUserInfoWithPagging indicates an expected call of GetAllUserInfoWithPagging.
func (mr *MockUserStoreMethodMockRecorder) GetAllUserInfoWithPagging(size, cursor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUserInfoWithPagging", reflect.TypeOf((*MockUserStoreMethod)(nil).GetAllUserInfoWithPagging), size, cursor)
}

// GetUserInfoByID mocks base method.
func (m *MockUserStoreMethod) GetUserInfoByID(userid int) (user.UserStoreInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserInfoByID", userid)
	ret0, _ := ret[0].(user.UserStoreInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserInfoByID indicates an expected call of GetUserInfoByID.
func (mr *MockUserStoreMethodMockRecorder) GetUserInfoByID(userid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserInfoByID", reflect.TypeOf((*MockUserStoreMethod)(nil).GetUserInfoByID), userid)
}

// GetUserInfoByUsername mocks base method.
func (m *MockUserStoreMethod) GetUserInfoByUsername(username string) (user.UserStoreInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserInfoByUsername", username)
	ret0, _ := ret[0].(user.UserStoreInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserInfoByUsername indicates an expected call of GetUserInfoByUsername.
func (mr *MockUserStoreMethodMockRecorder) GetUserInfoByUsername(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserInfoByUsername", reflect.TypeOf((*MockUserStoreMethod)(nil).GetUserInfoByUsername), username)
}

// UpdateUser mocks base method.
func (m *MockUserStoreMethod) UpdateUser(userinfo user.UserStoreInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", userinfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserStoreMethodMockRecorder) UpdateUser(userinfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserStoreMethod)(nil).UpdateUser), userinfo)
}
