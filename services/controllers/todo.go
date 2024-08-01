package controllers

import (
	"api-postgresql/db"
	"api-postgresql/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TodoController struct{}

func NewTodoController() *TodoController {
	return &TodoController{}
}

func (t *TodoController) Create(c echo.Context) error {
	var todo models.Todo
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, models.HTTPErrorResponse{ErrorMessage: "Invalid request payload"})
	}

	conn, err := db.OpenConnection()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HTTPErrorResponse{ErrorMessage: "Failed to connect to database"})
	}
	defer conn.Close()

	var id int
	err = conn.QueryRow(`INSERT INTO todos (title, description, done, in_progress, priority) VALUES ($1, $2, $3, $4, $5) RETURNING id`, todo.Title, todo.Description, todo.Done, todo.InProgress, todo.Priority).Scan(&id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HTTPErrorResponse{ErrorMessage: "Failed to create todo"})
	}

	return c.JSON(http.StatusCreated, map[string]int{"id": id})
}

func (t *TodoController) Update(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.HTTPErrorResponse{ErrorMessage: "Invalid ID"})
	}

	var todo models.Todo
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, models.HTTPErrorResponse{ErrorMessage: "Invalid request payload"})
	}

	conn, err := db.OpenConnection()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HTTPErrorResponse{ErrorMessage: "Failed to connect to database"})
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE todos SET title=$1, description=$2, done=$3, in_progress=$4, priority=$5 WHERE id=$6`,
		todo.Title, todo.Description, todo.Done, todo.InProgress, todo.Priority, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HTTPErrorResponse{ErrorMessage: "Failed to update todo"})
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HTTPErrorResponse{ErrorMessage: "Failed to retrieve update result"})
	}

	return c.JSON(http.StatusOK, map[string]int64{"updated": rowsAffected})
}

func (t *TodoController) Delete(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.HTTPErrorResponse{ErrorMessage: "Invalid ID"})
	}

	conn, err := db.OpenConnection()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HTTPErrorResponse{ErrorMessage: "Failed to connect to database"})
	}
	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM todos WHERE id=$1`, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HTTPErrorResponse{ErrorMessage: "Failed to delete todo"})
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HTTPErrorResponse{ErrorMessage: "Failed to retrieve delete result"})
	}

	return c.JSON(http.StatusOK, map[string]int64{"deleted": rowsAffected})
}

func (t *TodoController) Get(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.HTTPErrorResponse{ErrorMessage: "Invalid ID"})
	}

	conn, err := db.OpenConnection()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HTTPErrorResponse{ErrorMessage: "Failed to connect to database"})
	}
	defer conn.Close()

	var todo models.Todo
	row := conn.QueryRow(`SELECT id, title, description, done, in_progress, priority FROM todos WHERE id = $1`, id)
	err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done, &todo.InProgress, &todo.Priority)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HTTPErrorResponse{ErrorMessage: "Failed to retrieve todo"})
	}

	return c.JSON(http.StatusOK, todo)
}

func (t *TodoController) GetAll(c echo.Context) error {
	conn, err := db.OpenConnection()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HTTPErrorResponse{ErrorMessage: "Failed to connect to database"})
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT id, title, description, done, in_progress, priority FROM todos`)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HTTPErrorResponse{ErrorMessage: "Failed to retrieve todos"})
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done, &todo.InProgress, &todo.Priority)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, models.HTTPErrorResponse{ErrorMessage: "Failed to scan todo"})
		}
		todos = append(todos, todo)
	}

	return c.JSON(http.StatusOK, todos)
}