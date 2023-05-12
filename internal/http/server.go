package http

import (
	"fmt"
	"net/http"
	"os"

	"github.com/cbot918/igcgo/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type HttpServer struct{}

func New() *HttpServer {
	return &HttpServer{}
}

func (r *HttpServer) Run() {
	// Config Setup
	cfg, err := config.New()
	if err != nil {
		fmt.Println("config.New error")
		os.Exit(1)
	}

	// Http Server Init
	app := fiber.New()

	// Logger and Recover
	app.Use(logger.New())
	app.Use(recover.New())

	// Serve Spa
	app.Static("/", "./igcgo-ui/build")

	// Route Register
	app = RegistRoute(app)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(http.StatusNotFound).SendString("Not Found Handler")
	})

	err = app.Listen(
		fmt.Sprintf(":%s", cfg.Server.Port),
		// fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	)
	if err != nil {
		panic(err)
	}
}
