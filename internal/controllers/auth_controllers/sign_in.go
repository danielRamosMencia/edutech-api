package auth_controllers

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/constants"
	"github.com/danielRamosMencia/edutech-api/internal/models/auth_models"
	"github.com/danielRamosMencia/edutech-api/internal/services/auth_services"
	"github.com/danielRamosMencia/edutech-api/internal/utils"
	"github.com/danielRamosMencia/edutech-api/internal/validations"
	"github.com/gofiber/fiber/v2"
)

func SignIn(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), constants.ContextTimeOut)
	defer cancel()

	var login auth_models.Login

	err := c.BodyParser(&login)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Campos para solicitud de inicio de sesión incorrectos",
			"code":  "auth-err-000",
		})
	}

	err = validations.Validate.Struct(login)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": validations.MapValidatorErrors(err),
			"code":  "auth-err-000",
		})
	}

	sessionData, status, message, err := auth_services.SelectSessionData(ctx, login)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"error": message,
			"code":  "auth-err-000",
		})
	}

	token, maxAge, err := utils.GenerateJWT(sessionData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Inautorizado",
			"code":  "auth-err-000",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HTTPOnly: true,
		MaxAge:   int(maxAge),
		SameSite: fiber.CookieSameSiteLaxMode,
	})

	return c.Status(status).JSON(fiber.Map{
		"token":       token,
		"sessionData": sessionData,
	})
}
