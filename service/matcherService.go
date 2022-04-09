package service

import (
	"DriverLocation/model"
	"DriverLocation/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
	"time"
)

type MatcherService interface {
	Match(primitive.ObjectID, primitive.ObjectID) primitive.ObjectID
	GetAvailableDrivers(request model.FindDriverRequest) model.TheNearestDriverResponse
}

type matcherService struct {
	matcherRepository repository.MatcherRepository
}

func (m matcherService) GetAvailableDrivers(passengerInfo model.FindDriverRequest) model.TheNearestDriverResponse {
	driver := GetAvailableDrivers(passengerInfo)
	return driver
}

func (m matcherService) Match(passengerId primitive.ObjectID, driverId primitive.ObjectID) primitive.ObjectID {
	UpdateHasPassengerByDriverId(driverId)
	UpdateHasDriverByPassengerId(passengerId)
	//todo:  rollback
	matcher := model.Matcher{
		DriverId:    driverId,
		PassengerId: passengerId,
		IsActive:    true,
		CreatedDate: strconv.FormatInt(time.Now().UnixMilli(), 10),
	}
	id := m.matcherRepository.Match(matcher)
	return id
}

func NewMatcherService(matcherRepository *repository.MatcherRepository) MatcherService {
	return matcherService{matcherRepository: *matcherRepository}
}
