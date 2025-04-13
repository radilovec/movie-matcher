package database

import (
	"context"
	"movies-service/config"
	"movies-service/pkg/logger"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func ConnectToMongo() *mongo.Database {
	cfg := config.GetDBConfig()
	docs := "www.mongodb.com/docs/drivers/go/current/"

	if cfg.MongoURI == "" {
		logger.LogFatal("Set your 'MONGODB_URI' environment variable. " +
			"See: " + docs +
			"usage-examples/#environment-variable")
	}

	client, err := mongo.Connect(options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		logger.LogError(err.Error())
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		logger.LogError(err.Error())
		panic(err)
	}

	logger.LogInfo("Connected successfully")
	return client.Database(cfg.DBName)
}
