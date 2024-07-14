package handlers

import (
	"api-postgresql/models"
	"api-postgresql/services/controllers"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetAll(c echo.Context, dbConfig *models.DBConfig) error {
	todos, err := controllers.GetAll(dbConfig)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HTTPErrorResponse{
			ErrorMessage: err.Error(),
			Message:      "Failed to get todos",
		})
	}

	return c.JSON(http.StatusOK, models.HTTPResponse{
		Data: todos,
	})
}