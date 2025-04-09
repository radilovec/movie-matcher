package utils

import (
	"fmt"
	"server/internal/models"
	"server/internal/repository"
	"server/logger"
	"sync"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

func CreateUniqueColl(db *mongo.Database, collections []models.MovieCollectionType) ([]models.Movie, error) {
	var wg sync.WaitGroup
	mu := sync.Mutex{}
	uniqueMovies := make(map[int]models.Movie)

	for _, coll := range collections {
		wg.Add(1)

		go func(coll models.MovieCollectionType) {
			defer wg.Done()

			movies, err := repository.GetMoviesByColl(db, coll)
			if err != nil {
				logger.LogError(fmt.Sprintf("Error fetching from collection %s: %v", coll, err))
				return
			}

			mu.Lock()
			for _, movie := range movies {
				if _, exists := uniqueMovies[movie.ID]; !exists {
					uniqueMovies[movie.ID] = movie
				}
			}
			mu.Unlock()
		}(coll)
	}

	wg.Wait()

	result := make([]models.Movie, 0, len(uniqueMovies))

	for _, movie := range uniqueMovies {
		result = append(result, movie)
	}

	return result, nil
}
