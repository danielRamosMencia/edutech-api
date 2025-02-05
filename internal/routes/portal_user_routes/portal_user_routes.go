package portal_user_routes

import (
	"github.com/danielRamosMencia/edutech-api/internal/controllers/portal_user_controllers"
	"github.com/danielRamosMencia/edutech-api/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetPortalUserRoutes(router fiber.Router) {
	r := router.Group("/portal-users")
	r.Use(middlewares.AuthRequired)

	r.Get("/", portal_user_controllers.GetPortalUsers)
	// r.Get("/:id", portal_user_controllers.GetPortalUser)
	// r.Post("/", portal_user_controllers.PostPortalUser)
	// r.Put("/:id", portal_user_controllers.PutPortalUser)
	// r.Patch("/:id", portal_user_controllers.PatchPortalUser)
	// r.Delete("/:id", portal_user_controllers.DeletePortalUser)
}
