package main

import (
	"github.com/lawphotog/go-rest-api/internal/controllers"
	"github.com/lawphotog/go-rest-api/internal/repositories"
	"github.com/lawphotog/go-rest-api/internal/services"
)

func getPassengerController() *controllers.PassengerController {
	
	passengerRepository := repositories.NewPassengerRepository()
	passengerService := services.NewPassengerService(passengerRepository)
	return controllers.NewPassengerController(passengerService)
}
