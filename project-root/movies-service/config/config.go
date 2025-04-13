package config

import (
	"movies-service/internal/utils/logger"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI string
	DBName   string
	Port     string
}

type DBConfig struct {
	MongoURI string
	DBName   string
}

func InitConfig() {
	if err := godotenv.Load("../.env"); err != nil {
		logger.Error(err.Error())
		return
	}
	logger.Info("Load .env file")
}

func GetConfig() Config {
	return Config{
		MongoURI: os.Getenv("MONGO_URI"),
		DBName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("PORT"),
	}
}

func GetDBConfig() DBConfig {
	return DBConfig{
		MongoURI: os.Getenv("MONGO_URI"),
		DBName:   os.Getenv("DB_NAME"),
	}
}
