package controllers

import (
	"api-postgresql/db"
	"api-postgresql/models"
)

func GetAll() ([]models.Todo, error) {
	var todos []models.Todo

	conn, err := db.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM todos`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var todo models.Todo
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done, &todo.InProgress, &todo.Priority)
		if err != nil {
			continue
		}

		todos = append(todos, todo)
	}

	return todos, nil
}
