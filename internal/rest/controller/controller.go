package controller

import (
	"github.com/gofiber/fiber/v2"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (ctr *Controller) Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
