package controllers

import (
	"api-postgresql/db"
	"api-postgresql/models"
)

func Update(id int64, todo models.Todo) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE todos SET title=$1, description=$2, done=$3, in_progress=$4, priority=$5 WHERE id=$6`,
		todo.Title, todo.Description, todo.Done, todo.InProgress, todo.Priority, id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
