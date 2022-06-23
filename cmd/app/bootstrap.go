package main

import (
	"github.com/lawphotog/go-rest-api/packages/controllers"
	"github.com/lawphotog/go-rest-api/packages/repositories"
	"github.com/lawphotog/go-rest-api/packages/services"
)

func getPassengerController() *controllers.PassengerController {
	
	passengerRepository := repositories.NewPassengerRepository()
	passengerService := services.NewPassengerService(passengerRepository)
	return controllers.NewPassengerController(passengerService)
}
