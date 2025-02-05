package registration_status_routes

import (
	"github.com/danielRamosMencia/edutech-api/internal/controllers/registration_status_controllers"
	"github.com/danielRamosMencia/edutech-api/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetRegistrationStatusRoutes(router fiber.Router) {
	r := router.Group("/registration-status")
	r.Use(middlewares.AuthRequired)

	r.Get("/", registration_status_controllers.GetRegistrationStatuses)
	r.Get("/options", registration_status_controllers.GetRegStatusOptions)
	r.Get("/:id", registration_status_controllers.GetRegistrationStatus)
	r.Post("/", registration_status_controllers.PostRegistrationStatus)
	r.Put("/:id", registration_status_controllers.UpdateRegistrationStatus)
	r.Patch("/:id", registration_status_controllers.PatchRegistrationStatus)
	r.Delete("/:id", registration_status_controllers.DeleteRegistrationStatus)
}
