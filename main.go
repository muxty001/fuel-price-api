package main

import (
	"log"

	"fuel-api/database"
	"fuel-api/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDB()

	r := gin.Default()

	routes.AuthRoutes(r)
	routes.StationRoutes(r)
	routes.FuelPriceRoutes(r)

	r.Run(":3000")
}