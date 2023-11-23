package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type Person struct{ Email string }

func GetNews(c *fiber.Ctx) error {
	// r := new(models.RequestModel)
	// if err := c.BodyParser(r); err != nil {
	// 	return err
	// }
	userId := c.Params("user")

	log.Info("Hello, World! " + userId)
	var user Person = Person{
		Email: "mikolaj73@gmail.com",
	}
	return c.JSON(user)
}
