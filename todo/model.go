package todo

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name" validate:"required,min=3,max=50"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required,min=6"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TodoList struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	Public      bool      `json:"public"`
	UserID      uuid.UUID `json:"user_id"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type TodoItem struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	DueDate     time.Time `json:"due_date,omitempty"`
	Priority    int       `json:"priority,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type UserList struct {
	UserID uuid.UUID `json:"user_id"`
	ListID int       `json:"list_id"`
}

type ListItem struct {
	TodoID int `json:"todo_id"`
	ItemID int `json:"item_id"`
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Public      *bool   `json:"public"`
	UpdatedAt   *string `json:"updated_at"`
}

type UpdateItemInput struct {
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
	DueDate     *string `json:"due_date"`
	Priority    *int    `json:"priority"`
	UpdatedAt   *string `json:"updated_at"`
}
