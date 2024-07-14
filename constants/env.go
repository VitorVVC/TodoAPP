package constants

import "os"

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

var (
	ApiPort      = getEnv("API_PORT", "9000")
	PostgresHost = getEnv("POSTGRES_HOST", "localhost")
	PostgresPort = getEnv("POSTGRES_PORT", "5432")
	PostgresUser = getEnv("POSTGRES_USER", "user")
	PostgresPass = getEnv("POSTGRES_PASS", "pass")
	PostgresName = getEnv("POSTGRES_NAME", "dbname")
)

const (
	FutureConst = "Hello World!"
)
