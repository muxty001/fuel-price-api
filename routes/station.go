package routes

import (
	"fuel-api/controllers"

	"github.com/gin-gonic/gin"
)

func StationRoutes(r *gin.Engine) {
	r.GET("/stations", controllers.GetStations)

	r.POST("/stations", controllers.CreateStation)

	r.PUT("/stations/:id", controllers.UpdateStation)
}