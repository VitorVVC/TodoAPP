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

	e.POST("/", func(c echo.Context) error {
		return handlers.Create(c, &cfg.DB)
	})
	e.PUT("/:id", func(c echo.Context) error {
		return handlers.Update(c, &cfg.DB)
	})
	e.DELETE("/:id", func(c echo.Context) error {
		return handlers.Delete(c, &cfg.DB)
	})
	e.GET("/", func(c echo.Context) error {
		return handlers.GetAll(c, &cfg.DB)
	})
	e.GET("/:id", func(c echo.Context) error {
		return handlers.Get(c, &cfg.DB)
	})

	port := cfg.GetServerPort()
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
