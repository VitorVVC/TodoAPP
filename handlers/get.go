package handlers

import (
	"api-postgresql/models"
	"api-postgresql/services/controllers"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func Get(c echo.Context, dbConfig *models.DBConfig) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response := models.HTTPErrorResponse{
			Message: "Invalid ID",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	todo, err := controllers.Get(dbConfig, int64(id))
	if err != nil {
		response := models.HTTPErrorResponse{
			ErrorMessage: err.Error(),
			Message:      "Failed to get todo",
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := models.HTTPResponse{
		Data: todo,
	}
	return c.JSON(http.StatusOK, response)
}
