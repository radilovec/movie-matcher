package main

import (
	"server/config"
	"server/internal/database"
	"server/internal/models"
	"server/internal/repository"
	"server/internal/services"
	"server/internal/utils"
	"server/logger"
	"sync"
)

func main() {
	const totalPages = 5
	logger.InitLogger()
	config.InitConfig()

	db := database.ConnectDB()
	categories := []models.MovieCollectionType{models.NowPlaying, models.Popular, models.TopRated, models.Upcoming}

	var wg sync.WaitGroup
	for _, category := range categories {
		wg.Add(1)
		go func(category models.MovieCollectionType) {
			defer wg.Done()
			for page := 1; page <= totalPages; page++ {
				movies, err := services.FetchMovies(category, page)

				if err != nil {
					logger.LogError(err.Error())
					panic(err)
				}

				repository.SaveMovies(db, category, movies)
			}
		}(category)
	}

	wg.Wait()
	logger.LogInfo("all the movies have been fetched and uploaded")

	uniqueMovies, _ := utils.CreateUniqueColl(db, categories)
	repository.SaveMovies(db, models.Unique, uniqueMovies)

}
