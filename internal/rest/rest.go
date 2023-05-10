package rest

import (
	"fmt"
	"net/http"
	"os"

	"github.com/cbot918/health-helper/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Rest struct{}

func New() *Rest {
	return &Rest{}
}

func (r *Rest) Run() {
	// Config Setup
	cfg, err := config.New()
	if err != nil {
		fmt.Println("config.New error")
		os.Exit(1)
	}

	// Rest Server Init
	app := fiber.New()

	// Logger and Recover
	app.Use(logger.New())
	app.Use(recover.New())

	// Serve Spa
	app.Static("/", "./igcgo-ui/build")

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(http.StatusNotFound).SendString("Not Found")
	})
	app = RegistRoute(app)

	err = app.Listen(
		fmt.Sprintf(":%s", cfg.Server.Port),
		// fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	)
	if err != nil {
		panic(err)
	}
}
