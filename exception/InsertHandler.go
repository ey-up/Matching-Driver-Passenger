package exception

import (
	"DriverLocation/model"
	"github.com/gofiber/fiber/v2"
)

func InsertHandler(ctx *fiber.Ctx, err error) error {
	if err != nil {
		return ctx.JSON(model.WebResponse{
			Code:   400,
			Status: "insert failed",
			Data:   err.Error(),
		})
	}
	return nil
}
