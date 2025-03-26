package repository

import (
	"context"
	"server/internal/models"
	"server/logger"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func SaveMovies(db *mongo.Database, category models.MovieCollectionType, movies []models.Movie) error {
	coll := db.Collection(string(category))

	for _, movie := range movies {
		filter := bson.M{"_id": movie.ID}
		update := bson.M{"$set": movie}
		opts := options.UpdateOne().SetUpsert(true) // Вставит, если нет

		_, err := coll.UpdateOne(context.Background(), filter, update, opts)
		if err != nil {
			logger.LogError("Failed to save movie [" + movie.Title + "]: " + err.Error())
			return err
		}

	}

	logger.LogInfo(string(category) + " movies uploaded in MongoDB")
	return nil
}
