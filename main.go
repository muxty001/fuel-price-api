package main

import (
	"fuel-api/database"
	"fuel-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()
	
	r := gin.Default()

	routes.AuthRoutes(r)
	routes.StationRoutes(r)

	r.Run(":3000")
}
