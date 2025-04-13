package repository

import (
	"context"
	"fetch-and-store/internal/domain/models"
	"fetch-and-store/pkg/logger"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func SaveMovies(db *mongo.Database, category models.MovieCollectionType, movies []models.Movie) error {
	coll := db.Collection(string(category))

	for _, movie := range movies {
		filter := bson.M{"_id": movie.ID}
		update := bson.M{"$set": movie}
		opts := options.UpdateOne().SetUpsert(true)
		_, err := coll.UpdateOne(context.Background(), filter, update, opts)
		if err != nil {
			logger.LogError(err.Error())
			return err
		}

	}

	return nil
}

func GetMoviesByColl(db *mongo.Database, category models.MovieCollectionType) ([]models.Movie, error) {
	coll := db.Collection(string(category))

	cursor, err := coll.Find(context.Background(), bson.D{})

	if err != nil {
		logger.LogError(err.Error())
		return nil, err
	}
	defer cursor.Close(context.Background())

	movies := make([]models.Movie, 0)

	if err = cursor.All(context.Background(), &movies); err != nil {
		logger.LogError(err.Error())
		return nil, err
	}

	return movies, nil
}
