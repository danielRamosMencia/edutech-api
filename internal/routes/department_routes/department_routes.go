package department_routes

import (
	"github.com/danielRamosMencia/edutech-api/internal/controllers/department_controllers"
	"github.com/danielRamosMencia/edutech-api/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetDepartmentRoutes(router fiber.Router) {
	r := router.Group("/departments")
	r.Use(middlewares.AuthRequired)

	r.Get("/", department_controllers.GetDepartments)
	r.Get("/options", department_controllers.GetDepartmentOptions)
	r.Get("/:id", department_controllers.GetDepartment)
	// r.Post("/", department_controllers.PostDepartment)
	// r.Put("/:id", department_controllers.PutDepartment)
	// r.Patch("/:id", department_controllers.PatchDepartment)
	// r.Delete("/:id", department_controllers.DeleteDepartment)
}
