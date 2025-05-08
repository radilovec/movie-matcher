package handler

import (
	"context"
	"movies-service/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getMoviesByCategory(c *gin.Context) {
	collectionStr := c.Param("collection")
	collection := model.MovieCollectionType(collectionStr)

	movies, err := h.Services.GetByCategory(context.Background(), collection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movies)
}

func (h *Handler) getMovieById(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	movie, err := h.Services.GetById(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if movie.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Movie not found"})
		return
	}

	c.JSON(http.StatusOK, movie)
}

func (h *Handler) getMovieByName(c *gin.Context) {
	name := c.Query("name")
	movie, err := h.Services.GetByName(context.Background(), name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movie)
}

func (h *Handler) getRandomMovies(c *gin.Context) {
	limit := 10
	movies, err := h.Services.GetRandom(context.Background(), limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movies)
}
