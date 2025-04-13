package database

import (
	"context"
	"movies-service/config"
	"movies-service/internal/utils/logger"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func NewMongo() *mongo.Database {
	cfg := config.GetDBConfig()

	client, err := mongo.Connect(options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		logger.Error(err.Error())
		panic(err)
	}

	logger.Info("Connected to MongoDB successfully")
	return client.Database(cfg.DBName)
}
