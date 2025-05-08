package handler

import (
	"movies-service/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		Services: services,
	}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	router.GET("/movies/:collection", h.getMoviesByCategory)
	router.GET("/movies/id", h.getMovieById)
	router.GET("/movies/name", h.getMovieByName)
	router.GET("/movies/random", h.getRandomMovies)

	return router
}
