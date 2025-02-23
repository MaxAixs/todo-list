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
		SELECT ti.user_id, u.email, ti.id, ti.created_at
		FROM todo_items ti
		JOIN users u ON ti.user_id = u.id
		WHERE ti.done = true AND ti.sent_analys = false;
	`

	rows, err := a.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error getting users with done items: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var item analyticsService.TaskDoneItem
		if err := rows.Scan(&item.UserID, &item.Email, &item.ItemID, &item.CreatedAt); err != nil {
			fmt.Printf("error scan users with done items: %v", err)
		}
		items = append(items, item)
	}

	return items, nil
}
