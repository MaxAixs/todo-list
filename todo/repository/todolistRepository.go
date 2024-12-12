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

func (t *TodoListRepository) CreateList(userID uuid.UUID, list todo.TodoList) (int, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	var listID int
	createListQuery := `INSERT INTO todo_lists (title,description,public,user_id) VALUES ($1, $2, $3, $4) returning id`
	row := tx.QueryRow(createListQuery, list.Title, list.Description, list.Public, userID)
	if err := row.Scan(&listID); err != nil {
		return 0, fmt.Errorf("cant creating todo_list: %w", err)
	}

	createUsersListQuery := `INSERT INTO user_lists (user_id, list_id) VALUES ($1, $2)`
	_, err = tx.Exec(createUsersListQuery, userID, listID)
	if err != nil {
		return 0, fmt.Errorf("error creating users_lists: %w", err)
	}

	return listID, tx.Commit()
}

func (t *TodoListRepository) DeleteListById(userID uuid.UUID, listID int) error {
	deleteListQuery := `DELETE FROM todo_lists WHERE  id = $1 AND user_id = $2`
	_, err := t.db.Exec(deleteListQuery, listID, userID)
	if err != nil {
		return fmt.Errorf("error deleting todo_list: %w", err)
	}

	return nil
}

func (t *TodoListRepository) GetListById(userID uuid.UUID, listID int) (*todo.TodoList, error) {
	var list todo.TodoList

	getListByIdQuery := `SELECT tl.id, tl.title, tl.description, tl.public, tl.user_id, tl.created_at, tl.updated_at FROM todo_lists tl INNER JOIN user_lists ul ON tl.id = ul.list_id WHERE ul.list_id = $1 AND ul.user_id = $2`
	row := t.db.QueryRow(getListByIdQuery, listID, userID)
	if err := row.Scan(&list.ID, &list.Title, &list.Description, &list.Public, &list.UserID, &list.CreatedAt, &list.UpdatedAt); err != nil {
		return nil, fmt.Errorf("error getting todo_list: %w", err)
	}

	return &list, nil
}

func (t *TodoListRepository) GetAllLists(userID uuid.UUID) ([]todo.TodoList, error) {
	var lists []todo.TodoList
	getAllListQuery := `SELECT tl.id, tl.title, tl.description, tl.public, tl.user_id, tl.created_at, tl.updated_at  FROM todo_lists tl INNER JOIN user_lists ul ON tl.id = ul.list_id WHERE ul.user_id = $1`
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

func (t *TodoListRepository) UpdateList(userID uuid.UUID, listID int, input todo.UpdateListInput) error {
	setQuery, args := buildListUpdateSet(userID, listID, input)

	query := fmt.Sprintf(`
		UPDATE todo_lists tl 
		SET %s 
		FROM user_lists ul 
		WHERE tl.id = ul.list_id 
		  AND ul.user_id = $%d 
		  AND ul.list_id = $%d
	`, setQuery, len(args)-1, len(args))

	_, err := t.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("error updating todo_list: %w", err)
	}

	return nil
}
