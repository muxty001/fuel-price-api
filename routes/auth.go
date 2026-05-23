package routes

import "github.com/gin-gonic/gin"

func AuthRoutes(r *gin.Engine) {
	r.GET("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Login route working",
		})
	})
}
