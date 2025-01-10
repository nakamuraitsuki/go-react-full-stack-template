package migration

import (
	"github.com/jmoiron/sqlx"
)

func Migrate(db *sqlx.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INT PRIMARY KEY AUTO_INCREMENT,
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL UNIQUE,
			password_hash VARCHAR(255) NOT NULL,
			default_todo_list_id INT NOT NULL,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
			FOREIGN KEY (default_todo_list_id) REFERENCES todo_lists(id)
		)
	`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS todo_lists (
			id INT PRIMARY KEY AUTO_INCREMENT,
			user_id INT NOT NULL,
			name VARCHAR(255) NOT NULL,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(id)
		)
	`)

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS todos (
			id INT PRIMARY KEY AUTO_INCREMENT,
			todo_list_id INT NOT NULL,
			title VARCHAR(255) NOT NULL,
			completed BOOLEAN NOT NULL DEFAULT FALSE,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			FOREIGN KEY (todo_list_id) REFERENCES todo_lists(id)
		)
	`)
	if err != nil {
		return err
	}

	return nil
}
