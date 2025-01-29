package municipality_services

import (
	"context"
	"errors"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/models/municipality_models"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func UpdateMunicipality(ctx context.Context, municipalityId string, input municipality_models.UpdateMunicipality, author string) (int, string, error) {
	const query = `
	UPDATE "Municipality" 
	SET 
		"name" = $1, 
		"code" = $2,
		"active" = $3,
		"department_id" = $4,
		"modified_by" = $5,
		"updated_at" = CURRENT_TIMESTAMP
	WHERE "id" = $6;
	`

	result, err := db.Connx.ExecContext(
		ctx,
		query,
		input.Name,
		input.Code,
		input.Active,
		input.DepartmentId,
		author,
		municipalityId,
	)

	if err != nil {
		zap_logger.Logger.Info("Error updating municipality =>", zap.Error(err))
		return 500, ErrUpdate, err
	}

	row, err := result.RowsAffected()
	switch {
	case row == 0:
		return 404, NotFound, errors.New("municipality not found")
	case err != nil:
		zap_logger.Logger.Info("Error updating municipality =>", zap.Error(err))
		return 500, ErrUpdate, err
	}

	return 200, SuccessUpdate, nil
}
