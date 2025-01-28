package validations

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func ActiveRequired(c *fiber.Ctx) (bool, string, error) {
	var reqBody map[string]interface{}

	err := c.BodyParser(&reqBody)
	if err != nil {
		return false, "Solicitud para cambiar campo activo incorrecta", err
	}

	hasActive, ok := reqBody["active"]
	if !ok {
		return false, "No se encontró el campo 'active'", errors.New("request body missing active field")
	}

	active, ok := hasActive.(bool)
	if !ok {
		return false, "El campo activo no es un valor booleano", errors.ErrUnsupported
	}

	return active, "", nil
}
