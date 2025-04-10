package jobs

import (
	"server/config"
	"server/internal/models"
	"server/internal/repository"
	"server/internal/services"
	"server/logger"
	"sync"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

func FetchAndStoreMovies(db *mongo.Database) {
	cfg := config.GetFetchConfig()
	categories := []models.MovieCollectionType{models.NowPlaying, models.Popular, models.TopRated, models.Upcoming}
	unqMovies := make(map[int]models.Movie)
	mu := sync.Mutex{}

	var wg sync.WaitGroup
	for _, category := range categories {
		wg.Add(1)
		go func(category models.MovieCollectionType) {
			defer wg.Done()
			for page := 1; page <= cfg.TotalPages; page++ {
				movies, err := services.FetchMovies(category, page)

				if err != nil {
					logger.LogError(err.Error())
					panic(err)
				}

				mu.Lock()
				for _, movie := range movies {
					if _, exists := unqMovies[movie.ID]; !exists {
						unqMovies[movie.ID] = movie
					}
				}
				mu.Unlock()

				repository.SaveMovies(db, category, movies)
			}
		}(category)
	}

	wg.Wait()

	unqMoviesColl := make([]models.Movie, 0, len(unqMovies))

	for _, movie := range unqMovies {
		unqMoviesColl = append(unqMoviesColl, movie)
	}

	repository.SaveMovies(db, models.Unique, unqMoviesColl)
	logger.LogInfo("all the movies have been fetched and uploaded, including unique")
}
