package handlers

import (
	"api-postgresql/models"
	"api-postgresql/services/controllers"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Create(c echo.Context) error {
	var todo models.Todo

	err := json.NewDecoder(c.Request().Body).Decode(&todo)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.HTTPErrorResponse{
			Message: "Invalid request payload",
		})
	}

	id, err := controllers.Create(todo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HTTPErrorResponse{
			ErrorMessage: err.Error(),
			Message:      "Failed to create todo",
		})
	}

	return c.JSON(http.StatusCreated, models.HTTPResponse{
		Data: models.CreateTodoResponse{
			ID: id,
		},
	})
}
