package service

import (
	"context"
	"movies-service/internal/model"
	"movies-service/internal/repositrory"
)

type Service struct {
	Movie
}

type Movie interface {
	GetById(ctx context.Context, id string) (model.Movie, error)
	GetByName(ctx context.Context, name string) (model.Movie, error)
	GetByCategory(ctx context.Context, category model.MovieCollectionType) ([]model.Movie, error)
	GetRandom(ctx context.Context, limit int) ([]model.Movie, error)
}

func NewService(repos *repositrory.Repository) *Service {
	return &Service{
		Movie: NewMovieService(repos.Movie),
	}
}
