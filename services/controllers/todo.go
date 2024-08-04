package controllers

import (
	"api-postgresql/db"
	"api-postgresql/models"
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

func (t *TodoController) Create(c echo.Context) error {
	var todo models.Todo
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, models.HTTPErrorResponse{ErrorMessage: "Invalid request payload"})
	}

	if err := t.validator.Struct(todo); err != nil {
		return c.JSON(http.StatusBadRequest, models.HTTPErrorResponse{ErrorMessage: "Validation failed: " + err.Error()})
	}

	todo.UUID = uuid.New()

	conn, err := db.OpenConnection()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HTTPErrorResponse{ErrorMessage: "Failed to connect to database: " + err.Error()})
	}
	defer conn.Close()

	_, err = conn.Exec(`INSERT INTO todos (uuid, title, description, done, in_progress, priority) VALUES ($1, $2, $3, $4, $5, $6)`, todo.UUID, todo.Title, todo.Description, todo.Done, todo.InProgress, todo.Priority)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HTTPErrorResponse{ErrorMessage: "Failed to create todo: " + err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Todo created successfully"})
}

func (t *TodoController) Update(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.HTTPErrorResponse{ErrorMessage: "Invalid ID"})
	}

	var todo models.Todo
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, models.HTTPErrorResponse{ErrorMessage: "Invalid request payload"})
	}

	if err := t.validator.Struct(todo); err != nil {
		return c.JSON(http.StatusBadRequest, models.HTTPErrorResponse{ErrorMessage: "Validation failed: " + err.Error()})
	}

	conn, err := db.OpenConnection()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HTTPErrorResponse{ErrorMessage: "Failed to connect to database: " + err.Error()})
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE todos SET title=$1, description=$2, done=$3, in_progress=$4, priority=$5 WHERE id=$6`,
		todo.Title, todo.Description, todo.Done, todo.InProgress, todo.Priority, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HTTPErrorResponse{ErrorMessage: "Failed to update todo: " + err.Error()})
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HTTPErrorResponse{ErrorMessage: "Failed to retrieve update result: " + err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]int64{"updated": rowsAffected})
}

func (t *TodoController) Delete(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.HTTPErrorResponse{ErrorMessage: "Invalid ID"})
	}

	conn, err := db.OpenConnection()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HTTPErrorResponse{ErrorMessage: "Failed to connect to database: " + err.Error()})
	}
	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM todos WHERE id=$1`, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HTTPErrorResponse{ErrorMessage: "Failed to delete todo: " + err.Error()})
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HTTPErrorResponse{ErrorMessage: "Failed to retrieve delete result: " + err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]int64{"deleted": rowsAffected})
}

func (t *TodoController) Get(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.HTTPErrorResponse{ErrorMessage: "Invalid ID"})
	}

	conn, err := db.OpenConnection()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HTTPErrorResponse{ErrorMessage: "Failed to connect to database"})
	}
	defer conn.Close()

	var todo models.Todo
	row := conn.QueryRow(`SELECT uuid, title, description, done, in_progress, priority FROM todos WHERE id = $1`, id)
	err = row.Scan(&todo.UUID, &todo.Title, &todo.Description, &todo.Done, &todo.InProgress, &todo.Priority)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HTTPErrorResponse{ErrorMessage: "Failed to retrieve todo"})
	}

	return c.JSON(http.StatusOK, todo)
}

func (t *TodoController) GetAll(c echo.Context) error {
	conn, err := db.OpenConnection()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HTTPErrorResponse{ErrorMessage: "Failed to connect to database: " + err.Error()})
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT uuid, title, description, done, in_progress, priority FROM todos`)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HTTPErrorResponse{ErrorMessage: "Failed to retrieve todos: " + err.Error()})
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		err = rows.Scan(&todo.UUID, &todo.Title, &todo.Description, &todo.Done, &todo.InProgress, &todo.Priority)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, models.HTTPErrorResponse{ErrorMessage: "Failed to scan todo: " + err.Error()})
		}
		todos = append(todos, todo)
	}

	return c.JSON(http.StatusOK, todos)
}
