package repository

import (
	"database/sql"
	"fmt"
	"todo-list/pkg/analyticsService"
)

type AnalyticRepo struct {
	db *sql.DB
}

func NewAnalyticRepo(db *sql.DB) *AnalyticRepo {
	return &AnalyticRepo{db: db}
}

func (a *AnalyticRepo) GetUsersWithDoneItem() ([]analyticsService.TaskDoneItem, error) {
	var items []analyticsService.TaskDoneItem

	query := `
		UPDATE todo_items ti
		SET sent_analys = true
		FROM users u
		WHERE ti.user_id = u.id
		AND ti.done = true
		AND ti.sent_analys = false
		RETURNING u.email, ti.user_id, ti.id AS item_id, ti.created_at;
	`

	rows, err := a.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error updating and getting users with done items: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var item analyticsService.TaskDoneItem
		if err := rows.Scan(&item.Email, &item.UserID, &item.ItemID, &item.CreatedAt); err != nil {
			fmt.Printf("error scan users with done items: %v", err)
		}
		items = append(items, item)
	}

	return items, nil
}
