package controllers

import "github.com/gofiber/fiber/v3"

func UserController(c fiber.Ctx) error {
	return c.SendString("welcome to the api👋!")
}
