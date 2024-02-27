package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/models"
)

func (h *Handler) createPerson(c *gin.Context) {
	var input models.CreatePersonInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}

	output, err := h.Usecase.CreatePerson(input.Name, input.Age)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, output)
}
