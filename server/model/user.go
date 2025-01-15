package model

type User struct {
	ID           		int    `db:"id"`
	Name         		string `db:"name"`
	Email        		string `db:"email"`
	PasswordHash 		string `db:"password_hash"`
	DefaultTodoListID 	int    `db:"default_todo_list_id"`
	CreatedAt    		string `db:"created_at"`
	UpdatedAt    		string `db:"updated_at"`
}
