package repository

import (
	"database/sql"
	"fmt"
	"todo-list/pkg/notifyService"
)

type DeadLineRepo struct {
	db *sql.DB
}

func NewDeadRepo(db *sql.DB) *DeadLineRepo {
	return &DeadLineRepo{db: db}
}

func (d *DeadLineRepo) GetDeadlineItems() ([]notifyService.TaskDeadlineInfo, error) {
	var users []notifyService.TaskDeadlineInfo

	query := `
		UPDATE todo_items ti
		SET sent_deadline = true
		FROM users u
		WHERE ti.user_id = u.id
		AND ti.due_date BETWEEN NOW() AND NOW() + INTERVAL '5 minutes'
		AND ti.sent_deadline = false
		RETURNING u.email, ti.user_id, ti.id AS item_id, ti.description;
	`

	rows, err := d.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("cant get and update deadline Items: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user notifyService.TaskDeadlineInfo
		if err := rows.Scan(&user.Email, &user.UserID, &user.ItemID, &user.Description); err != nil {
			return nil, fmt.Errorf("cant scan user: %v", err)
		}
		users = append(users, user)
	}

	return users, nil
}
