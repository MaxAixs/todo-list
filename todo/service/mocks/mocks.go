// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"
	todo "todo-list/todo"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockAuthorization is a mock of Authorization interface.
type MockAuthorization struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationMockRecorder
}

// MockAuthorizationMockRecorder is the mock recorder for MockAuthorization.
type MockAuthorizationMockRecorder struct {
	mock *MockAuthorization
}

// NewMockAuthorization creates a new mock instance.
func NewMockAuthorization(ctrl *gomock.Controller) *MockAuthorization {
	mock := &MockAuthorization{ctrl: ctrl}
	mock.recorder = &MockAuthorizationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorization) EXPECT() *MockAuthorizationMockRecorder {
	return m.recorder
}

// AuthUser mocks base method.
func (m *MockAuthorization) AuthUser(user todo.User) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthUser", user)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthUser indicates an expected call of AuthUser.
func (mr *MockAuthorizationMockRecorder) AuthUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthUser", reflect.TypeOf((*MockAuthorization)(nil).AuthUser), user)
}

// GenerateToken mocks base method.
func (m *MockAuthorization) GenerateToken(email, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", email, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockAuthorizationMockRecorder) GenerateToken(email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockAuthorization)(nil).GenerateToken), email, password)
}

// ParseToken mocks base method.
func (m *MockAuthorization) ParseToken(tokenString string) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", tokenString)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseToken indicates an expected call of ParseToken.
func (mr *MockAuthorizationMockRecorder) ParseToken(tokenString interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockAuthorization)(nil).ParseToken), tokenString)
}

// MockTodoList is a mock of TodoList interface.
type MockTodoList struct {
	ctrl     *gomock.Controller
	recorder *MockTodoListMockRecorder
}

// MockTodoListMockRecorder is the mock recorder for MockTodoList.
type MockTodoListMockRecorder struct {
	mock *MockTodoList
}

// NewMockTodoList creates a new mock instance.
func NewMockTodoList(ctrl *gomock.Controller) *MockTodoList {
	mock := &MockTodoList{ctrl: ctrl}
	mock.recorder = &MockTodoListMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTodoList) EXPECT() *MockTodoListMockRecorder {
	return m.recorder
}

// CreateList mocks base method.
func (m *MockTodoList) CreateList(userID uuid.UUID, list todo.TodoList) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateList", userID, list)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateList indicates an expected call of CreateList.
func (mr *MockTodoListMockRecorder) CreateList(userID, list interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateList", reflect.TypeOf((*MockTodoList)(nil).CreateList), userID, list)
}

// DeleteListById mocks base method.
func (m *MockTodoList) DeleteListById(userID uuid.UUID, listID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteListById", userID, listID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteListById indicates an expected call of DeleteListById.
func (mr *MockTodoListMockRecorder) DeleteListById(userID, listID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteListById", reflect.TypeOf((*MockTodoList)(nil).DeleteListById), userID, listID)
}

// GetAllLists mocks base method.
func (m *MockTodoList) GetAllLists(userID uuid.UUID) ([]todo.TodoList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllLists", userID)
	ret0, _ := ret[0].([]todo.TodoList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllLists indicates an expected call of GetAllLists.
func (mr *MockTodoListMockRecorder) GetAllLists(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllLists", reflect.TypeOf((*MockTodoList)(nil).GetAllLists), userID)
}

// GetListById mocks base method.
func (m *MockTodoList) GetListById(userID uuid.UUID, listID int) (*todo.TodoList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListById", userID, listID)
	ret0, _ := ret[0].(*todo.TodoList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListById indicates an expected call of GetListById.
func (mr *MockTodoListMockRecorder) GetListById(userID, listID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListById", reflect.TypeOf((*MockTodoList)(nil).GetListById), userID, listID)
}

// UpdateList mocks base method.
func (m *MockTodoList) UpdateList(userID uuid.UUID, listID int, list todo.UpdateListInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateList", userID, listID, list)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateList indicates an expected call of UpdateList.
func (mr *MockTodoListMockRecorder) UpdateList(userID, listID, list interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateList", reflect.TypeOf((*MockTodoList)(nil).UpdateList), userID, listID, list)
}

// MockTodoItem is a mock of TodoItem interface.
type MockTodoItem struct {
	ctrl     *gomock.Controller
	recorder *MockTodoItemMockRecorder
}

// MockTodoItemMockRecorder is the mock recorder for MockTodoItem.
type MockTodoItemMockRecorder struct {
	mock *MockTodoItem
}

// NewMockTodoItem creates a new mock instance.
func NewMockTodoItem(ctrl *gomock.Controller) *MockTodoItem {
	mock := &MockTodoItem{ctrl: ctrl}
	mock.recorder = &MockTodoItemMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTodoItem) EXPECT() *MockTodoItemMockRecorder {
	return m.recorder
}

// CreateItem mocks base method.
func (m *MockTodoItem) CreateItem(userID uuid.UUID, listID int, item todo.TodoItem) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateItem", userID, listID, item)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateItem indicates an expected call of CreateItem.
func (mr *MockTodoItemMockRecorder) CreateItem(userID, listID, item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateItem", reflect.TypeOf((*MockTodoItem)(nil).CreateItem), userID, listID, item)
}

// DeleteItemById mocks base method.
func (m *MockTodoItem) DeleteItemById(userID uuid.UUID, itemID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteItemById", userID, itemID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteItemById indicates an expected call of DeleteItemById.
func (mr *MockTodoItemMockRecorder) DeleteItemById(userID, itemID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteItemById", reflect.TypeOf((*MockTodoItem)(nil).DeleteItemById), userID, itemID)
}

// GetAllItems mocks base method.
func (m *MockTodoItem) GetAllItems(userID uuid.UUID, listID int) ([]todo.TodoItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllItems", userID, listID)
	ret0, _ := ret[0].([]todo.TodoItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllItems indicates an expected call of GetAllItems.
func (mr *MockTodoItemMockRecorder) GetAllItems(userID, listID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllItems", reflect.TypeOf((*MockTodoItem)(nil).GetAllItems), userID, listID)
}

// GetItemById mocks base method.
func (m *MockTodoItem) GetItemById(userID uuid.UUID, itemID int) (*todo.TodoItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItemById", userID, itemID)
	ret0, _ := ret[0].(*todo.TodoItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItemById indicates an expected call of GetItemById.
func (mr *MockTodoItemMockRecorder) GetItemById(userID, itemID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItemById", reflect.TypeOf((*MockTodoItem)(nil).GetItemById), userID, itemID)
}

// UpdateItem mocks base method.
func (m *MockTodoItem) UpdateItem(userID uuid.UUID, itemID int, item todo.UpdateItemInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateItem", userID, itemID, item)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateItem indicates an expected call of UpdateItem.
func (mr *MockTodoItemMockRecorder) UpdateItem(userID, itemID, item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateItem", reflect.TypeOf((*MockTodoItem)(nil).UpdateItem), userID, itemID, item)
}
