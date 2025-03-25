package main

import (
	"server/config"
	"server/internal/database"
	"server/internal/models"
	"server/internal/repository"
	"server/internal/services"
	"server/logger"
	"sync"
)

func main() {
	logger.InitLogger()
	config.InitConfig()

	db := database.ConnectDB()
	categories := []models.MovieCollectionType{models.NowPlaying, models.Popular, models.TopRated, models.Upcoming}

	movieChans := make(map[models.MovieCollectionType]chan []models.Movie)

	for _, category := range categories {
		movieChans[category] = make(chan []models.Movie, 100)
	}

	var wg sync.WaitGroup
	for _, category := range categories {
		wg.Add(1)
		go func(category models.MovieCollectionType) {
			defer wg.Done()
			services.FetchAllMovies(category, movieChans[category])
			repository.SaveMovies(db, category, <-movieChans[category])
		}(category)
	}

	wg.Wait()
	logger.LogInfo("фильмы загружены и записаны в монгу")

}
