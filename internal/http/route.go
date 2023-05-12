package http

import (
	"database/sql"

	"github.com/cbot918/igcgo/internal/http/controller"
	"github.com/gofiber/fiber/v2"
)

func RegistRoute(http *fiber.App, db *sql.DB) *fiber.App {
	common := controller.NewCommonController()

	auth := controller.NewAuthController(db)

	http.Get("/ping", common.Ping)
	http.Post("/auth", auth.Auth)

	return http
}
