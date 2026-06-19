package routes

import (
	"fuel-api/controllers"

	"github.com/gin-gonic/gin"

	"fuel-api/middleware"
)

func StationRoutes(r *gin.Engine) {
	r.GET("/stations", controllers.GetStations)
	r.GET("/stations/:id", controllers.GetStation)

	r.POST("/stations",
		middleware.AuthMiddleware(),
		controllers.CreateStation,
	)

	r.PUT("/stations/:id",
		middleware.AuthMiddleware(),
		controllers.UpdateStation,
	)

	r.DELETE("/stations/:id",
		middleware.AuthMiddleware(),
		controllers.DeleteStation,
	)
}