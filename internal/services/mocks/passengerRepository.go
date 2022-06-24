package mocks

import (
	"github.com/lawphotog/go-rest-api/internal/repositories"
	"github.com/stretchr/testify/mock"
)

type PassengerRepository struct {
	mock.Mock
}

func (p *PassengerRepository) GetPassengers() ([]repositories.Passenger, error) {
	args := p.Called()
	passengers := args.Get(0).([]repositories.Passenger)
	return passengers, args.Error(1)
}

func (p *PassengerRepository) AddPassenger(passenger repositories.Passenger) error {
	args := p.Called(passenger)
	return args.Error(0)
}
