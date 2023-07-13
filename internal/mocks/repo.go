// Code generated by MockGen. DO NOT EDIT.
// Source: repo.go

// Package mocks is a generated GoMock package.
package mocks

import (
	endeus "endeus/wawan/pkg/endeus"
	model "endeus/wawan/pkg/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCaraMasakRepository is a mock of CaraMasakRepository interface.
type MockCaraMasakRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCaraMasakRepositoryMockRecorder
}

// MockCaraMasakRepositoryMockRecorder is the mock recorder for MockCaraMasakRepository.
type MockCaraMasakRepositoryMockRecorder struct {
	mock *MockCaraMasakRepository
}

// NewMockCaraMasakRepository creates a new mock instance.
func NewMockCaraMasakRepository(ctrl *gomock.Controller) *MockCaraMasakRepository {
	mock := &MockCaraMasakRepository{ctrl: ctrl}
	mock.recorder = &MockCaraMasakRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCaraMasakRepository) EXPECT() *MockCaraMasakRepositoryMockRecorder {
	return m.recorder
}

// GetByResepID mocks base method.
func (m *MockCaraMasakRepository) GetByResepID(resepID uint) (model.CaraMasak, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByResepID", resepID)
	ret0, _ := ret[0].(model.CaraMasak)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByResepID indicates an expected call of GetByResepID.
func (mr *MockCaraMasakRepositoryMockRecorder) GetByResepID(resepID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByResepID", reflect.TypeOf((*MockCaraMasakRepository)(nil).GetByResepID), resepID)
}

// MockBahanRepository is a mock of BahanRepository interface.
type MockBahanRepository struct {
	ctrl     *gomock.Controller
	recorder *MockBahanRepositoryMockRecorder
}

// MockBahanRepositoryMockRecorder is the mock recorder for MockBahanRepository.
type MockBahanRepositoryMockRecorder struct {
	mock *MockBahanRepository
}

// NewMockBahanRepository creates a new mock instance.
func NewMockBahanRepository(ctrl *gomock.Controller) *MockBahanRepository {
	mock := &MockBahanRepository{ctrl: ctrl}
	mock.recorder = &MockBahanRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBahanRepository) EXPECT() *MockBahanRepositoryMockRecorder {
	return m.recorder
}

// GetByResepID mocks base method.
func (m *MockBahanRepository) GetByResepID(resepID uint) (model.Bahan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByResepID", resepID)
	ret0, _ := ret[0].(model.Bahan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByResepID indicates an expected call of GetByResepID.
func (mr *MockBahanRepositoryMockRecorder) GetByResepID(resepID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByResepID", reflect.TypeOf((*MockBahanRepository)(nil).GetByResepID), resepID)
}

// MockCategoryRepository is a mock of CategoryRepository interface.
type MockCategoryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCategoryRepositoryMockRecorder
}

// MockCategoryRepositoryMockRecorder is the mock recorder for MockCategoryRepository.
type MockCategoryRepositoryMockRecorder struct {
	mock *MockCategoryRepository
}

// NewMockCategoryRepository creates a new mock instance.
func NewMockCategoryRepository(ctrl *gomock.Controller) *MockCategoryRepository {
	mock := &MockCategoryRepository{ctrl: ctrl}
	mock.recorder = &MockCategoryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCategoryRepository) EXPECT() *MockCategoryRepositoryMockRecorder {
	return m.recorder
}

// GetAll mocks base method.
func (m *MockCategoryRepository) GetAll() ([]model.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]model.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockCategoryRepositoryMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockCategoryRepository)(nil).GetAll))
}

// MockResepRepository is a mock of ResepRepository interface.
type MockResepRepository struct {
	ctrl     *gomock.Controller
	recorder *MockResepRepositoryMockRecorder
}

// MockResepRepositoryMockRecorder is the mock recorder for MockResepRepository.
type MockResepRepositoryMockRecorder struct {
	mock *MockResepRepository
}

// NewMockResepRepository creates a new mock instance.
func NewMockResepRepository(ctrl *gomock.Controller) *MockResepRepository {
	mock := &MockResepRepository{ctrl: ctrl}
	mock.recorder = &MockResepRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResepRepository) EXPECT() *MockResepRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockResepRepository) Create(param endeus.ResepParam) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", param)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockResepRepositoryMockRecorder) Create(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockResepRepository)(nil).Create), param)
}

// GetAll mocks base method.
func (m *MockResepRepository) GetAll(str string) ([]model.Resep, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", str)
	ret0, _ := ret[0].([]model.Resep)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockResepRepositoryMockRecorder) GetAll(str interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockResepRepository)(nil).GetAll), str)
}

// GetAllByCategoryID mocks base method.
func (m *MockResepRepository) GetAllByCategoryID(categoryID uint) ([]model.Resep, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllByCategoryID", categoryID)
	ret0, _ := ret[0].([]model.Resep)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllByCategoryID indicates an expected call of GetAllByCategoryID.
func (mr *MockResepRepositoryMockRecorder) GetAllByCategoryID(categoryID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllByCategoryID", reflect.TypeOf((*MockResepRepository)(nil).GetAllByCategoryID), categoryID)
}

// GetByID mocks base method.
func (m *MockResepRepository) GetByID(resepID uint) (model.Resep, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", resepID)
	ret0, _ := ret[0].(model.Resep)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockResepRepositoryMockRecorder) GetByID(resepID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockResepRepository)(nil).GetByID), resepID)
}

// Update mocks base method.
func (m *MockResepRepository) Update(id uint, param endeus.ResepParam) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, param)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockResepRepositoryMockRecorder) Update(id, param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockResepRepository)(nil).Update), id, param)
}