package main

import (
	"movies-service/config"
	"movies-service/internal/database"
	"movies-service/internal/handler"
	"movies-service/internal/repositrory"
	"movies-service/internal/server"
	"movies-service/internal/service"
	"movies-service/internal/utils/logger"
)

func main() {
	logger.InitFileLogger()
	defer logger.SyncLogger()

	config.InitConfig()
	cfg := config.GetConfig()
	db := database.NewMongo()
	repositrory := repositrory.NewRepository(db)
	service := service.NewService(repositrory)
	handlers := handler.NewHandler(service)

	srv := new(server.Server)
	if err := srv.Run(cfg.Port, handlers.Init()); err != nil {
		logger.Error(err.Error())
		return
	}

}
