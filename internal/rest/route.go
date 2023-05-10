package rest

import (
	"github.com/cbot918/health-helper/internal/rest/controller"
	"github.com/gofiber/fiber/v2"
)

func RegistRoute(server *fiber.App) *fiber.App {
	ctr := controller.New()

	server.Get("/", ctr.Hello)

	return server
}
