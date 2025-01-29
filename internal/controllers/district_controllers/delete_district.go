package district_controllers

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/constants"
	"github.com/danielRamosMencia/edutech-api/internal/services/district_services"
	"github.com/gofiber/fiber/v2"
)

func DeleteDistrict(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), constants.ContextTimeOut)
	defer cancel()

	districtId := c.Params("id")

	status, message, err := district_services.DeleteDistrict(ctx, districtId)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"error": message,
			"code":  "district-err-005",
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"message": message,
	})
}
