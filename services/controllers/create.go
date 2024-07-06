package controllers

import (
	"api-postgresql/db"
	"api-postgresql/models"
)

func Create(conf *models.DBConfig, todo models.Todo) (int, error) {
	conn, err := db.OpenConnection(conf)
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	var id int
	err = conn.QueryRow(
		`INSERT INTO todos (title, description, done, in_progress, priority) VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		todo.Title, todo.Description, todo.Done, todo.InProgress, todo.Priority).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
