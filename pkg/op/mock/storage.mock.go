// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/caos/oidc/pkg/op/u (interfaces: Storage)

// Package mock is a generated GoMock package.
package mock

import (
	oidc "github.com/caos/oidc/pkg/oidc"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockStorage is a mock of Storage interface
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// AuthRequestByCode mocks base method
func (m *MockStorage) AuthRequestByCode(arg0 oidc.Client, arg1, arg2 string) (*oidc.AuthRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthRequestByCode", arg0, arg1, arg2)
	ret0, _ := ret[0].(*oidc.AuthRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthRequestByCode indicates an expected call of AuthRequestByCode
func (mr *MockStorageMockRecorder) AuthRequestByCode(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthRequestByCode", reflect.TypeOf((*MockStorage)(nil).AuthRequestByCode), arg0, arg1, arg2)
}

// AuthRequestByID mocks base method
func (m *MockStorage) AuthRequestByID(arg0 string) (*oidc.AuthRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthRequestByID", arg0)
	ret0, _ := ret[0].(*oidc.AuthRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthRequestByID indicates an expected call of AuthRequestByID
func (mr *MockStorageMockRecorder) AuthRequestByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthRequestByID", reflect.TypeOf((*MockStorage)(nil).AuthRequestByID), arg0)
}

// AuthorizeClientIDCodeVerifier mocks base method
func (m *MockStorage) AuthorizeClientIDCodeVerifier(arg0, arg1 string) (oidc.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthorizeClientIDCodeVerifier", arg0, arg1)
	ret0, _ := ret[0].(oidc.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthorizeClientIDCodeVerifier indicates an expected call of AuthorizeClientIDCodeVerifier
func (mr *MockStorageMockRecorder) AuthorizeClientIDCodeVerifier(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthorizeClientIDCodeVerifier", reflect.TypeOf((*MockStorage)(nil).AuthorizeClientIDCodeVerifier), arg0, arg1)
}

// AuthorizeClientIDSecret mocks base method
func (m *MockStorage) AuthorizeClientIDSecret(arg0, arg1 string) (oidc.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthorizeClientIDSecret", arg0, arg1)
	ret0, _ := ret[0].(oidc.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthorizeClientIDSecret indicates an expected call of AuthorizeClientIDSecret
func (mr *MockStorageMockRecorder) AuthorizeClientIDSecret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthorizeClientIDSecret", reflect.TypeOf((*MockStorage)(nil).AuthorizeClientIDSecret), arg0, arg1)
}

// CreateAuthRequest mocks base method
func (m *MockStorage) CreateAuthRequest(arg0 *oidc.AuthRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAuthRequest", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAuthRequest indicates an expected call of CreateAuthRequest
func (mr *MockStorageMockRecorder) CreateAuthRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAuthRequest", reflect.TypeOf((*MockStorage)(nil).CreateAuthRequest), arg0)
}

// DeleteAuthRequestAndCode mocks base method
func (m *MockStorage) DeleteAuthRequestAndCode(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAuthRequestAndCode", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAuthRequestAndCode indicates an expected call of DeleteAuthRequestAndCode
func (mr *MockStorageMockRecorder) DeleteAuthRequestAndCode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAuthRequestAndCode", reflect.TypeOf((*MockStorage)(nil).DeleteAuthRequestAndCode), arg0, arg1)
}

// GetClientByClientID mocks base method
func (m *MockStorage) GetClientByClientID(arg0 string) (oidc.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClientByClientID", arg0)
	ret0, _ := ret[0].(oidc.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClientByClientID indicates an expected call of GetClientByClientID
func (mr *MockStorageMockRecorder) GetClientByClientID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClientByClientID", reflect.TypeOf((*MockStorage)(nil).GetClientByClientID), arg0)
}
