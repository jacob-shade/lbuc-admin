package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jacobshade/lbuc-admin/server/interactors"
	"github.com/jacobshade/lbuc-admin/server/model"
)

func GetChecksForPlayer(c *fiber.Ctx) error {
	// Get player id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	checks, err := interactors.GetChecksForPlayer(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Checks not found"})
	}

	return c.Status(fiber.StatusOK).JSON(checks)
}

func GetChecksForTask(c *fiber.Ctx) error {
	// Get task id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	checks, err := interactors.GetChecksForTask(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Checks not found"})
	}

	return c.Status(fiber.StatusOK).JSON(checks)
}

func UpdateCheck(c *fiber.Ctx) error {
	var check model.Check
	if err := c.BodyParser(&check); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := interactors.UpdateCheck(check); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Check not found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Check updated"})
}
