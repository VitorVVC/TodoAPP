package main

import (
	"api-postgresql/configs"
	"api-postgresql/handlers"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg, err := configs.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Erro ao carregar a configuração: %v", err))
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	e.POST("/", func(c echo.Context) error {
		return handlers.Create(c)
	})
	e.PUT("/:id", func(c echo.Context) error {
		return handlers.Update(c)
	})
	e.DELETE("/:id", func(c echo.Context) error {
		return handlers.Delete(c)
	})
	e.GET("/", func(c echo.Context) error {
		return handlers.GetAll(c)
	})
	e.GET("/:id", func(c echo.Context) error {
		return handlers.Get(c)
	})

	port := cfg.API.Port
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
