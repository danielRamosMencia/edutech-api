package auth_controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func SignOut(c *fiber.Ctx) error {
	expired := time.Now().Add(-time.Hour * 24)

	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   "",
		Expires: expired,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Sesión cerrada exitosamente",
	})
}
