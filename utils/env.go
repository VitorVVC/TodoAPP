package utils

import (
	"api-postgresql/constants"
	"github.com/spf13/viper"
	"log"
)

var Config = map[string]string{}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	Config[constants.API_PORT] = viper.GetString("api.port")
	Config[constants.POSTGRES_HOST] = viper.GetString("database.host")
	Config[constants.POSTGRES_PORT] = viper.GetString("database.port")
	Config[constants.POSTGRES_USER] = viper.GetString("database.user")
	Config[constants.POSTGRES_PASS] = viper.GetString("database.pass")
	Config[constants.POSTGRES_NAME] = viper.GetString("database.name")
}
