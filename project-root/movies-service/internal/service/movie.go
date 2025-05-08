package service

import (
	"context"
	"fmt"
	"movies-service/internal/model"
	"movies-service/internal/repositrory"
)

type MovieService struct {
	Repository repositrory.Movie
}

func NewMovieService(repo repositrory.Movie) *MovieService {
	return &MovieService{
		Repository: repo,
	}
}

func (s *MovieService) GetById(ctx context.Context, id string) (model.Movie, error) {
	return s.Repository.GetById(ctx, id)
}

func (s *MovieService) GetByName(ctx context.Context, name string) (model.Movie, error) {
	return s.Repository.GetByName(ctx, name)
}

func (s *MovieService) GetByCategory(ctx context.Context, category model.MovieCollectionType) ([]model.Movie, error) {
	if !category.IsValid() {
		return nil, fmt.Errorf("invalid category")
	}

	return s.Repository.GetByCategory(ctx, category)
}

func (s *MovieService) GetRandom(ctx context.Context, limit int) ([]model.Movie, error) {
	return s.Repository.GetRandom(ctx, limit)
}
