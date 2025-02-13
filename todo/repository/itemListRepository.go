package repository

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"todo-list/todo"
)

type ListItemRepository struct {
	db *sql.DB
}

func NewListItemRepository(db *sql.DB) *ListItemRepository {
	return &ListItemRepository{db: db}
}

func (l *ListItemRepository) CreateItem(todoID int, userID uuid.UUID, todoItems todo.TodoItem) (int, error) {
	tx, err := l.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	var itemID int
	createItemQuery := `INSERT INTO todo_items (description, done, due_date, user_id) 
                         VALUES ($1, $2, $3, $4) 
                         returning id`
	row := tx.QueryRow(createItemQuery, todoItems.Description, todoItems.Done, todoItems.DueDate, userID)
	err = row.Scan(&itemID)
	if err != nil {
		return 0, fmt.Errorf("can't create todo_items: %w", err)
	}

	createListItemQuery := `INSERT INTO list_items (todo_id, item_id) VALUES ($1, $2)`
	_, err = tx.Exec(createListItemQuery, todoID, itemID)
	if err != nil {
		return 0, fmt.Errorf("can't create list_items: %w", err)
	}

	return itemID, tx.Commit()
}

func (l *ListItemRepository) DeleteItem(UserID uuid.UUID, itemID int) error {
	deleteItemQuery := `DELETE FROM todo_items ti USING list_items li, user_lists ul WHERE ti.id = li.item_id AND li.todo_id = ul.list_id AND ul.user_id = $1 AND ti.id = $2`
	_, err := l.db.Exec(deleteItemQuery, UserID, itemID)
	if err != nil {
		return fmt.Errorf("cant delete item %w", err)
	}

	return nil
}

func (l *ListItemRepository) GetItemById(userID uuid.UUID, itemID int) (*todo.TodoItem, error) {
	var item todo.TodoItem
	getItemById := `SELECT ti.id, ti.description, ti.done, ti.priority FROM todo_items ti INNER JOIN list_items li on li.item_id = ti.id INNER JOIN user_lists ul on ul.list_id = li.todo_id WHERE  ti.id = $1 AND ul.user_id = $2`
	row := l.db.QueryRow(getItemById, itemID, userID)
	if err := row.Scan(&item.ID, &item.Description, &item.Done, &item.Priority); err != nil {
		return nil, fmt.Errorf("cant get item by id %w", err)
	}

	return &item, nil

}

func (l *ListItemRepository) GetAllItems(userID uuid.UUID, listID int) ([]todo.TodoItem, error) {
	var items []todo.TodoItem

	query := `
		SELECT ti.id, ti.description, ti.done, ti.due_date, ti.priority, ti.created_at, ti.updated_at
		FROM todo_items ti
		INNER JOIN list_items li ON li.item_id = ti.id
		INNER JOIN user_lists ul ON ul.list_id = li.todo_id
		WHERE li.todo_id = $1 AND ul.user_id = $2;
	`
	rows, err := l.db.Query(query, listID, userID)
	if err != nil {
		return nil, fmt.Errorf("cant get items: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var item todo.TodoItem
		if err := rows.Scan(&item.ID, &item.Description, &item.Done, &item.DueDate, &item.Priority, &item.CreatedAt, &item.UpdatedAt); err != nil {
			return nil, fmt.Errorf("error scanning item: %w", err)
		}
		items = append(items, item)
	}

	return items, nil
}

func (l *ListItemRepository) UpdateItem(userID uuid.UUID, itemID int, input todo.UpdateItemInput) error {
	setQuery, args := buildItemUpdateSet(userID, itemID, input)

	query := fmt.Sprintf(`
		UPDATE todo_items ti 
		SET %s 
		FROM list_items li, user_lists ul 
		WHERE ti.id = li.item_id 
		  AND li.todo_id = ul.list_id 
		  AND ul.user_id = $%d 
		  AND ti.id = $%d
	`, setQuery, len(args)-1, len(args))

	_, err := l.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("error updating todo_item: %w", err)
	}

	return nil
}
