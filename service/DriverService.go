package service

import (
	"DriverLocation/exception"
	"DriverLocation/model"
	"DriverLocation/repository"
	"github.com/gocarina/gocsv"
	"os"
	"strconv"
	"time"
)

type DriverService interface {
	Create(request model.CreateDriverRequest) (response model.CreateDriverResponse)
	InsertFromCsv()
}

type driverService struct {
	driverRepository repository.DriverRepository
}

func (d driverService) InsertFromCsv() {

	CoordinatesFile, err := os.OpenFile("Coordinates.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	exception.IsPanic(err)

	defer func(CoordinatesFile *os.File) {
		err := CoordinatesFile.Close()
		panic(err)
	}(CoordinatesFile)

	var Coordinates []*model.CreateDriverRequest
	err = gocsv.UnmarshalFile(CoordinatesFile, &Coordinates)
	exception.IsPanic(err)
	var Drivers []interface{}

	for _, Coordinate := range Coordinates {
		driver := model.Driver{
			Latitude:    Coordinate.Latitude,
			Longtitude:  Coordinate.Longtitude,
			HasCustomer: false,
			IsActive:    true,
			CreatedDate: strconv.FormatInt(time.Now().UnixMilli(), 10),
		}

		Drivers = append(Drivers, driver)
	}
	d.driverRepository.InsertFromCsv(Drivers)

	_, err = CoordinatesFile.Seek(0, 0)
	exception.IsPanic(err)
}

func (d driverService) Create(request model.CreateDriverRequest) (response model.CreateDriverResponse) {
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
