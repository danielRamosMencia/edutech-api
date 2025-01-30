package grade_controllers

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/constants"
	"github.com/danielRamosMencia/edutech-api/internal/helpers"
	"github.com/danielRamosMencia/edutech-api/internal/models/grade_models"
	"github.com/danielRamosMencia/edutech-api/internal/services/grade_services"
	"github.com/danielRamosMencia/edutech-api/internal/validations"
	"github.com/gofiber/fiber/v2"
)

func PostGrade(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), constants.ContextTimeOut)
	defer cancel()

	sessionData, err := helpers.GetClaims(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Inautorizado",
			"code":  "grade-err-002",
		})
	}

	var newGrade grade_models.CreateGrade
	err = c.BodyParser(&newGrade)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Campos para solicitud de crear grado incorrectos",
			"code":  "grade-err-002",
		})
	}

	err = validations.Validate.Struct(newGrade)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": validations.MapValidatorErrors(err),
			"code":  "grade-err-002",
		})
	}

	status, message, err := grade_services.InsertGrade(ctx, newGrade, sessionData.Username)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"error": message,
			"code":  "grade-err-002",
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"message": message,
	})
}
