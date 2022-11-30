package api

import (
	"github.com/gofiber/fiber/v2"
)

type JobHandler interface {
	Get(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	AddJob(c *fiber.Ctx) error
	UpdateJob(c *fiber.Ctx) error
	DeleteJob(c *fiber.Ctx) error
}
