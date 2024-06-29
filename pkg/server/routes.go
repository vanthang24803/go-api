package server

import (
	"net/http"

	v1 "app/pkg/router/v1"

	"github.com/gin-gonic/gin"
)

func (s *Server) Routes() http.Handler {
	r := gin.Default()

	v1.Routes(r.Group("/api/v1"))

	return r
}
