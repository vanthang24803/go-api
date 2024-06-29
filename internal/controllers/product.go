package controllers

import (
	"app/pkg/helpers"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ProductController(r *gin.RouterGroup) {
	products := r.Group("/products")
	{
		products.GET("/", GetProducts)
	}
}

func GetProducts(c *gin.Context) {
	products := os.Getenv("PORT")
	helpers.JSON(c, http.StatusOK, products)
}
