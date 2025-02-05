package employee_routes

import (
	"github.com/danielRamosMencia/edutech-api/internal/controllers/employee_controllers"
	"github.com/danielRamosMencia/edutech-api/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetEmployeeRoutes(router fiber.Router) {
	r := router.Group("/employees")
	r.Use(middlewares.AuthRequired)

	r.Get("/", employee_controllers.GetEmployees)
	r.Get("/:id", employee_controllers.GetEmployee)
	r.Post("/", employee_controllers.PostEmployee)
	r.Put("/:id", employee_controllers.PutEmployee)
	r.Patch("/:id", employee_controllers.PatchEmployee)
	r.Delete("/:id", employee_controllers.DeleteEmployee)
}
