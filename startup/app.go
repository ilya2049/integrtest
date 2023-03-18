package startup

import (
	"integrtest/infra/http/server"

	"github.com/gofiber/fiber/v2"
)

func InitializeHTTPServer() *fiber.App {
	// TODO: initialize repository implementation

	// TODO: initialize app service

	// TODO: inject all in the server

	return server.New()
}
