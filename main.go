package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type header struct {
	Authorization string `reqHeader:"Authorization"`
}

var user string

func main() {
	app := fiber.New()

	db := database()

	app.Use(func(c *fiber.Ctx) error {
		// placeholder for jwt
		h := new(header)
		err := c.ReqHeaderParser(h)

		if err != nil || h.Authorization == "" {
			return c.Status(403).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		c.Locals(user, strings.Trim(strings.Split(h.Authorization, " ")[1], " "))

		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		var todos []APITodo
		var user = c.Locals(user).(string)

		db.Model(&Todo{}).Where("user = ?", user).Find(&todos)
		fmt.Println(todos)

		return c.JSON(todos)
	})

	app.Post("/", func(c *fiber.Ctx) error {
		t := new(Todo)
		err := c.BodyParser(t)

		if err != nil || t.Description == "" {
			return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
				"error": "please provide description in json",
			})
		}

		var user = c.Locals(user).(string)
		db.Create(&Todo{User: user, Description: t.Description})

		return c.JSON(fiber.Map{"status": "ok"})
	})

	log.Fatal(app.Listen(":3000"))
}
