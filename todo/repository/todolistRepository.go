package repository

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"todo-list/todo"
)

type TodoListRepository struct {
	db *sql.DB
}

func NewTodoListRepository(db *sql.DB) *TodoListRepository {
	return &TodoListRepository{db: db}
}

func (t *TodoListRepository) CreateList(UserID uuid.UUID, list todo.TodoList) (int, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return 0, err
	}

	var listID int
	createListQuery := `INSERT INTO todo_lists (title, description,public,user_id) VALUES ($1, $2, $3, $4) returning id`
	row := tx.QueryRow(createListQuery, list.Title, list.Description, list.Public, UserID)
	if err := row.Scan(&listID); err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("cant creating todo_list: %w", err)
	}

	createUsersListQuery := `INSERT INTO users_lists (user_id, list_id) VALUES ($1, $2)`
	_, err = tx.Exec(createUsersListQuery, UserID, listID)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("error creating users_lists: %w", err)
	}

	return listID, tx.Commit()
}

func (t *TodoListRepository) DeleteList(UserID uuid.UUID, list todo.TodoList) error {
	deleteListQuery := `DELETE FROM todo_lists WHERE  id = $1 AND user_id = $2`
	_, err := t.db.Exec(deleteListQuery, list.ID, UserID)
	if err != nil {
		return fmt.Errorf("error deleting todo_list: %w", err)
	}

	return nil
}

func (t *TodoListRepository) GetListById(userID uuid.UUID, listID int) (*todo.TodoList, error) {
	var list todo.TodoList

	getListByIdQuery := `SELECT tl.id, tl.title, tl.description, tl.public, tl.user_id, tl.created_at, tl.updated_at FROM todo_lists tl INNER JOIN users_lists ul ON tl.id = ul.list_id WHERE ul.list_id = $1 AND ul.user_id = $2`
	row := t.db.QueryRow(getListByIdQuery, listID, userID)
	if err := row.Scan(&list.ID, &list.Title, &list.Description, &list.Public, &list.UserID, &list.CreatedAt, &list.UpdatedAt); err != nil {
		return nil, fmt.Errorf("error getting todo_list: %w", err)
	}

	return &list, nil
}

func (t *TodoListRepository) GetAllList(userID uuid.UUID) ([]todo.TodoList, error) {
	var lists []todo.TodoList
	getAllListQuery := `SELECT tl.id, tl.title, tl.description, tl.public, tl.user_id, tl.created_at, tl.updated_at  FROM todo_lists tl INNER JOIN users_lists ul ON tl.id = ul.list_id WHERE ul.user_id = $1`
	rows, err := t.db.Query(getAllListQuery, userID)
	if err != nil {
		return nil, fmt.Errorf("error getting todo_lists: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var list todo.TodoList
		if err := rows.Scan(&list.ID, &list.Title, &list.Description, &list.Public, &list.UserID, &list.CreatedAt, &list.UpdatedAt); err != nil {
			return nil, fmt.Errorf("cant scaning todo_list %w", err)
		}
		lists = append(lists, list)
	}

	return lists, nil
}

func (t *TodoListRepository) UpdateList(UserID uuid.UUID, list todo.TodoList) error {
	return nil
}
