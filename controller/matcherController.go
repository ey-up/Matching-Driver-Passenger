package controller

import (
	"DriverLocation/exception"
	"DriverLocation/model"
	"DriverLocation/service"
	"github.com/gofiber/fiber/v2"
)

type MatcherController struct {
	MatcherService service.MatcherService
}

func NewMatcherController(matcherService *service.MatcherService) MatcherController {
	return MatcherController{MatcherService: *matcherService}
}

func (matcherController MatcherController) Route(app *fiber.App) {
	matcher := app.Group("/api/matcher")
	matcher.Put("/", matcherController.Match)
	matcher.Put("/findDriver", matcherController.FindTheNearestDriver)
}

// Match ShowAccount godoc
// @Summary Match a passenger with a driver
// @Tags Matcher
// @Description Matcher the passenger with the driver
// @Accept  json
// @Param        passenger  body      model.MatchRequest  true  "PassengerId DriverId"
// @Produce  json
// @Success 200 {object} model.WebResponse
// @Failure 400 {object} model.WebResponse
// @Router /api/matcher [Put]
func (matcherController MatcherController) Match(ctx *fiber.Ctx) error {
	var request model.MatchRequest
	err := ctx.BodyParser(&request)
	exception.IsPanic(err)
	//todo: request null check
	response := matcherController.MatcherService.Match(request.PassengerId, request.DriverId)
	return ctx.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

// FindTheNearestDriver ShowAccount godoc
// @Summary Find driver
// @Tags Matcher
// @Description Find the nearest driver to passenger
// @Accept  json
// @Param        passenger  body      model.FindDriverRequest  true  "PassengerId DriverId"
// @Produce  json
// @Success 200 {object} model.WebResponse
// @Failure 400 {object} model.WebResponse
// @Router /api/matcher/findDriver [Put]
func (matcherController MatcherController) FindTheNearestDriver(ctx *fiber.Ctx) error {
	var request model.FindDriverRequest
	err := ctx.BodyParser(&request)
	exception.IsPanic(err)

	nearestDriver := matcherController.MatcherService.GetAvailableDrivers(request)
	return ctx.Status(200).JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   nearestDriver,
	})
}
