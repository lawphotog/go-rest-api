package services_test

import (
	"errors"
	"testing"

	"github.com/lawphotog/go-rest-api/packages/repositories"
	"github.com/lawphotog/go-rest-api/packages/services"
	"github.com/lawphotog/go-rest-api/packages/services/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "passenger service")
}

var _ = Describe("when GetPassengers is called", func ()  {
	Describe("when is ok", func() {
		It("should return list of passengers", func() {
			mockRepo := &mocks.PassengerRepository{}
			expected := []repositories.Passenger{
				{
					Id: "1",
					Name: "Jim",
				},
				{
					Id: "2",
					Name: "Paul",
				},
			}
			mockRepo.On("GetPassengers").Return(expected, nil)
			service := services.NewPassengerService(mockRepo)
			actual, _ := service.GetPassengers()

			Expect(actual[0].Id).Should(Equal(expected[0].Id))
			Expect(actual[1].Id).Should(Equal(expected[1].Id))
		})

		It("should not return error value", func() {
			mockRepo := &mocks.PassengerRepository{}
			expected := []repositories.Passenger{
				{
					Id: "1",
					Name: "Jim",
				},
				{
					Id: "2",
					Name: "Paul",
				},
			}
			mockRepo.On("GetPassengers").Return(expected, nil)
			service := services.NewPassengerService(mockRepo)
			_, err := service.GetPassengers()

			Expect(err).Should(BeNil())
		})
	})

	Describe("when repository returns error", func() {
		It("should return list of passengers", func() {
			mockRepo := &mocks.PassengerRepository{}
			mockPassengers := []repositories.Passenger{
				{
					Id: "1",
					Name: "Jim",
				},
				{
					Id: "2",
					Name: "Paul",
				},
			}
			expected := errors.New("can't fetch data")
			mockRepo.On("GetPassengers").Return(mockPassengers, expected)
			service := services.NewPassengerService(mockRepo)
			_, err := service.GetPassengers()
			
			Expect(err.Error()).Should(Equal("can't fetch data"))
		})
	})
})
