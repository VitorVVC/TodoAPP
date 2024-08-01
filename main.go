package main

import (
	"api-postgresql/configs"
	"api-postgresql/routes"
	_ "api-postgresql/utils" // Importar utils para garantir que o init() seja chamado
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Erro ao carregar a configuração: %v", err)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	api := e.Group("")
	routes.Todo(api)

	port := cfg.API.Port
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
