package routes

import (
	"api-postgresql/constants"
	"api-postgresql/services/controllers"
	"github.com/labstack/echo/v4"
)

func User(app *echo.Group) {
	group := app.Group(string(constants.UserRoute))
	controller := controllers.NewUser()

	group.POST("", controller.Create)
	//group.PUT("/:id", controller.Update)
	//group.DELETE("/:id", controller.Delete)
	//group.GET("", controller.GetAll)
	//group.GET("/:id", controller.Get)
}
