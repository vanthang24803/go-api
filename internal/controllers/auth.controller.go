package controllers

import (
	"app/internal/repositories"

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
	result := controller.authRepo.GetAll()

	c.JSON(result.Status, result)
}
