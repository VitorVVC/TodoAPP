package handlers

import (
	"api-postgresql/models"
	"api-postgresql/services/controllers"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func Delete(c echo.Context, dbConfig *models.DBConfig) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.HTTPErrorResponse{
			Message: "Invalid ID",
		})
	}

	rows, err := controllers.Delete(dbConfig, int64(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HTTPErrorResponse{
			ErrorMessage: err.Error(),
			Message:      "Failed to delete todo",
		})
	}

	if rows == 0 {
		return c.JSON(http.StatusNotFound, models.HTTPErrorResponse{
			Message: "Todo not found",
		})
	}

	return c.JSON(http.StatusOK, models.HTTPResponse{
		Data: models.DeleteTodoResponse{
			Message: "Todo deleted successfully",
		},
	})
}
