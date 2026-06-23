package controllers

import "github.com/gin-gonic/gin"

func Home(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Fuel Price API is running successfully",
		"author":  "Muxty",
		"version": "1.0.0",
	})
}