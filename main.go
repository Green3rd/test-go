package main

import (
	"os"
	"test/controllers"
	"test/initializer"

	"github.com/gofiber/fiber/v2"
)

func init() {
	initializer.LoadEnv()
}

func main() {
	app := fiber.New(fiber.Config{
		AppName: "BFF Cafe v0.0.1",
	})

	api := app.Group("/api")

	api.Get("/news/:user", controllers.GetNews)
	api.Get("/stocks/:user", controllers.GetNews)

	app.Listen(":" + os.Getenv("PORT"))
}
