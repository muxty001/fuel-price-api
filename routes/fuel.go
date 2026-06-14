package routes

import (
	"fuel-api/controllers"
	"fuel-api/middleware"

	"github.com/gin-gonic/gin"
)

func FuelPriceRoutes(r *gin.Engine) {
	r.POST(
		"/fuel-prices",
		middleware.AuthMiddleware(),
		controllers.CreateFuelPrice,
	)
}