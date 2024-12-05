package handler

import "github.com/gofiber/fiber/v2"

func CheckStatus(c *fiber.Ctx) error {
	return c.SendString("api is OK!")
}
