package controller

import (
	"DriverLocation/exception"
	"DriverLocation/model"
	"DriverLocation/service"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService *service.UserService) UserController {
	return UserController{UserService: *userService}
}

func (c UserController) Route(app *fiber.App) {
	app.Post("/api/user", c.Create)
}

func (c *UserController) Create(ctx *fiber.Ctx) error {
	var request model.CreateUserRequest
	err := ctx.BodyParser(&request)
	exception.IsPanic(err)
	response := c.UserService.Create(request)
	return ctx.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}
