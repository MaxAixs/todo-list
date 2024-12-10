package repository

import (
	"database/sql"
	"fmt"
	"todo-list/todo"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (a *AuthRepository) CreateUser(user todo.User) (int, error) {
	var ID int
	query := `INSERT INTO users (name, email, password_hash) VALUES ($1, $2, $3) returning id`

	row := a.db.QueryRow(query, user.Name, user.Email, user.Password)
	if err := row.Scan(&ID); err != nil {
		return 0, fmt.Errorf("cant create user: %w", err)
	}

	return ID, nil
}

func (a *AuthRepository) GetUser(email string, password string) (todo.User, error) {
	var user todo.User
	query := `SELECT id,name,email,created_at,updated_at FROM users WHERE email = $1 AND password_hash = $2`

	row := a.db.QueryRow(query, email, password)
	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return todo.User{}, fmt.Errorf("cant get user: %w", err)
	}

	return user, nil
}
