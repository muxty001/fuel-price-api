package controllers

import "github.com/gin-gonic/gin"



	type Station struct {
		Name string `json:"name"`
		Location string `json:"location"`
	}

	var stations []Station

	func GetStations(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": stations,
		})
	}

	func CreateStation(c *gin.Context){
		var station Station

		c.ShouldBindJSON(&station)

		stations = append(stations, station)

		c.JSON(200, gin.H{
			"message": "Station created successfully",
			"data": station,
		})
	}
