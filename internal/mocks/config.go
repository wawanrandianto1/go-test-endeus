// Code generated by MockGen. DO NOT EDIT.
// Source: config.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockConfig is a mock of Config interface.
type MockConfig struct {
	ctrl     *gomock.Controller
	recorder *MockConfigMockRecorder
}

// MockConfigMockRecorder is the mock recorder for MockConfig.
type MockConfigMockRecorder struct {
	mock *MockConfig
}

// NewMockConfig creates a new mock instance.
func NewMockConfig(ctrl *gomock.Controller) *MockConfig {
	mock := &MockConfig{ctrl: ctrl}
	mock.recorder = &MockConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfig) EXPECT() *MockConfigMockRecorder {
	return m.recorder
}

// DBAddress mocks base method.
func (m *MockConfig) DBAddress() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DBAddress")
	ret0, _ := ret[0].(string)
	return ret0
}

// DBAddress indicates an expected call of DBAddress.
func (mr *MockConfigMockRecorder) DBAddress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DBAddress", reflect.TypeOf((*MockConfig)(nil).DBAddress))
}

// DBHost mocks base method.
func (m *MockConfig) DBHost() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DBHost")
	ret0, _ := ret[0].(string)
	return ret0
}

// DBHost indicates an expected call of DBHost.
func (mr *MockConfigMockRecorder) DBHost() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DBHost", reflect.TypeOf((*MockConfig)(nil).DBHost))
}

// DBName mocks base method.
func (m *MockConfig) DBName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DBName")
	ret0, _ := ret[0].(string)
	return ret0
}

// DBName indicates an expected call of DBName.
func (mr *MockConfigMockRecorder) DBName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DBName", reflect.TypeOf((*MockConfig)(nil).DBName))
}

// DBPassword mocks base method.
func (m *MockConfig) DBPassword() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DBPassword")
	ret0, _ := ret[0].(string)
	return ret0
}

// DBPassword indicates an expected call of DBPassword.
func (mr *MockConfigMockRecorder) DBPassword() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DBPassword", reflect.TypeOf((*MockConfig)(nil).DBPassword))
}

// DBPort mocks base method.
func (m *MockConfig) DBPort() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DBPort")
	ret0, _ := ret[0].(string)
	return ret0
}

// DBPort indicates an expected call of DBPort.
func (mr *MockConfigMockRecorder) DBPort() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DBPort", reflect.TypeOf((*MockConfig)(nil).DBPort))
}

// DBUsername mocks base method.
func (m *MockConfig) DBUsername() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DBUsername")
	ret0, _ := ret[0].(string)
	return ret0
}

// DBUsername indicates an expected call of DBUsername.
func (mr *MockConfigMockRecorder) DBUsername() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DBUsername", reflect.TypeOf((*MockConfig)(nil).DBUsername))
}

// Environment mocks base method.
func (m *MockConfig) Environment() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Environment")
	ret0, _ := ret[0].(string)
	return ret0
}

// Environment indicates an expected call of Environment.
func (mr *MockConfigMockRecorder) Environment() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Environment", reflect.TypeOf((*MockConfig)(nil).Environment))
}

// IsDev mocks base method.
func (m *MockConfig) IsDev() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDev")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsDev indicates an expected call of IsDev.
func (mr *MockConfigMockRecorder) IsDev() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDev", reflect.TypeOf((*MockConfig)(nil).IsDev))
}

// LogLevel mocks base method.
func (m *MockConfig) LogLevel() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LogLevel")
	ret0, _ := ret[0].(string)
	return ret0
}

// LogLevel indicates an expected call of LogLevel.
func (mr *MockConfigMockRecorder) LogLevel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogLevel", reflect.TypeOf((*MockConfig)(nil).LogLevel))
}

// ServerPort mocks base method.
func (m *MockConfig) ServerPort() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerPort")
	ret0, _ := ret[0].(string)
	return ret0
}

// ServerPort indicates an expected call of ServerPort.
func (mr *MockConfigMockRecorder) ServerPort() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerPort", reflect.TypeOf((*MockConfig)(nil).ServerPort))
}
