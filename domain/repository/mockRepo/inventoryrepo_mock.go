// Code generated by MockGen. DO NOT EDIT.
// Source: domain/repository/inventoryRepository.go

// Package mockRepo is a generated GoMock package.
package mockRepo

import (
	entity "70_Off/domain/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockInventory is a mock of Inventory interface.
type MockInventory struct {
	ctrl     *gomock.Controller
	recorder *MockInventoryMockRecorder
}

// MockInventoryMockRecorder is the mock recorder for MockInventory.
type MockInventoryMockRecorder struct {
	mock *MockInventory
}

// NewMockInventory creates a new mock instance.
func NewMockInventory(ctrl *gomock.Controller) *MockInventory {
	mock := &MockInventory{ctrl: ctrl}
	mock.recorder = &MockInventoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInventory) EXPECT() *MockInventoryMockRecorder {
	return m.recorder
}

// CreateInventory mocks base method.
func (m *MockInventory) CreateInventory(inventory *entity.Inventory) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInventory", inventory)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateInventory indicates an expected call of CreateInventory.
func (mr *MockInventoryMockRecorder) CreateInventory(inventory interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInventory", reflect.TypeOf((*MockInventory)(nil).CreateInventory), inventory)
}

// DeleteInventory mocks base method.
func (m *MockInventory) DeleteInventory(inventory *entity.Inventory) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteInventory", inventory)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteInventory indicates an expected call of DeleteInventory.
func (mr *MockInventoryMockRecorder) DeleteInventory(inventory interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteInventory", reflect.TypeOf((*MockInventory)(nil).DeleteInventory), inventory)
}

// GetInventoryByID mocks base method.
func (m *MockInventory) GetInventoryByID(id uint) (*entity.Inventory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInventoryByID", id)
	ret0, _ := ret[0].(*entity.Inventory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInventoryByID indicates an expected call of GetInventoryByID.
func (mr *MockInventoryMockRecorder) GetInventoryByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInventoryByID", reflect.TypeOf((*MockInventory)(nil).GetInventoryByID), id)
}

// GetInventoryByProductItemID mocks base method.
func (m *MockInventory) GetInventoryByProductItemID(id uint) (*entity.Inventory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInventoryByProductItemID", id)
	ret0, _ := ret[0].(*entity.Inventory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInventoryByProductItemID indicates an expected call of GetInventoryByProductItemID.
func (mr *MockInventoryMockRecorder) GetInventoryByProductItemID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInventoryByProductItemID", reflect.TypeOf((*MockInventory)(nil).GetInventoryByProductItemID), id)
}

// ReduceQuantity mocks base method.
func (m *MockInventory) ReduceQuantity(inventory *entity.Inventory, quantity uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReduceQuantity", inventory, quantity)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReduceQuantity indicates an expected call of ReduceQuantity.
func (mr *MockInventoryMockRecorder) ReduceQuantity(inventory, quantity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReduceQuantity", reflect.TypeOf((*MockInventory)(nil).ReduceQuantity), inventory, quantity)
}

// UpdateInventory mocks base method.
func (m *MockInventory) UpdateInventory(inventory *entity.Inventory) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInventory", inventory)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateInventory indicates an expected call of UpdateInventory.
func (mr *MockInventoryMockRecorder) UpdateInventory(inventory interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInventory", reflect.TypeOf((*MockInventory)(nil).UpdateInventory), inventory)
}