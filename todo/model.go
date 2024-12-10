package todo

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TodoList struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	Public      bool      `json:"public"`
	UserId      int       `json:"user_id"`
	User        *User     `json:"user,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type TodoItem struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	DueDate     time.Time `json:"due_date,omitempty"`
	Priority    int       `json:"priority,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type UserList struct {
	UserID int `json:"user_id"`
	ListID int `json:"list_id"`
}

type ListItem struct {
	TodoID int `json:"todo_id"`
	ItemID int `json:"item_id"`
}
