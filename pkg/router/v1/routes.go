package router

import (
	"github.com/gin-gonic/gin"

	routes "app/internal/controllers"
)

func Routes(r *gin.RouterGroup) {
	v1 := r.Group("/api/v1")
	{
		routes.ProductController(v1)
	}
}
