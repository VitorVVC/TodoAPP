package models

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

func (c *Config) GetDB() DBConfig {
	return c.DB
}

func (c *Config) GetServerPort() string {
	return c.API.Port
}
