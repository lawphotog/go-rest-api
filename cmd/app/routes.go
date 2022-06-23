package main

import (
	"github.com/gin-gonic/gin"
)

func registerRoutes(r *gin.Engine) *gin.Engine {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "service is up",
		})
	})

	passengerController := getPassengerController()
	r.GET("/passenger", passengerController.GetPassengers)
	r.POST("/passenger/add", passengerController.AddPassenger)

	return r
}
