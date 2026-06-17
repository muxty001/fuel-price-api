package routes

import (
	"fuel-api/controllers"
	"fuel-api/middleware"

	"github.com/gin-gonic/gin"
)

func FuelPriceRoutes(r *gin.Engine) {
	r.GET("/fuel-prices", controllers.GetFuelPrices)
	r.GET("/fuel-prices/:id", controllers.GetFuelPrice)
	r.GET("/fuel-prices-details", controllers.GetFuelPricesWithStations)

	r.POST(
		"/fuel-prices",
		middleware.AuthMiddleware(),
		controllers.CreateFuelPrice,
	)
}