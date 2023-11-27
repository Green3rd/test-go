package controllers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type Person struct{ Email string }

type CatStatus struct {
	Verified  bool
	Feedback  string
	SentCount int
}
type CatResponse struct {
	Status    CatStatus
	_id       string
	__id      int
	User      string
	Text      string
	UpdatedAt string
	SendDate  string
	Deleted   bool
	Used      bool
	Source    string
	Type      string
}

func GetNews(c *fiber.Ctx) error {
	// Get Request Body
	// r := new(models.RequestModel)
	// if err := c.BodyParser(r); err != nil {
	// 	return err
	// }

	userId := c.Params("user")

	log.Info("Hello, World! " + userId)
	var user Person = Person{
		Email: "mikolaj73@gmail.com",
	}
	getSomething(c)
	return c.JSON(user)
}

// Get something
func getSomething(c *fiber.Ctx) (err error) {
	agent := fiber.Get("https://cat-fact.herokuapp.com/facts")
	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errs": errs,
		})
	}

	var result []CatResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err": err,
		})
	}
	log.Info("Got Cat result with length ", len(result), " with the first text ", result[0].Text)
	return c.Status(statusCode).JSON(result)
}
