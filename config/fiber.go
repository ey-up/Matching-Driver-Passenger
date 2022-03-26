package config

import (
	"DriverLocation/exception"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		AppName:      viper.GetString("port"),
		ErrorHandler: exception.ErrorHandler,
	}
}
