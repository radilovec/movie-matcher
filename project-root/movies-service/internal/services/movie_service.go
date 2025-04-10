package services

import (
	"encoding/json"
	"fmt"
	"movies-service/config"
	"movies-service/internal/domain/models"
	"movies-service/pkg/logger"

	"net/http"
)

func FetchMovies(category models.MovieCollectionType, page int) ([]models.Movie, error) {
	if !category.IsValid() {
		err := fmt.Errorf("invalid movie collection type")
		return nil, err
	}

	tmdbCfg := config.GetTmdbConfig()
	url := fmt.Sprintf("https://api.themoviedb.org/3/movie/%s?page=%d&api_key=%s", string(category), page, tmdbCfg.ApiKey)
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
