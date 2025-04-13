package handlers

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) Init() *gin.Engine {
	router := gin.New()

	movies := router.Group("/movies")
	{
		movies.GET("/:collection", h.getMoviesByCollection)
		{
			movies.GET("/:id", h.getMovieById)
			movies.POST("", h.createMovie)
			movies.PUT("/:id", h.updateMovie)
			movies.DELETE("/:id", h.deleteMovie)
		}

	}

	return router
}

func (h *Handler) getMoviesByCollection(c *gin.Context) {

}

func (h *Handler) getMovieById(c *gin.Context) {

}

func (h *Handler) createMovie(c *gin.Context) {}

func (h *Handler) updateMovie(c *gin.Context) {}

func (h *Handler) deleteMovie(c *gin.Context) {}
