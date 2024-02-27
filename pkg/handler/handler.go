package handler

import (
	"github.com/gin-gonic/gin"
	"main.go/pkg/usecase"
)

type Handler struct {
	Usecase *usecase.Usecase
}

func NewHandler(usecase *usecase.Usecase) *Handler {
	return &Handler{
		Usecase: usecase,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.POST("/api/person", h.createPerson)

	return router
}
