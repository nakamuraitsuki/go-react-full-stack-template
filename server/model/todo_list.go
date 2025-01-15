package model

import "time"

type TodoList struct {
	ID        int       `db:"id"`
	UserID    int       `db:"user_id"`
	Name     string		`db:"name"`
	IsDefault bool      `db:"is_default"` 
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}