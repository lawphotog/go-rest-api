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

func (p *PassengerService) AddPassenger(passenger domains.Passenger) error {
	err := p.passengerRepository.AddPassenger(repositories.Passenger(passenger))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func mapPassengers(passengers []repositories.Passenger) []domains.Passenger {
	domainPassengers := []domains.Passenger{}
	for _, v := range passengers {
		domainPassengers = append(domainPassengers, domains.Passenger(v))
	}
	return domainPassengers
}

type Repository interface {
	GetPassengers() ([]repositories.Passenger, error)
	AddPassenger(passenger repositories.Passenger) error
}
