package controllers

import "github.com/gofiber/fiber/v3"

func Welcome(c fiber.Ctx) error {
	return c.SendString("welcome to the api👋!")
}
