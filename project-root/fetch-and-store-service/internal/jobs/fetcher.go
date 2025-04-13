package jobs

import (
	"movies-service/internal/domain/models"
	"movies-service/internal/repository"
	"movies-service/internal/services"
	"movies-service/pkg/logger"
	"strconv"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

func FetchAndStoreMovies(db *mongo.Database) {
	categories := []models.MovieCollectionType{models.NowPlaying, models.Popular, models.TopRated, models.Upcoming}
	unqMovies := make(map[int]models.Movie)
	mu := sync.Mutex{}

	var wg sync.WaitGroup
	for _, category := range categories {
		_, totalPages := services.GetTotalPages(category)

		wg.Add(1)
		go func(category models.MovieCollectionType) {
			defer wg.Done()
			for page := 1; page <= totalPages; page++ {
				time.Sleep(200 * time.Millisecond)
				movies, err := services.FetchMovies(category, page)
				logger.LogInfo("FETCHED " + string(category) + " " + strconv.Itoa(page))

				if err != nil {
					logger.LogError(err.Error())
					return
				}

				mu.Lock()
				for _, movie := range movies {
					if _, exists := unqMovies[movie.ID]; !exists {
						unqMovies[movie.ID] = movie
					}
				}
				mu.Unlock()

				repository.SaveMovies(db, category, movies)
				logger.LogInfo("SAVED " + string(category) + " " + strconv.Itoa(page))
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
