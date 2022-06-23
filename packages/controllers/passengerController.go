package controllers

import (
	"log"

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
	passengers, err := p.passengerService.GetPassengers()
	if err != nil {
		c.JSON(200, gin.H{
			"error": "An error has occurred",
		})
		return
	}
	c.JSON(200, gin.H{
		"passengers": passengers,
	})
}

func (p *PassengerController) AddPassenger(c *gin.Context) {
	var passenger domains.Passenger
	if err := c.BindJSON(&passenger); err != nil {
		log.Println(err)
		c.JSON(200, gin.H{
			"error": "An error has occurred",
		})
		return
	}
	err := p.passengerService.AddPassenger(passenger)
	if err != nil {
		log.Println(err)
		c.JSON(200, gin.H{
			"error": "An error has occurred",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "a new passenger is added successfully",
	})
}

type Service interface {
	GetPassengers() ([]domains.Passenger, error)
	AddPassenger(passenger domains.Passenger) error
}
