package service

import (
	"DriverLocation/model"
	"DriverLocation/repository"
	"strconv"
	"time"
)

type PassengerService interface {
	Create(request model.CreatePassengerRequest) (response model.CreatePassengerResponse)
}

type passengerService struct {
	passengerRepository repository.PassengerRepository
}

func NewPassengerService(passengerRepository *repository.PassengerRepository) PassengerService {
	return passengerService{passengerRepository: *passengerRepository}
}

func (p passengerService) Create(request model.CreatePassengerRequest) (response model.CreatePassengerResponse) {
	passenger := model.Passenger{
		Latitude:    request.Latitude,
		Longtitude:  request.Longtitude,
		HasDriver:   false,
		IsActive:    true,
		CreatedDate: strconv.FormatInt(time.Now().UnixMilli(), 10),
	}
	id := p.passengerRepository.Create(passenger)
	response = model.CreatePassengerResponse{
		Id:          id,
		Latitude:    passenger.Latitude,
		Longtitude:  passenger.Longtitude,
		HasDriver:   passenger.HasDriver,
		IsActive:    passenger.IsActive,
		CreatedDate: passenger.CreatedDate,
	}
	return response
}
