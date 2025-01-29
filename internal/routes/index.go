package routes

import (
	"github.com/danielRamosMencia/edutech-api/internal/routes/auth_routes"
	"github.com/danielRamosMencia/edutech-api/internal/routes/country_routes"
	"github.com/danielRamosMencia/edutech-api/internal/routes/department_routes"
	muncipality_routes "github.com/danielRamosMencia/edutech-api/internal/routes/municipality_routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func AppRoutes(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173/", // frontend url
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "Content-Type, Authorization",
		AllowCredentials: true,
	}))
	app.Use(logger.New(logger.Config{}))
	app.Use(requestid.New())
	app.Use(helmet.New())

	v1 := app.Group("api/v1", func(c *fiber.Ctx) error {
		c.Set("Version", "1")
		return c.Next()
	})

	// Health check
	v1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Server ready to Go!")
	})

	country_routes.SetCountryRoutes(v1)
	auth_routes.SetAuthRoutes(v1)
	department_routes.SetDepartmentRoutes(v1)
	muncipality_routes.SetMunicipalityRoutes(v1)
}
