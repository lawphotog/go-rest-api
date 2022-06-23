package services

import (
	"log"

	"github.com/lawphotog/go-rest-api/packages/domains"
	"github.com/lawphotog/go-rest-api/packages/repositories"
)

type PassengerService struct {
	passengerRepository Repository
}

func NewPassengerService(passengerRepository Repository) *PassengerService {
	return &PassengerService{
		passengerRepository: passengerRepository,
	}
}

func (p *PassengerService) GetPassengers() ([]domains.Passenger, error) {
	passengers, err := p.passengerRepository.GetPassengers()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return mapPassengers(passengers), nil
}

func mapPassengers(passengers []repositories.Passenger) []domains.Passenger {
	domainPassengers := []domains.Passenger{}
	for _, v := range passengers {
		domainPassengers = append(domainPassengers, mapPassenger(v))
	}
	return domainPassengers
}

func mapPassenger(p repositories.Passenger) domains.Passenger {
	return domains.Passenger{
		Id: p.Id,
		Name: p.Name,
	}
}

type Repository interface {
	GetPassengers() ([]repositories.Passenger, error)
}
