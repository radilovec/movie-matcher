package repositrory

import (
	"context"
	"movies-service/internal/model"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repository struct {
	Movie
}

type Movie interface {
	GetById(ctx context.Context, id string) (model.Movie, error)
	GetByName(ctx context.Context, name string) (model.Movie, error)
	GetByCategory(ctx context.Context, category model.MovieCollectionType) ([]model.Movie, error)
	GetRandom(ctx context.Context, limit int) ([]model.Movie, error)
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		Movie: NewMovieRepository(db),
	}
}
