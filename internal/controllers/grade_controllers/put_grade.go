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

func PutGrade(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), constants.ContextTimeOut)
	defer cancel()

	sessionData, err := helpers.GetClaims(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Inautorizado",
			"code":  "grade-err-003",
		})
	}

	var updateGrade grade_models.UpdateGrade
	err = c.BodyParser(&updateGrade)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Campos para solicitud de actualizar grado incorrectos",
			"code":  "grade-err-003",
		})
	}

	err = validations.Validate.Struct(updateGrade)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": validations.MapValidatorErrors(err),
			"code":  "grade-err-003",
		})
	}

	gradeId := c.Params("id")

	status, message, err := grade_services.UpdateGrade(ctx, updateGrade, gradeId, sessionData.Username)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"error": message,
			"code":  "grade-err-003",
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"message": message,
	})
}
