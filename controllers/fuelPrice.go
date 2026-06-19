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

type FuelPriceDetails struct {
	StationName string `json:"station_name"`
	Location    string `json:"location"`
	Petrol      string `json:"petrol"`
	Diesel      string `json:"diesel"`
	Kerosene    string `json:"kerosene"`
}

func CreateFuelPrice(c *gin.Context) {
	var price FuelPrice

	if err := c.ShouldBindJSON(&price); err != nil {
		c.JSON(400, gin.H{
			"error": "Station ID, petrol, diesel and kerosene are required",
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

func GetFuelPrices(c *gin.Context) {
	rows, err := database.DB.Query(`
		SELECT id, station_id, petrol, diesel, kerosene
		FROM fuel_prices
	`)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to fetch fuel prices",
		})
		return
	}
	defer rows.Close()

	var prices []FuelPrice

	for rows.Next() {
		var price FuelPrice

		err := rows.Scan(
			&price.ID,
			&price.StationID,
			&price.Petrol,
			&price.Diesel,
			&price.Kerosene,
		)

		if err != nil {
			c.JSON(500, gin.H{
				"error": "Failed to read fuel prices",
			})
			return
		}

		prices = append(prices, price)
	}

	c.JSON(200, gin.H{
		"data": prices,
	})
}

func GetFuelPrice(c *gin.Context) {
	id := c.Param("id")

	var price FuelPrice

	query := `
	SELECT id, station_id, petrol, diesel, kerosene FROM fuel_prices WHERE id = $1
	`
	err := database.DB.QueryRow(
		query,
		id,
	).Scan(
		&price.ID,
		&price.StationID,
		&price.Petrol,
		&price.Diesel,
		&price.Kerosene,
	)

	if err != nil {
		c.JSON(404, gin.H{
			"error": "Fuel price not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": price,
	})
}

func GetFuelPricesWithStations(c *gin.Context) {
	query := `
		SELECT
			stations.name,
			stations.location,
			fuel_prices.petrol,
			fuel_prices.diesel,
			fuel_prices.kerosene
		FROM fuel_prices
		JOIN stations
			ON fuel_prices.station_id = stations.id
	`

	rows, err := database.DB.Query(query)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to fetch fuel prices",
		})
		return
	}
	defer rows.Close()

	var prices []FuelPriceDetails

	for rows.Next() {
		var price FuelPriceDetails

		err := rows.Scan(
			&price.StationName,
			&price.Location,
			&price.Petrol,
			&price.Diesel,
			&price.Kerosene,
		)

		if err != nil {
			c.JSON(500, gin.H{
				"error": "Failed to read fuel prices",
			})
			return
		}

		prices = append(prices, price)
	}

	c.JSON(200, gin.H{
		"data": prices,
	})
}

func DeleteFuelPrice(c *gin.Context) {
	id := c.Param("id")

	query := `
		DELETE FROM fuel_prices
		WHERE id = $1
	`

	_, err := database.DB.Exec(query, id)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to delete fuel price",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Fuel price deleted successfully",
	})
}