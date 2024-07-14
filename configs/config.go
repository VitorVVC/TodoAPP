package configs

import (
	"api-postgresql/constants"
	"api-postgresql/models"
	"log"
)

func LoadConfig() (*models.Config, error) {
	cfg := &models.Config{
		API: models.APIConfig{
			Port: constants.ApiPort,
		},
		DB: models.DBConfig{
			Host:     constants.PostgresHost,
			Port:     constants.PostgresPort,
			User:     constants.PostgresUser,
			Pass:     constants.PostgresPass,
			Database: constants.PostgresName,
		},
	}

	return cfg, nil
}

func GetDBConfig() *models.DBConfig {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("Erro ao carregar a configuração: %v", err)
	}
	return &cfg.DB
}
