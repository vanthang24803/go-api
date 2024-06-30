package controllers

import (
	"app/internal/repositories"
	"app/pkg/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authRepo repositories.AuthRepository
}

func NewAuthController(repository repositories.AuthRepository) *AuthController {
	return &AuthController{
		authRepo: repository,
	}
}

func (controller *AuthController) FinAll(c *gin.Context) {

	result, err := controller.authRepo.GetAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	helpers.JSON(c, http.StatusOK, result)

}
