// Code generated by MockGen. DO NOT EDIT.
// Source: CleverIT-challenge/internal/core/domain/beers (interfaces: Repository)

// Package beersmocks is a generated GoMock package.
package beersmocks

import (
	beers "CleverIT-challenge/internal/core/domain/beers"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// FindAllBeers mocks base method.
func (m *MockRepository) FindAllBeers(arg0 context.Context) ([]beers.Beer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllBeers", arg0)
	ret0, _ := ret[0].([]beers.Beer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllBeers indicates an expected call of FindAllBeers.
func (mr *MockRepositoryMockRecorder) FindAllBeers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllBeers", reflect.TypeOf((*MockRepository)(nil).FindAllBeers), arg0)
}

// FindBeerByID mocks base method.
func (m *MockRepository) FindBeerByID(arg0 context.Context, arg1 int) (beers.Beer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindBeerByID", arg0, arg1)
	ret0, _ := ret[0].(beers.Beer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindBeerByID indicates an expected call of FindBeerByID.
func (mr *MockRepositoryMockRecorder) FindBeerByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindBeerByID", reflect.TypeOf((*MockRepository)(nil).FindBeerByID), arg0, arg1)
}

// SaveBeer mocks base method.
func (m *MockRepository) SaveBeer(arg0 context.Context, arg1 beers.Beer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveBeer", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveBeer indicates an expected call of SaveBeer.
func (mr *MockRepositoryMockRecorder) SaveBeer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveBeer", reflect.TypeOf((*MockRepository)(nil).SaveBeer), arg0, arg1)
}
