package controllers

import (
	"api-postgresql/db"
)

func Insert(todo Todo) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	sql := `INSERT INTO todos (title, description, done) VALUES ($1, $2, $3) RETURNING id;`

	var id int64
	err = conn.QueryRow(sql, todo.Title, todo.Description, todo.Done).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
