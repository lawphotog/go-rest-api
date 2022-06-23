package main

import (
	"github.com/lawphotog/go-rest-api/packages/controllers"
	"github.com/lawphotog/go-rest-api/packages/repositories"
	"github.com/lawphotog/go-rest-api/packages/services"

	"github.com/gin-gonic/gin"
)

func registerRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "service is up",
		})
	})

	passengerRepository := repositories.NewPassengerRepository()
	passengerService := services.NewPassengerService(passengerRepository)
	passengerController := controllers.NewPassengerController(passengerService)

	r.GET("/passenger", passengerController.GetPassengers)

	return r
}
