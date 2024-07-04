package configs

import (
	"github.com/spf13/viper"
	"log"
)

var Cfg *Config

type Config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

func init() {
	viper.SetDefault("api.port", "8080")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
	viper.SetDefault("database.user", "user")
	viper.SetDefault("database.pass", "pass")
	viper.SetDefault("database.name", "dbname")

	err := Load()
	if err != nil {
		log.Fatalf("Erro ao carregar a configuração: %v", err)
	}
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	Cfg = &Config{
		API: APIConfig{
			Port: viper.GetString("api.port"),
		},
		DB: DBConfig{
			Host:     viper.GetString("database.host"),
			Port:     viper.GetString("database.port"),
			User:     viper.GetString("database.user"),
			Pass:     viper.GetString("database.pass"),
			Database: viper.GetString("database.name"),
		},
	}

	return nil
}

func GetDB() DBConfig {
	return Cfg.DB
}

func GetServerPort() string {
	return Cfg.API.Port
}
