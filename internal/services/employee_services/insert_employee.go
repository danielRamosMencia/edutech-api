package employee_services

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/models/employee_models"
	"github.com/danielRamosMencia/edutech-api/internal/utils"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func InsertEmployee(ctx context.Context, input employee_models.CreateEmployee, author string) (int, string, error) {
	const query = `
	INSERT INTO "Employee" (
		"id",
		"name",
		"middle_name",
		"last_name",
		"middle_last_name",
		"dni",
		"rtn",
		"address",
		"email",
		"phone",
		"birthdate",
		"active",
		"municipality_id",
		"institution_id",
		"created_by",
		"created_at",
		"updated_at"
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7,
		$8,
		$9,
		$10,
		$11,
		$12,
		$13,
		$14,
		$15,
		CURRENT_TIMESTAMP,
		CURRENT_TIMESTAMP
	);
	`

	id := utils.GenerateId()

	_, err := db.Connx.ExecContext(
		ctx,
		query,
		id,
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
	)

	if err != nil {
		zap_logger.Logger.Info("Error inserting employee =>", zap.Error(err))
		return 500, ErrInsert, err
	}

	return 201, SuccessInsert, nil
}
