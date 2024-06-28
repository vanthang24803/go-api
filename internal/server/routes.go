package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.GET("/", s.GetHelloWorld)

	return r
}

func (s *Server) GetHelloWorld(c *gin.Context) {
	resp := make(map[string]string)

	resp["message"] = "Hello World Go API"

	c.JSON(http.StatusOK, resp)
}
