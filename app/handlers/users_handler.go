package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func UpdateUserSegments(c *fiber.Ctx) error {
	return c.SendString("Update user segments")
}

func GetUserSegments(c *fiber.Ctx) error {
	return c.SendString("Returns user segments by user id")
}