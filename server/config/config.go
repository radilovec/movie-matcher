package config

import (
	"os"
	"server/logger"

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

func InitConfig() {
	if err := godotenv.Load(); err != nil {
		logger.LogError(err.Error())
		return
	}
	logger.LogInfo("Load .evn file")
}

func GetTmdbConfig() TmdbConfig {

	return TmdbConfig{
		ApiKey:  os.Getenv("TMDB_API_KEY"),
		BaseURL: os.Getenv("TMDB_BASE_URL"),
	}
}

func GetDBConfig() DBConfig {
	return DBConfig{
		MongoURI: os.Getenv("MONGO_URI"),
		DBName:   os.Getenv("DB_NAME"),
	}
}
