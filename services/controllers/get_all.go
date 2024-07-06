package controllers

import (
	"api-postgresql/db"
	"api-postgresql/models"
)

func GetAll(conf *models.DBConfig) ([]models.Todo, error) {
	conn, err := db.OpenConnection(conf)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT id, title, description, done FROM todos`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}
