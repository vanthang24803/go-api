package server

import (
	"app/internal/controllers"
	"app/internal/repositories/impl"
	"app/pkg/database"
	"net/http"

	"app/internal/router"

	"github.com/gin-gonic/gin"
)

func (s *Server) Routes() http.Handler {
	// TODO: Connect Database
	Db := database.ConnectDatabase()

	// TODO: DI
	authRepository := impl.NewAuthRepositoryImpl(Db)
	authController := controllers.NewAuthController(authRepository)

	// TODO:Router
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{

		router.AuthRoute(v1, authController)
	}

	return r
}
