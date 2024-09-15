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

	app.Use(func(c *fiber.Ctx) error {
		h := new(header)
		err := c.ReqHeaderParser(h)
		if err != nil || h.Authorization == "" {
			return c.Status(403).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		c.Locals(user, strings.Split(h.Authorization, " ")[1])
		return c.Next()

	})

	app.Get("/", func(c *fiber.Ctx) error {
		var user = c.Locals(user).(string)
		println(user)
		var res = database[user]
		fmt.Println(res)

		return c.JSON(res)
	})

	app.Post("/", func(c *fiber.Ctx) error {
		t := new(todo)
		err := c.BodyParser(t)

		if err != nil || t.Description == "" {

			return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
				"error": "please provide description in json",
			})
		}

		var user = c.Locals(user).(string)

		database[user] = append(database[user], *t)
		return c.JSON(*t)
	})

	log.Fatal(app.Listen(":3000"))
}
