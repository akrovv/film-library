// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	domain "github.com/akrovv/filmlibrary/internal/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockActorService is a mock of ActorService interface.
type MockActorService struct {
	ctrl     *gomock.Controller
	recorder *MockActorServiceMockRecorder
}

// MockActorServiceMockRecorder is the mock recorder for MockActorService.
type MockActorServiceMockRecorder struct {
	mock *MockActorService
}

// NewMockActorService creates a new mock instance.
func NewMockActorService(ctrl *gomock.Controller) *MockActorService {
	mock := &MockActorService{ctrl: ctrl}
	mock.recorder = &MockActorServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockActorService) EXPECT() *MockActorServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockActorService) Create(dto *domain.CreateActor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", dto)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockActorServiceMockRecorder) Create(dto interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockActorService)(nil).Create), dto)
}

// Delete mocks base method.
func (m *MockActorService) Delete(dto *domain.DeleteActor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", dto)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockActorServiceMockRecorder) Delete(dto interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockActorService)(nil).Delete), dto)
}

// GetList mocks base method.
func (m *MockActorService) GetList() ([]domain.ActorWithMovie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetList")
	ret0, _ := ret[0].([]domain.ActorWithMovie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetList indicates an expected call of GetList.
func (mr *MockActorServiceMockRecorder) GetList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetList", reflect.TypeOf((*MockActorService)(nil).GetList))
}

// Update mocks base method.
func (m *MockActorService) Update(dto *domain.Actor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", dto)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockActorServiceMockRecorder) Update(dto interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockActorService)(nil).Update), dto)
}