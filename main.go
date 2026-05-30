package main

import (
	"fuel-api/database"
	"fuel-api/routes"

	"github.com/gin-gonic/gin"
	"log"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error .env file")
	}

	database.ConnectDB()
	
	r := gin.Default()

	routes.AuthRoutes(r)
	routes.StationRoutes(r)

	r.Run(":3000")
}
