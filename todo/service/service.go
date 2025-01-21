package service

import (
	"github.com/google/uuid"
	"todo-list/todo"
	"todo-list/todo/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mockAuth.go

type Authorization interface {
	AuthUser(user todo.User) (uuid.UUID, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(tokenString string) (uuid.UUID, error)
}

type TodoList interface {
	CreateList(userID uuid.UUID, list todo.TodoList) (int, error)
	GetAllLists(userID uuid.UUID) ([]todo.TodoList, error)
	GetListById(userID uuid.UUID, listID int) (*todo.TodoList, error)
	DeleteListById(userID uuid.UUID, listID int) error
	UpdateList(userID uuid.UUID, listID int, list todo.UpdateListInput) error
}

type TodoItem interface {
	CreateItem(userID uuid.UUID, listID int, item todo.TodoItem) (int, error)
	DeleteItemById(userID uuid.UUID, itemID int) error
	GetAllItems(userID uuid.UUID, listID int) ([]todo.TodoItem, error)
	GetItemById(userID uuid.UUID, itemID int) (*todo.TodoItem, error)
	UpdateItem(userID uuid.UUID, itemID int, item todo.UpdateItemInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		TodoList:      NewListService(repo.TodoList),
		TodoItem:      NewItemService(repo.ItemList, repo.TodoList),
	}
}
