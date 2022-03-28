package router

import (
	"DriverLocation/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Route(router *fiber.App) {
	router.Use(logger.New())
	router.Use(recover.New())

	users := router.Group("/users")
	users.Get("/test1", func(ctx *fiber.Ctx) error {
		return ctx.JSON(model.WebResponse{
			Code:   200,
			Status: "OK",
			Data:   "test1"})
	})
	users.Get("/test2", func(ctx *fiber.Ctx) error {
		return ctx.JSON(model.WebResponse{
			Code:   200,
			Status: "false",
			Data:   "test2"})
	})
	//users.Get("/", controller.GetUser)
	//users.Post("/", controller.CreateUser)
}
