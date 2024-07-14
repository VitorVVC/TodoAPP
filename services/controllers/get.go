package controllers

import (
	"api-postgresql/db"
	"api-postgresql/models"
)

func Get(id int64) (models.Todo, error) {
	var todo models.Todo

	conn, err := db.OpenConnection()
	if err != nil {
		return todo, err
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM todos WHERE id = $1`, id)
	err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done, &todo.InProgress, &todo.Priority)

	return todo, err
}
