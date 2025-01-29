package department_services

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/models/department_models"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func UpdateDepartment(ctx context.Context, departmentId string, input department_models.UpdateDeparment, author string) (int, string, error) {
	const query = `
	UPDATE "Department" SET
		"name" = $1,
		"code" = $2,
		"active" = $3,
		"country_id" = $4,
		"modified_by" = $5,
		"updated_at" = CURRENT_TIMESTAMP
	WHERE
		"id" = $6;
	`

	result, err := db.Connx.ExecContext(
		ctx,
		query,
		input.Name,
		input.Code,
		input.Active,
		input.CountryId,
		author,
		departmentId,
	)

	if err != nil {
		zap_logger.Logger.Info("Error updating department =>", zap.Error(err))
		return 500, ErrUpdate, err
	}

	row, err := result.RowsAffected()
	switch {
	case row == 0:
		return 404, NotFound, nil
	case err != nil:
		zap_logger.Logger.Info("Error updating department =>", zap.Error(err))
		return 500, ErrUpdate, err
	}

	return 200, SuccessUpdate, nil
}
