package employee_services

import (
	"context"
	"errors"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/models/employee_models"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func UpdateEmployee(ctx context.Context, employeeId string, input employee_models.UpdateEmployee, author string) (int, string, error) {
	const query = `
	UPDATE "Employee" SET
		"name" = $1,
		"middle_name" = $2,
		"last_name" = $3,
		"middle_last_name" = $4,
		"dni" = $5,
		"rtn" = $6,
		"address" = $7,
		"email" = $8,
		"phone" = $9,
		"birthdate" = $10,
		"active" = $11,
		"municipality_id" = $12,
		"institution_id" = $13,
		"modified_by" = $14,
		"updated_at" = CURRENT_TIMESTAMP
	WHERE
		"id" = $15;
	`

	result, err := db.Connx.ExecContext(
		ctx,
		query,
		input.Name,
		input.MiddleName,
		input.LastName,
		input.MiddleLastName,
		input.Dni,
		input.Rtn,
		input.Address,
		input.Email,
		input.Phone,
		input.Birthdate,
		input.Active,
		input.MunicipalityId,
		input.InstitutionId,
		author,
		employeeId,
	)

	if err != nil {
		zap_logger.Logger.Info("Error updating employee =>", zap.Error(err))
		return 500, ErrUpdate, err
	}

	row, err := result.RowsAffected()
	switch {
	case row == 0:
		return 404, NotFound, errors.New("employee not found")
	case err != nil:
		zap_logger.Logger.Info("Error updating employee =>", zap.Error(err))
		return 500, ErrUpdate, err
	}

	return 200, SuccessUpdate, nil
}
