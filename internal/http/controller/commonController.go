package controller

import (
	"github.com/gofiber/fiber/v2"
)

type Common struct{}

func NewCommonController() *Common {
	return &Common{}
}

func (com *Common) Ping(c *fiber.Ctx) error {
	return c.SendString("Pong")
}
