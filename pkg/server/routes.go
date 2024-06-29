package server

import (
	routes "app/internal/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) Routes() http.Handler {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		routes.ProductController(v1)
		routes.AuthController(v1)
	}
	return r
}
