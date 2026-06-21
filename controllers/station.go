package controllers

import (
	"fuel-api/database"

	"github.com/gin-gonic/gin"
)

type Station struct {
	ID       int    `json:"id"`
	Name     string `json:"name" binding:"required"`
	Location string `json:"location" binding:"required"`
}

func GetStations(c *gin.Context) {
	rows, err := database.DB.Query(
		"SELECT id, name, location FROM stations",
	)

	if err != nil {
    c.JSON(500, gin.H{
        "error": err.Error(),
    })
    return
}
	defer rows.Close()

	var stations []Station

	for rows.Next() {
		var station Station

		rows.Scan(
			&station.ID,
			&station.Name,
			&station.Location,
		)

		stations = append(stations, station)
	}

	c.JSON(200, gin.H{
		"data": stations,
	})
}

func GetStation(c *gin.Context) {
	id := c.Param("id")

	var station Station

	query := `
		SELECT id, name, location
		FROM stations
		WHERE id = $1
	`

	err := database.DB.QueryRow(
		query,
		id,
	).Scan(
		&station.ID,
		&station.Name,
		&station.Location,
	)

	if err != nil {
		c.JSON(404, gin.H{
			"error": "Station not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": station,
	})
}

func CreateStation(c *gin.Context) {
	var station Station

	if err := c.ShouldBindJSON(&station); err != nil {
	c.JSON(400, gin.H{
		"error": "Name and location are required",
	})
	return
}

	query := `
		INSERT INTO stations (name, location)
		VALUES ($1, $2)
		RETURNING id
	`

	err := database.DB.QueryRow(
		query,
		station.Name,
		station.Location,
	).Scan(&station.ID)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Station created successfully",
		"data":    station,
	})
}

func UpdateStation(c *gin.Context) {
	id := c.Param("id")

	var station Station

	if err := c.ShouldBindJSON(&station); err != nil {
	c.JSON(400, gin.H{
		"error": "Name and location are required",
	})
	return
}

	query := `
		UPDATE stations
		SET name=$1, location=$2
		WHERE id=$3
	`

	_, err := database.DB.Exec(
		query,
		station.Name,
		station.Location,
		id,
	)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Station updated successfully",
	})
}

func DeleteStation(c *gin.Context) {
	id := c.Param("id")

	query := `DELETE FROM stations WHERE id=$1`

	_, err := database.DB.Exec(query, id)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Station deleted successfully",
	})
}


