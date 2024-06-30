package router

import (
	"app/internal/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoute(r *gin.RouterGroup, authController *controllers.AuthController) *gin.RouterGroup {
	route := r.Group("/auth")
	{
		route.GET("/", authController.FinAll)
	}

	return r
}
