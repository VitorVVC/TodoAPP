package utils

import (
	"api-postgresql/constants"
	"go.uber.org/zap"
	"os"
)

var Defaults = map[string]interface{}{
	constants.ApiPort:      "your_port",
	constants.PostgresHost: "your_host",
	constants.PostgresPort: "5432",
	constants.PostgresUser: "your_user",
	constants.PostgresPass: "your_pass",
	constants.PostgresName: "your_name",
}

func EnvString(key string) string {
	value := os.Getenv(key)
	if value == "" {
		valueInterface, ok := Defaults[key]
		if !ok {
			zap.L().Fatal("missing env", zap.Error(ErrMissingEnv), zap.String("env", key))
		}

		value, ok = valueInterface.(string)
		if !ok {
			zap.L().Fatal("wrong type", zap.Error(ErrWrongEnvType), zap.String("env", key))
		}
	}

	return value
}
