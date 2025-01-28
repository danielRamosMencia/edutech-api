package auth_routes

import (
	"github.com/danielRamosMencia/edutech-api/internal/controllers/auth_controllers"
	"github.com/danielRamosMencia/edutech-api/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetAuthRoutes(router fiber.Router) {
	r := router.Group("/auth")

	r.Post("/sign-in", auth_controllers.SignIn)
	r.Post("/sign-out", middlewares.AuthRequired, auth_controllers.SignOut)
}
