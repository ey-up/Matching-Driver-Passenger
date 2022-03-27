package main

import (
	"DriverLocation/config"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("%s", err.Error())
	}
	app := fiber.New(config.NewFiberConfig())
	app.Use(logger.New())
	app.Use(recover.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return fiber.NewError(782, "Custom error message")
	})
	fmt.Println("App started port: " + viper.GetString("port"))
	log.Fatal(app.Listen(":" + viper.GetString("port")))
}
