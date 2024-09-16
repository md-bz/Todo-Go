package main

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type header struct {
	Authorization string `reqHeader:"Authorization"`
}

type updateTodo struct {
	OldDescription string `json:"oldDescription"`
	NewDescription string `json:"newDescription"`
}

var user string

func main() {
	app := fiber.New()

	db := database()

	app.Use(func(c *fiber.Ctx) error {
		// placeholder for jwt
		h := new(header)
		err := c.ReqHeaderParser(h)

		auth := strings.Split(h.Authorization, " ")

		if err != nil || h.Authorization == "" || auth[0] != "Bearer" {
			return c.Status(403).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		c.Locals(user, strings.Trim(auth[1], " "))

		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		var todos []APITodo
		var user = c.Locals(user).(string)

		db.Model(&Todo{}).Where("user = ?", user).Find(&todos)

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

		return c.JSON(APITodo{t.Description, t.Done})
	})

	app.Delete("/", func(c *fiber.Ctx) error {
		t := new(Todo)
		err := c.BodyParser(t)

		if err != nil || t.Description == "" {
			return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
				"error": "please provide description in json",
			})
		}

		var user = c.Locals(user).(string)
		db.Where(&Todo{Description: t.Description, User: user}).Delete(&t)

		return c.JSON(APITodo{t.Description, t.Done})
	})

	app.Patch("/", func(c *fiber.Ctx) error {
		t := new(updateTodo)
		err := c.BodyParser(t)

		if err != nil || t.NewDescription == "" || t.OldDescription == "" {
			return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
				"error": "please provide OldDescription,NewDescription in json",
			})
		}

		var user = c.Locals(user).(string)

		res := new(Todo)
		db.Where(&Todo{Description: t.OldDescription, User: user}).Find(&res)

		res.Description = t.NewDescription
		db.Save(&res)
		return c.JSON(APITodo{res.Description, res.Done})
	})

	app.Post("/toggle", func(c *fiber.Ctx) error {
		t := new(Todo)
		err := c.BodyParser(t)

		if err != nil || t.Description == "" {
			return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
				"error": "please provide OldDescription,NewDescription in json",
			})
		}

		var user = c.Locals(user).(string)

		res := db.Where(&Todo{Description: t.Description, User: user}).First(&t)
		if res.Error != nil && res.Error.Error() == "record not found" {
			return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
				"error": "No todo found with that Description",
			})
		}

		t.Done = !t.Done
		db.Save(&t)
		return c.JSON(APITodo{t.Description, t.Done})
	})

	log.Fatal(app.Listen(":3000"))
}
