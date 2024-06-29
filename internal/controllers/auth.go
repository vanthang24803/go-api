package controllers

import (
	"app/internal/models"
	"app/pkg/database"
	"app/pkg/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AuthController(r *gin.RouterGroup) {
	route := r.Group("/auth")
	{
		route.GET("/", GetHelloWorld)
		route.GET("/create", Save)
	}
}

func Save(c *gin.Context) {

	user := models.User{
		ID:        uuid.New(),
		FirstName: "May",
		LastName:  "Nguyen",
		Avatar:    "https://i.pinimg.com/564x/96/ff/11/96ff113f9df327237fa1d1bb7f2d3f22.jpg",
	}

	result := database.DB.Create(&user)

	if result.Error != nil {
		helpers.JSON(c, http.StatusInternalServerError, result.Error.Error())
		return
	}

	helpers.JSON(c, http.StatusCreated, user)
}

func GetHelloWorld(c *gin.Context) {
	helpers.JSON(c, http.StatusOK, "Hello World Auth Route")
}
