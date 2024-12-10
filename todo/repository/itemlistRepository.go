package repository

import (
	"database/sql"
	"fmt"
	"todo-list/todo"
)

type ListItemRepository struct {
	db *sql.DB
}

func NewListItemRepository(db *sql.DB) *ListItemRepository {
	return &ListItemRepository{db: db}
}

func (l *ListItemRepository) CreateItem(todoID int, todoItems todo.TodoItem) (int, error) {
	tx, err := l.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemID int
	createItemQuery := `INSERT INTO todo_items (description,done,priority) VALUES ($1, $2, $3) returning id`
	row := tx.QueryRow(createItemQuery, todoItems.Description, todoItems.Done, todoItems.Priority)
	if err := row.Scan(&itemID); err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("cant create todo_items %w", err)
	}

	createListItemQuery := `INSERT INTO list_items (list_id, item_id) (VALUES $1, $2)`
	_, err = tx.Exec(createListItemQuery, todoID, itemID)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("cant create list_items %w", err)
	}

	return itemID, nil
}

func (l *ListItemRepository) DeleteItem(UserID int, itemID int) error {
	deleteItemQuery := `DELETE FROM todo_items ti USING list_items li, user_lists ul WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $1 AND ti.id = $2`
	_, err := l.db.Exec(deleteItemQuery, UserID, itemID)
	if err != nil {
		return fmt.Errorf("cant delete item %w", err)
	}

	return nil
}
func (l *ListItemRepository) GetItemById(userID, itemID int) (*todo.TodoItem, error) {
	var item todo.TodoItem
	getItemById := `SELECT ti.id, ti.description, ti.done, ti.priority FROM todo_items ti INNER JOIN list_items li on li.item_id = ti.id INNER JOIN user_lists ul on ul.list_id = li.list_id WHERE  ti.id = $1 AND ul.user_id = $2`
	row := l.db.QueryRow(getItemById, itemID, userID)
	if err := row.Scan(&item.ID, &item.Description, &item.Done, &item.Priority); err != nil {
		return nil, fmt.Errorf("cant get item by id %w", err)
	}

	return &item, nil

}

func (l *ListItemRepository) GetAllItems(userID int, listID int) ([]todo.TodoItem, error) {
	var items []todo.TodoItem
	getAllItemsQuery := `SELECT ti.id, ti.description, ti.done, ti.priority FROM todo_items ti INNER JOIN list_items li on li.item_id = ti.id INNER JOIN user_lists ul on ul.list_id = li.list_id WHERE li.list_id = $1 AND ul.user_id = $2`
	rows, err := l.db.Query(getAllItemsQuery, listID, userID)
	if err != nil {
		return nil, fmt.Errorf("cant get Items %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var item todo.TodoItem
		if err := rows.Scan(&item.ID, &item.Description, &item.Done, &item.Priority); err != nil {
			return nil, fmt.Errorf("error scanning item: %w", err)
		}
		items = append(items, item)
	}

	return items, nil
}

func (l *ListItemRepository) UpdateItems() error {
	return nil
}
