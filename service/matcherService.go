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
}

type matcherService struct {
	matcherRepository repository.MatcherRepository
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
