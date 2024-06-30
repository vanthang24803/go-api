package server

import (
	"app/internal/controllers"
	"app/internal/repositories/impl"
	"app/pkg/database"
	"app/pkg/middlewares"
	"net/http"

	"app/internal/router"

	"github.com/gin-gonic/gin"
)

func (s *Server) Application() http.Handler {
	// TODO: Connect Database
	Db := database.ConnectDatabase()

	// TODO: DI
	authRepository := impl.NewAuthRepositoryImpl(Db)
	authController := controllers.NewAuthController(authRepository)

	// TODO: Main
	app := gin.Default()

	// TODO: Middlewares

	app.Use(middlewares.Logger())

	// TODO: Router
	v1 := app.Group("/api/v1")
	{

		router.AuthRoute(v1, authController)
	}

	return app
}
