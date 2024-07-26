package configs

import (
	"api-postgresql/constants"
	"api-postgresql/models"
	"api-postgresql/utils"
	"log"
)

func LoadConfig() (*models.Config, error) {
	cfg := &models.Config{
		API: models.APIConfig{
			Port: utils.EnvString(constants.ApiPort),
		},
		DB: models.DBConfig{
			Host:     utils.EnvString(constants.PostgresHost),
			Port:     utils.EnvString(constants.PostgresPort),
			User:     utils.EnvString(constants.PostgresUser),
			Pass:     utils.EnvString(constants.PostgresPass),
			Database: utils.EnvString(constants.PostgresName),
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
