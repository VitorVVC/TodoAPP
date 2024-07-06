package controllers

import (
	"api-postgresql/db"
	"api-postgresql/models"
)

func Delete(conf *models.DBConfig, id int64) (int64, error) {
	conn, err := db.OpenConnection(conf)
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM todos WHERE id=$1`, id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
