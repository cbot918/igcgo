package http

import (
	"github.com/cbot918/igcgo/internal/http/controller"
	"github.com/gofiber/fiber/v2"
)

func RegistRoute(rest *fiber.App) *fiber.App {
	common := controller.NewCommonController()

	auth := controller.NewAuthController()

	rest.Get("/ping", common.Ping)
	rest.Post("/auth", auth.Auth)

	return rest
}
