package routes

import (
	"api-postgresql/constants"
	"api-postgresql/services/controllers"
	"github.com/labstack/echo/v4"
)

func Todo(app *echo.Group) {
	group := app.Group(string(constants.RootRoute))
	controller := controllers.NewTodoController()

	group.POST("", controller.Create)
	group.PUT("/:id", controller.Update)
	group.DELETE("/:id", controller.Delete)
	group.GET("", controller.GetAll)
	group.GET("/:id", controller.Get)
}
