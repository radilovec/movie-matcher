package config

import (
	"movies-service/pkg/logger"
	"os"
	"strconv"

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

func GetFetchConfig() FetchConfig {

	interval, err := strconv.Atoi(os.Getenv("INTERVAL"))
	if err != nil {
		interval = 60
	}

	return FetchConfig{
		Interval: interval,
	}
}
