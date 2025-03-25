package repository

import (
	"context"
	"server/internal/models"
	"server/logger"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func SaveMovies(db *mongo.Database, category models.MovieCollectionType, movies []models.Movie) error {
	coll := db.Collection("movie_" + string(category))

	var operations []mongo.WriteModel

	for _, movie := range movies {
		filter := bson.M{"_id": movie.ID}
		update := bson.M{"$set": movie}

		operations = append(operations, mongo.NewUpdateManyModel().SetFilter(filter).SetUpdate(update))

		if len(operations) > 0 {
			_, err := coll.BulkWrite(context.Background(), operations)

			if err != nil {
				logger.LogError(err.Error())
				return err
			}
		}
		logger.LogInfo("записаны в бд")

	}

	return nil
}
