package service

import (
	"DriverLocation/db"
	"DriverLocation/exception"
	"DriverLocation/model"
	"DriverLocation/repository"
	"fmt"
	"github.com/gocarina/gocsv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math"
	"os"
	"reflect"
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
			Latitude:     Coordinate.Latitude,
			Longtitude:   Coordinate.Longtitude,
			HasPassenger: false,
			IsActive:     true,
			CreatedDate:  strconv.FormatInt(time.Now().UnixMilli(), 10),
		}

		Drivers = append(Drivers, driver)
	}
	d.driverRepository.InsertFromCsv(Drivers)

	_, err = CoordinatesFile.Seek(0, 0)
	exception.IsPanic(err)
}

func (d driverService) Create(request model.CreateDriverRequest) (response model.CreateDriverResponse) {
	driver := model.Driver{
		Latitude:     request.Latitude,
		Longtitude:   request.Longtitude,
		HasPassenger: false,
		IsActive:     true,
		CreatedDate:  strconv.FormatInt(time.Now().UnixMilli(), 10),
	}
	id := d.driverRepository.Create(driver)
	response = model.CreateDriverResponse{
		Id:           id,
		Latitude:     driver.Latitude,
		Longtitude:   driver.Longtitude,
		HasPassenger: driver.HasPassenger,
		IsActive:     driver.IsActive,
		CreatedDate:  driver.CreatedDate,
	}
	return response
}

func GetAvailableDrivers(passengerInfo model.FindDriverRequest) model.TheNearestDriverResponse {
	driverInfo := model.Driver{
		IsActive:     true,
		HasPassenger: false,
		Longtitude:   passengerInfo.Longtitude,
		Latitude:     passengerInfo.Latitude,
	}
	database := db.Connection()
	drivers := repository.NewDriverRepository(database).FindAvailableDrivers(driverInfo)
	fmt.Println("cap of drivers: ", cap(drivers))
	var driver, distance = getTheNearestDriver(drivers, passengerInfo)
	theNearestDriver := model.TheNearestDriverResponse{
		Id:           driver.Id,
		Latitude:     driver.Latitude,
		Longtitude:   driver.Longtitude,
		HasPassenger: driver.HasPassenger,
		IsActive:     driver.IsActive,
		CreatedDate:  driver.CreatedDate,
		Distance:     distance,
	}
	return theNearestDriver
}

func getTheNearestDriver(drivers []model.Driver, passengerInfo model.FindDriverRequest) (model.Driver, float64) {
	var min = math.MaxFloat64
	var index int64 = -1
	var distance float64 = 0
	for i, driver := range drivers {
		fmt.Println(reflect.TypeOf(i))

		fmt.Println(i, " ", driver)
		lat1 := passengerInfo.Latitude
		lon1 := passengerInfo.Longtitude

		lat2 := driver.Latitude
		lon2 := driver.Longtitude

		dLat := toRadians(lat2 - lat1)
		dLon := toRadians(lon2 - lon1)

		lat1 = toRadians(lat1)
		lat2 = toRadians(lat2)

		var a = math.Pow(math.Sin(dLat/2), 2) +
			math.Pow(math.Sin(dLon/2), 2)*
				math.Cos(lat1)*
				math.Cos(lat2)
		distance = 2 * math.Asin(math.Sqrt(a)) * 6371.009
		
		fmt.Println("total: ", distance)
		if distance < min {
			min = distance
			index = int64(i)
		}
	}
	if index == -1 {
		return model.Driver{}, 0
	}
	fmt.Println("min: ", min)
	return drivers[index], min
}

func toRadians(val float64) float64 {
	return val * math.Pi / 180
}

func NewDriverService(driverRepository *repository.DriverRepository) DriverService {
	return driverService{driverRepository: *driverRepository}
}

func UpdateHasPassengerByDriverId(id primitive.ObjectID) {
	database := db.Connection()
	repository.NewDriverRepository(database).UpdateHasPassengerByDriverId(id)
}
