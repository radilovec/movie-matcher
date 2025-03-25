package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/config"
	"server/internal/models"
	"server/logger"
	"sync"
	"time"
)

func FetchMovies(category models.MovieCollectionType, page int) ([]models.Movie, error) {
	if !category.IsValid() {
		err := fmt.Errorf("invalid movie collection type")
		return nil, err
	}

	tmdbCfg := config.GetTmdbConfig()
	url := fmt.Sprintf("https://api.themoviedb.org/3/movie/%s?page=%d&api_key=%s", category, page, tmdbCfg.ApiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var data models.TmdbResponse

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	logger.LogInfo(string(category) + " movies fetched successfully")
	return data.Results, nil
}

func FetchAllMovies(category models.MovieCollectionType, movieChan chan<- []models.Movie) {
	var wg sync.WaitGroup

	for page := 1; page <= 5; page++ {
		wg.Add(1)

		go func(page int) {
			defer wg.Done()
			movies, err := FetchMovies(category, page)

			if err != nil || movies == nil {
				logger.LogError(err.Error())
				return
			}
			time.Sleep(300 * time.Millisecond)

			movieChan <- movies
		}(page)

	}

	wg.Wait()
	close(movieChan)
}
