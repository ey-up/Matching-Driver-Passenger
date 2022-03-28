package service

import (
	"DriverLocation/model"
	"DriverLocation/repository"
	"time"
)

type UserService interface {
	Create(request model.CreateUserRequest) (response model.CreateUserResponse)
}

type userService struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) UserService {
	return &userService{
		UserRepository: *userRepository,
	}
}

func (u userService) Create(request model.CreateUserRequest) (response model.CreateUserResponse) {

	user := model.User{
		Name:      request.Name,
		Email:     request.Email,
		CreatedAt: time.Now(),
	}
	u.UserRepository.Insert(user)
	response = model.CreateUserResponse{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
	return response
}
