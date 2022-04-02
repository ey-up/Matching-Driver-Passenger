package service

import (
	"DriverLocation/model"
	"DriverLocation/repository"
	"strconv"
	"time"
)

type DriverService interface {
	Create(request model.CreateDriverRequest) (response model.CreateDriverResponse)
}

type driverService struct {
	driverRepository repository.DriverRepository
}

func (d driverService) Create(request model.CreateDriverRequest) (response model.CreateDriverResponse) {
	// todo: validation
	driver := model.Driver{
		Latitude:    request.Latitude,
		Longtitude:  request.Longtitude,
		HasCustomer: false,
		IsActive:    true,
		CreatedDate: strconv.FormatInt(time.Now().UnixMilli(), 10),
	}
	id := d.driverRepository.Create(driver)
	response = model.CreateDriverResponse{
		Id:          id,
		Latitude:    driver.Latitude,
		Longtitude:  driver.Longtitude,
		HasCustomer: driver.HasCustomer,
		IsActive:    driver.IsActive,
		CreatedDate: driver.CreatedDate,
	}
	return response
}

func NewDriverService(driverRepository *repository.DriverRepository) DriverService {
	return driverService{driverRepository: *driverRepository}
}
