package routes

import (
	"fuel-api/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)

	r.POST("/login", controllers.Login)
}