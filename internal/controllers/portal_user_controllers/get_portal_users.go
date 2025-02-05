package portal_user_controllers

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/constants"
	"github.com/danielRamosMencia/edutech-api/internal/helpers"
	"github.com/danielRamosMencia/edutech-api/internal/services/portal_user_services"
	"github.com/gofiber/fiber/v2"
)

func GetPortalUsers(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), constants.ContextTimeOut)
	defer cancel()

	pagination := helpers.MapPagination(c)

	portalUsers, status, message, err := portal_user_services.SelectPortalUsers(ctx, pagination)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"error": message,
			"code":  "portal-user-err-000",
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"message": message,
		"records": portalUsers,
	})

}
