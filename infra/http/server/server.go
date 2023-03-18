package server

import (
	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	router := fiber.New()

	// TODO: implement router

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("It is a stub!")
	})

	return router
}
