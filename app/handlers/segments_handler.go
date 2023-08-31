package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func CreateNewSegment(c *fiber.Ctx) error {
	return c.SendString("Create new segment with name")
}

func DeleteSegment(c *fiber.Ctx) error {
	return c.SendString("Deletes segment by name")
}