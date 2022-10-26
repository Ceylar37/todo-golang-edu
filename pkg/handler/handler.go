package handler

import (
	"crud/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.LoadHTMLGlob("templates/*")

	api := router.Group("/api")
	{
		api.GET("/", h.getAllTodos)
		api.POST("/", h.createTodo)
		api.PUT("/changeIsDone/:id", h.changeIsDone)
		api.PUT("/changeIsFavourite/:id", h.changeIsFavourite)
		api.DELETE("/:id", h.deleteTodo)
		api.GET("/docs", h.getDocs)
	}

	return router
}
