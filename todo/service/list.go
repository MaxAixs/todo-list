package service

import (
	"fmt"
	"github.com/google/uuid"
	"todo-list/todo"
	"todo-list/todo/repository"
)

type ListService struct {
	repo repository.TodoList
}

func NewListService(repo repository.TodoList) *ListService {
	return &ListService{repo: repo}
}

func (l *ListService) CreateList(UserID uuid.UUID, list todo.TodoList) (int, error) {
	return l.repo.CreateList(UserID, list)
}

func (l *ListService) GetAllLists(UserID uuid.UUID) ([]todo.TodoList, error) {
	return l.repo.GetAllLists(UserID)
}

func (l *ListService) GetListById(UserID uuid.UUID, listID int) (*todo.TodoList, error) {
	return l.repo.GetListById(UserID, listID)
}

func (l *ListService) DeleteListById(UserID uuid.UUID, listID int) error {
	return l.repo.DeleteListById(UserID, listID)
}

func (l *ListService) UpdateList(UserID uuid.UUID, listID int, list todo.UpdateListInput) error {
	if err := ValidateInputUpdateList(list); err != nil {
		return err
	}
	return l.repo.UpdateList(UserID, listID, list)
}

func ValidateInputUpdateList(list todo.UpdateListInput) error {
	if list.Title == nil &&
		list.Description == nil &&
		list.Public == nil &&
		list.UpdatedAt == nil {
		return fmt.Errorf("update input cannot be empty: at least one field must be provided")
	}

	return nil
}
