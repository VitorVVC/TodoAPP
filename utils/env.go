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

	Config[constants.ApiPort] = viper.GetString("api.port")
	Config[constants.PostgresHost] = viper.GetString("database.host")
	Config[constants.PostgresPort] = viper.GetString("database.port")
	Config[constants.PostgresUser] = viper.GetString("database.user")
	Config[constants.PostgresPass] = viper.GetString("database.pass")
	Config[constants.PostgresName] = viper.GetString("database.name")
}
