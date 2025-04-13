package main

import (
	"fetch-and-store/config"
	"fetch-and-store/internal/jobs"
	"fetch-and-store/pkg/database"
	"fetch-and-store/pkg/logger"
	"time"
)

func main() {
	logger.InitLogger()
	config.InitConfig()

	fcnfg := config.GetFetchConfig()

	db := database.ConnectToMongo()

	jobs.FetchAndStoreMovies(db)
	ticker := time.NewTicker(time.Minute * time.Duration(fcnfg.Interval))
	defer ticker.Stop()

	for range ticker.C {
		go jobs.FetchAndStoreMovies(db)
		logger.LogInfo("success")
	}
}
