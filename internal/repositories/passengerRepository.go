package repositories

type PassengerRepository struct {
}

//mock data
var passengers []Passenger = []Passenger{
	{
		Id:   "1",
		Name: "John",
	},
	{
		Id:   "2",
		Name: "Joe",
	},
}

func NewPassengerRepository() *PassengerRepository {
	return &PassengerRepository{}
}

func (p *PassengerRepository) GetPassengers() ([]Passenger, error) {
	return passengers, nil
}

func (p *PassengerRepository) AddPassenger(passenger Passenger) error {
	passengers = append(passengers, passenger)
	return nil
}
