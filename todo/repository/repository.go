package repository

import (
	"database/sql"
	"github.com/google/uuid"
	"todo-list/pkg/notifyService"
	"todo-list/pkg/notifyService/repository"
	"todo-list/todo"
)

type Authorization interface {
	CreateUser(user todo.User) (uuid.UUID, error)
	GetUser(email string, password string) (todo.User, error)
}

type TodoList interface {
	CreateList(userID uuid.UUID, list todo.TodoList) (int, error)
	DeleteListById(userID uuid.UUID, list int) error
	GetListById(userID uuid.UUID, listID int) (*todo.TodoList, error)
	GetAllLists(userID uuid.UUID) ([]todo.TodoList, error)
	UpdateList(UserID uuid.UUID, listID int, input todo.UpdateListInput) error
}

type ItemList interface {
	CreateItem(todoID int, userID uuid.UUID, todoItems todo.TodoItem) (int, error)
	DeleteItem(userID uuid.UUID, itemID int) error
	GetItemById(userID uuid.UUID, itemID int) (*todo.TodoItem, error)
	GetAllItems(userID uuid.UUID, listID int) ([]todo.TodoItem, error)
	UpdateItem(userID uuid.UUID, itemID int, item todo.UpdateItemInput) error
}

type Notifier interface {
	GetDeadlineItems() ([]notifyService.TaskDeadlineInfo, error)
}

type Repository struct {
	Authorization
	TodoList
	ItemList
	Notifier
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
		TodoList:      NewTodoListRepository(db),
		ItemList:      NewListItemRepository(db),
		Notifier:      repository.NewDeadRepo(db),
	}
}
