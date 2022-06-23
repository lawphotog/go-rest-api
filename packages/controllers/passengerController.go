package controllers

import (
	"github.com/lawphotog/go-rest-api/packages/domains"

	"github.com/gin-gonic/gin"
)

type PassengerController struct {
	passengerService Service
}

func NewPassengerController(passengerService Service) *PassengerController {
	return &PassengerController{
		passengerService: passengerService,
	}
}

func (p *PassengerController) GetPassengers(c *gin.Context) {
	passengers := p.passengerService.GetPassengers()
	c.JSON(200, gin.H{
		"passengers": passengers,
	})
}

type Service interface {
	GetPassengers() []domains.Passenger
}
