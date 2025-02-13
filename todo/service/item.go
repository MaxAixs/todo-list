package service

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"todo-list/todo"
	"todo-list/todo/repository"
)

type ItemService struct {
	repo     repository.ItemList
	repoList repository.TodoList
}

func NewItemService(repo repository.ItemList, repoList repository.TodoList) *ItemService {
	return &ItemService{repo: repo, repoList: repoList}
}

func (i *ItemService) CreateItem(userID uuid.UUID, listID int, item todo.TodoItem) (int, error) {
	_, err := i.repoList.GetListById(userID, listID)
	if err != nil {
		return 0, errors.New("list doesn't exist or does not belong to user")
	}

	return i.repo.CreateItem(listID, userID, item)
}

func (i *ItemService) DeleteItemById(userID uuid.UUID, itemID int) error {
	return i.repo.DeleteItem(userID, itemID)
}

func (i *ItemService) GetAllItems(userID uuid.UUID, listID int) ([]todo.TodoItem, error) {
	return i.repo.GetAllItems(userID, listID)
}

func (i *ItemService) GetItemById(userID uuid.UUID, itemID int) (*todo.TodoItem, error) {
	return i.repo.GetItemById(userID, itemID)
}

func (i *ItemService) UpdateItem(userID uuid.UUID, itemID int, item todo.UpdateItemInput) error {
	err := ValidateItemInput(item)
	if err != nil {
		return err
	}

	return i.repo.UpdateItem(userID, itemID, item)
}

func ValidateItemInput(item todo.UpdateItemInput) error {
	if item.Description == nil &&
		item.Done == nil &&
		item.DueDate == nil &&
		item.Priority == nil &&
		item.UpdatedAt == nil {
		return fmt.Errorf("update input cannot be empty: at least one field must be provided")
	}

	return nil
}
