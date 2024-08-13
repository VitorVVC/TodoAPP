package controllers

import (
	"api-postgresql/db"
	"api-postgresql/models"
	"api-postgresql/utils"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type TodoController struct {
	validator *validator.Validate
}

func NewTodoController() *TodoController {
	return &TodoController{
		validator: validator.New(),
	}
}

func (t *TodoController) Create(ctx echo.Context) error {
	var data models.Todo
	if err := ctx.Bind(&data); err != nil {
		return utils.HTTPFail(ctx, http.StatusBadRequest, err, "failed to parse body")
	}

	if err := t.validator.Struct(data); err != nil {
		return utils.HTTPFail(ctx, http.StatusBadRequest, err, "failed to validate body")
	}

	data.UUID = uuid.New()

	conn, err := db.OpenConnection()
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusInternalServerError, err, "failed to connect to database")
	}
	defer conn.Close()

	_, err = conn.Exec(`INSERT INTO todos (uuid, title, description, done, in_progress, priority) VALUES ($1, $2, $3, $4, $5, $6)`, data.UUID, data.Title, data.Description, data.Done, data.InProgress, data.Priority)
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusInternalServerError, err, "Failed to create todo: ")
	}

	return utils.HTTPCreated(ctx, data)
}

func (t *TodoController) Update(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusBadRequest, err, "invalid id")
	}

	var data models.Todo
	if err := ctx.Bind(&data); err != nil {
		return utils.HTTPFail(ctx, http.StatusBadRequest, err, "failed to parse body")
	}

	if err := t.validator.Struct(data); err != nil {
		return utils.HTTPFail(ctx, http.StatusBadRequest, err, "failed to validate body")
	}

	conn, err := db.OpenConnection()
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusInternalServerError, err, "Failed to connect to database")
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE todos SET title=$1, description=$2, done=$3, in_progress=$4, priority=$5 WHERE id=$6`,
		data.Title, data.Description, data.Done, data.InProgress, data.Priority, id)
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusInternalServerError, err, "failed to update todo")
	}

	_, err = res.RowsAffected()
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusInternalServerError, err, "failed to retrieve update result")
	}

	return utils.HTTPCreated(ctx, data)
}

func (t *TodoController) Delete(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusBadRequest, err, "invalid id")
	}

	conn, err := db.OpenConnection()
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusInternalServerError, err, "failed to connect to database")
	}
	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM todos WHERE id=$1`, id)
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusInternalServerError, err, "failed to delete todo")
	}

	_, err = res.RowsAffected()
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusInternalServerError, err, "failed to retrieve delete result")
	}

	return utils.HTTPSucess(ctx, res)
}

func (t *TodoController) Get(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusBadRequest, err, "invalid ID")
	}

	conn, err := db.OpenConnection()
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusInternalServerError, err, "failed to connect to database")
	}
	defer conn.Close()

	var data models.Todo
	row := conn.QueryRow(`SELECT uuid, title, description, done, in_progress, priority FROM todos WHERE id = $1`, id)
	err = row.Scan(&data.UUID, &data.Title, &data.Description, &data.Done, &data.InProgress, &data.Priority)
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusInternalServerError, err, "failed to retrieve todo")
	}

	return utils.HTTPSucess(ctx, data)
}

func (t *TodoController) GetAll(ctx echo.Context) error {
	conn, err := db.OpenConnection()
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusInternalServerError, err, "failed to connect to database: ")
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT uuid, title, description, done, in_progress, priority FROM todos`)
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusInternalServerError, err, "failed to retrieve todos: ")
	}
	defer rows.Close()

	var data []models.Todo
	for rows.Next() {
		var todo models.Todo
		err = rows.Scan(&todo.UUID, &todo.Title, &todo.Description, &todo.Done, &todo.InProgress, &todo.Priority)
		if err != nil {
			return utils.HTTPFail(ctx, http.StatusInternalServerError, err, "failed to scan todo: ")
		}
		data = append(data, todo)
	}

	return utils.HTTPSucess(ctx, data)
}
