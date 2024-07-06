package controllers

import (
	"api-postgresql/db"
	"api-postgresql/models"
)

func Get(conf *models.DBConfig, id int64) (models.Todo, error) {
	conn, err := db.OpenConnection(conf)
	if err != nil {
		return models.Todo{}, err
	}
	defer conn.Close()

	var todo models.Todo
	row := conn.QueryRow(`SELECT id, title, description, done, in_progress, priority FROM todos WHERE id = $1`, id)
	err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done, &todo.InProgress, &todo.Priority)
	if err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}
