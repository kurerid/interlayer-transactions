package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/models"
)

func (h *Handler) registerAccount(c *gin.Context) {
	var input models.RegisterAccountInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	output, err := h.Usecase.RegisterAccount(input.Email, input.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, output)
}
