package configs

import (
	"api-postgresql/constants"
	"api-postgresql/models"
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault(constants.API_PORT, constants.DefaultAPIPort)
	viper.SetDefault(constants.POSTGRES_HOST, constants.DefaultPostgresHost)
	viper.SetDefault(constants.POSTGRES_PORT, constants.DefaultPostgresPort)
	viper.SetDefault(constants.POSTGRES_USER, constants.DefaultPostgresUser)
	viper.SetDefault(constants.POSTGRES_PASS, constants.DefaultPostgresPass)
	viper.SetDefault(constants.POSTGRES_NAME, constants.DefaultPostgresName)
}

func LoadConfig() (*models.Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}

	cfg := &models.Config{
		API: models.APIConfig{
			Port: viper.GetString(constants.API_PORT),
		},
		DB: models.DBConfig{
			Host:     viper.GetString(constants.POSTGRES_HOST),
			Port:     viper.GetString(constants.POSTGRES_PORT),
			User:     viper.GetString(constants.POSTGRES_USER),
			Pass:     viper.GetString(constants.POSTGRES_PASS),
			Database: viper.GetString(constants.POSTGRES_NAME),
		},
	}

	return cfg, nil
}
