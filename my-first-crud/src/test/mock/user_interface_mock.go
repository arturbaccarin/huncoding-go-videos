// Code generated by MockGen. DO NOT EDIT.
// Source: src/model/service/user_interface.go
//
// Generated by this command:
//
//	mockgen -source=src/model/service/user_interface.go -destination=src/test/mock/user_interface_mock.go -package=mock
//

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	rest_err "github.com/arturbaccarin/go-my-first-crud/src/configuration/rest_err"
	model "github.com/arturbaccarin/go-my-first-crud/src/model"
	gomock "go.uber.org/mock/gomock"
)

// MockUserDomainService is a mock of UserDomainService interface.
type MockUserDomainService struct {
	ctrl     *gomock.Controller
	recorder *MockUserDomainServiceMockRecorder
}

// MockUserDomainServiceMockRecorder is the mock recorder for MockUserDomainService.
type MockUserDomainServiceMockRecorder struct {
	mock *MockUserDomainService
}

// NewMockUserDomainService creates a new mock instance.
func NewMockUserDomainService(ctrl *gomock.Controller) *MockUserDomainService {
	mock := &MockUserDomainService{ctrl: ctrl}
	mock.recorder = &MockUserDomainServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserDomainService) EXPECT() *MockUserDomainServiceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserDomainService) CreateUser(arg0 model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0)
	ret0, _ := ret[0].(model.UserDomainInterface)
	ret1, _ := ret[1].(*rest_err.RestErr)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserDomainServiceMockRecorder) CreateUser(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserDomainService)(nil).CreateUser), arg0)
}

// DeleteUser mocks base method.
func (m *MockUserDomainService) DeleteUser(arg0 string) *rest_err.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0)
	ret0, _ := ret[0].(*rest_err.RestErr)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserDomainServiceMockRecorder) DeleteUser(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserDomainService)(nil).DeleteUser), arg0)
}

// FindUserByEmail mocks base method.
func (m *MockUserDomainService) FindUserByEmail(arg0 string) (model.UserDomainInterface, *rest_err.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByEmail", arg0)
	ret0, _ := ret[0].(model.UserDomainInterface)
	ret1, _ := ret[1].(*rest_err.RestErr)
	return ret0, ret1
}

// FindUserByEmail indicates an expected call of FindUserByEmail.
func (mr *MockUserDomainServiceMockRecorder) FindUserByEmail(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmail", reflect.TypeOf((*MockUserDomainService)(nil).FindUserByEmail), arg0)
}

// FindUserById mocks base method.
func (m *MockUserDomainService) FindUserById(arg0 string) (model.UserDomainInterface, *rest_err.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserById", arg0)
	ret0, _ := ret[0].(model.UserDomainInterface)
	ret1, _ := ret[1].(*rest_err.RestErr)
	return ret0, ret1
}

// FindUserById indicates an expected call of FindUserById.
func (mr *MockUserDomainServiceMockRecorder) FindUserById(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserById", reflect.TypeOf((*MockUserDomainService)(nil).FindUserById), arg0)
}

// LoginUser mocks base method.
func (m *MockUserDomainService) LoginUser(arg0 model.UserDomainInterface) (model.UserDomainInterface, string, *rest_err.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginUser", arg0)
	ret0, _ := ret[0].(model.UserDomainInterface)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(*rest_err.RestErr)
	return ret0, ret1, ret2
}

// LoginUser indicates an expected call of LoginUser.
func (mr *MockUserDomainServiceMockRecorder) LoginUser(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginUser", reflect.TypeOf((*MockUserDomainService)(nil).LoginUser), arg0)
}

// UpdateUser mocks base method.
func (m *MockUserDomainService) UpdateUser(arg0 string, arg1 model.UserDomainInterface) *rest_err.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(*rest_err.RestErr)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserDomainServiceMockRecorder) UpdateUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserDomainService)(nil).UpdateUser), arg0, arg1)
}
