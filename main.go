package main

import (
	"fuel-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.AuthRoutes(r)
	routes.StationRoutes(r)

	r.Run(":3000")
}
