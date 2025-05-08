package config

import (
	"auth-service/pkg/logger"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	MongoURI string
	DBName   string
}

type TmdbConfig struct {
	ApiKey  string
	BaseURL string
}

type FetchConfig struct {
	Interval int
}

func InitConfig() {
	if err := godotenv.Load(); err != nil {
		logger.LogError(err.Error())
		return
	}
	logger.LogInfo("Load .evn file")
}

func GetDBConfig() DBConfig {
	return DBConfig{
		MongoURI: os.Getenv("MONGO_URI"),
		DBName:   os.Getenv("DB_NAME"),
	}
}
