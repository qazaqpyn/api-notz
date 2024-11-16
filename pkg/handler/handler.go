package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/qazaqpyn/api-notz/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.GET("/refresh", h.refreshTokens)
	}

	api := router.Group("/api")
	{
		api.Use(h.authorizarion)
		notes := api.Group("/notes")
		{
			notes.GET("/", h.getAllNotes)
			notes.POST("/", h.createNote)
			notes.GET("/:id", h.getNoteById)
			notes.PUT("/:id", h.updateNote)
			notes.DELETE("/:id", h.deleteNote)
		}
	}

	return router
}
