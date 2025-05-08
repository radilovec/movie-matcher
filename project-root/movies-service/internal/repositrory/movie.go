package repositrory

import (
	"context"
	"movies-service/internal/model"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MovieRepository struct {
	DB *mongo.Database
}

func NewMovieRepository(db *mongo.Database) *MovieRepository {
	return &MovieRepository{
		DB: db,
	}
}

func (r *MovieRepository) GetByCategory(ctx context.Context, category model.MovieCollectionType) ([]model.Movie, error) {
	collection := r.DB.Collection(string(category))

	var movies []model.Movie
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var movie model.Movie
		if err := cursor.Decode(&movie); err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	return movies, nil
}

func (r *MovieRepository) GetById(ctx context.Context, id string) (model.Movie, error) {
	collection := r.DB.Collection("unique")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return model.Movie{}, err
	}

	var movie model.Movie
	err = collection.FindOne(ctx, bson.M{"id": intID}).Decode(&movie)
	return movie, err
}

func (r *MovieRepository) GetByName(ctx context.Context, name string) (model.Movie, error) {
	collection := r.DB.Collection("unique")
	var movie model.Movie

	filter := bson.M{
		"title": bson.M{
			"$regex":   name,
			"$options": "i",
		},
	}

	err := collection.FindOne(ctx, filter).Decode(&movie)
	if err != nil {
		return model.Movie{}, err
	}

	return movie, nil
}

func (r *MovieRepository) GetRandom(ctx context.Context, limit int) ([]model.Movie, error) {
	collection := r.DB.Collection("unique")

	pipeline := []bson.M{
		{"$sample": bson.M{"size": limit}},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var movies []model.Movie
	for cursor.Next(ctx) {
		var movie model.Movie
		if err := cursor.Decode(&movie); err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	return movies, nil
}
