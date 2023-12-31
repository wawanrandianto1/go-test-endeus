// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	endeus "endeus/wawan/pkg/endeus"
	model "endeus/wawan/pkg/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockResepService is a mock of ResepService interface.
type MockResepService struct {
	ctrl     *gomock.Controller
	recorder *MockResepServiceMockRecorder
}

// MockResepServiceMockRecorder is the mock recorder for MockResepService.
type MockResepServiceMockRecorder struct {
	mock *MockResepService
}

// NewMockResepService creates a new mock instance.
func NewMockResepService(ctrl *gomock.Controller) *MockResepService {
	mock := &MockResepService{ctrl: ctrl}
	mock.recorder = &MockResepServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResepService) EXPECT() *MockResepServiceMockRecorder {
	return m.recorder
}

// CreateNew mocks base method.
func (m *MockResepService) CreateNew(param endeus.ResepParam) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNew", param)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateNew indicates an expected call of CreateNew.
func (mr *MockResepServiceMockRecorder) CreateNew(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNew", reflect.TypeOf((*MockResepService)(nil).CreateNew), param)
}

// GetAll mocks base method.
func (m *MockResepService) GetAll(str string) ([]endeus.ResepCustomResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", str)
	ret0, _ := ret[0].([]endeus.ResepCustomResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockResepServiceMockRecorder) GetAll(str interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockResepService)(nil).GetAll), str)
}

// GetAllByCategoryID mocks base method.
func (m *MockResepService) GetAllByCategoryID(catID uint) ([]endeus.ResepCustomResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllByCategoryID", catID)
	ret0, _ := ret[0].([]endeus.ResepCustomResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllByCategoryID indicates an expected call of GetAllByCategoryID.
func (mr *MockResepServiceMockRecorder) GetAllByCategoryID(catID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllByCategoryID", reflect.TypeOf((*MockResepService)(nil).GetAllByCategoryID), catID)
}

// GetByID mocks base method.
func (m *MockResepService) GetByID(id uint) (model.Resep, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(model.Resep)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockResepServiceMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockResepService)(nil).GetByID), id)
}

// Update mocks base method.
func (m *MockResepService) Update(id uint, param endeus.ResepParam) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, param)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockResepServiceMockRecorder) Update(id, param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockResepService)(nil).Update), id, param)
}

// MockCategoryService is a mock of CategoryService interface.
type MockCategoryService struct {
	ctrl     *gomock.Controller
	recorder *MockCategoryServiceMockRecorder
}

// MockCategoryServiceMockRecorder is the mock recorder for MockCategoryService.
type MockCategoryServiceMockRecorder struct {
	mock *MockCategoryService
}

// NewMockCategoryService creates a new mock instance.
func NewMockCategoryService(ctrl *gomock.Controller) *MockCategoryService {
	mock := &MockCategoryService{ctrl: ctrl}
	mock.recorder = &MockCategoryServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCategoryService) EXPECT() *MockCategoryServiceMockRecorder {
	return m.recorder
}

// GetAll mocks base method.
func (m *MockCategoryService) GetAll() ([]model.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]model.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockCategoryServiceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockCategoryService)(nil).GetAll))
}
