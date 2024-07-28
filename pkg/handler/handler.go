package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.GET("/")
		}
	}
	return router
}
