package controllers

import (
	"fuel-api/database"

	"github.com/gin-gonic/gin"
)

type FuelPrice struct {
	ID        int     `json:"id"`
	StationID int     `json:"station_id" binding:"required"`
	Petrol    float64 `json:"petrol" binding:"required"`
	Diesel    float64 `json:"diesel" binding:"required"`
	Kerosene  float64 `json:"kerosene" binding:"required"`
}

func CreateFuelPrice(c *gin.Context) {
	var price FuelPrice

	if err := c.ShouldBindJSON(&price); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	query := `
		INSERT INTO fuel_prices (
			station_id,
			petrol,
			diesel,
			kerosene
		)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	err := database.DB.QueryRow(
		query,
		price.StationID,
		price.Petrol,
		price.Diesel,
		price.Kerosene,
	).Scan(&price.ID)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to save fuel price",
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "Fuel price added successfully",
		"data":    price,
	})
}