package db

import (
	"api-postgresql/configs"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // Importar o driver PostgreSQL
	"go.uber.org/zap"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDBConfig()
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Info("Opening database connection",
		zap.String("host", conf.Host),
		zap.String("port", conf.Port),
		zap.String("user", conf.User),
		zap.String("dbname", conf.Database),
	)

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	logger.Info("Connection string", zap.String("connection_string", sc))

	conn, err := sql.Open("postgres", sc)
	if err != nil {
		logger.Error("Failed to open database connection", zap.Error(err))
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		logger.Error("Failed to ping database", zap.Error(err))
		return nil, err
	}

	logger.Info("Successfully connected to the database")
	return conn, nil
}
