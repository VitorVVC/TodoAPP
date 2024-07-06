package handlers

import (
	"api-postgresql/models"
	"api-postgresql/services/controllers"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func Update(c echo.Context, dbConfig *models.DBConfig) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.HTTPErrorResponse{
			Message: "Invalid ID",
		})
	}

	var todo models.Todo
	err = json.NewDecoder(c.Request().Body).Decode(&todo)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.HTTPErrorResponse{
			Message: "Invalid request payload",
		})
	}

	rows, err := controllers.Update(dbConfig, int64(id), todo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HTTPErrorResponse{
			ErrorMessage: err.Error(),
			Message:      "Failed to update todo",
		})
	}

	if rows == 0 {
		return c.JSON(http.StatusNotFound, models.HTTPErrorResponse{
			Message: "Todo not found",
		})
	}

	return c.JSON(http.StatusOK, models.HTTPResponse{
		Data: models.UpdateTodoResponse{
			Message: "Todo updated successfully",
		},
	})
}
