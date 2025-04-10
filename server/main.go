package main

import (
	"server/config"
	"server/internal/database"
	"server/internal/jobs"
	"server/logger"
	"time"
)

func main() {
	logger.InitLogger()
	config.InitConfig()

	fcnfg := config.GetFetchConfig()

	db := database.ConnectDB()

	jobs.FetchAndStoreMovies(db)
	ticker := time.NewTicker(time.Second * time.Duration(fcnfg.Interval))
	defer ticker.Stop()

	for range ticker.C {
		go jobs.FetchAndStoreMovies(db)
		logger.LogInfo("success")
	}
}
