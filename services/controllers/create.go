package controllers

import (
	"api-postgresql/db"
	"api-postgresql/models"
)

func Create(todo models.Todo) (int, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	var id int
	err = conn.QueryRow(`INSERT INTO todos (title, description, done) VALUES ($1, $2, $3) RETURNING id`, todo.Title, todo.Description, todo.Done).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
