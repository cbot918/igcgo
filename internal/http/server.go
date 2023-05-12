package http

import (
	"database/sql"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewHttpServer(db *sql.DB) *fiber.App {

	// Http Server Init
	server := fiber.New()

	// Logger and Recover
	server.Use(logger.New())
	server.Use(recover.New())

	// Route Register
	server = RegistRoute(server, db)

	// 404 Handler
	server.Use(func(c *fiber.Ctx) error {
		return c.Status(http.StatusNotFound).SendString("Not Found Handler")
	})

	return server
}
