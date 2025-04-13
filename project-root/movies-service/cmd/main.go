package main

import (
	"movies-service/config"
	"movies-service/internal/database"
	"movies-service/internal/server"
	"movies-service/internal/utils/logger"
)

func main() {
	logger.InitFileLogger()
	defer logger.SyncLogger()

	config.InitConfig()
	cfg := config.GetConfig()
	database.NewMongo()

	srv := new(server.Server)
	if err := srv.Run(cfg.Port); err != nil {
		logger.Error(err.Error())
		return
	}
}
