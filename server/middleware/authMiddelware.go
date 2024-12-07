package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jacobshade/lbuc-admin/server/database"
)

func AuthRequired() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Next()

		sess, err := database.Store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		email := sess.Get("email")
		if email == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		return c.Next()
	}
}
