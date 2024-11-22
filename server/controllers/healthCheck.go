package controllers

import "github.com/gofiber/fiber/v3"

func HealthCheck(c fiber.Ctx) error {
	return c.SendString("api is OK!")
}
