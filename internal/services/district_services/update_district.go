package district_services

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/models/district_models"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func UpdateDistrict(ctx context.Context, input district_models.UpdateDistrict, districtId string, author string) (int, string, error) {
	const query = `
	UPDATE "District" SET
		"name" = $1,
		"code" = $2,
		"active" = $3,
		"modified_by" = $4,
		"updated_at" = CURRENT_TIMESTAMP
	WHERE
		"id" = $5;
	`

	result, err := db.Connx.ExecContext(
		ctx,
		query,
		input.Name,
		input.Code,
		input.Active,
		author,
		districtId,
	)

	if err != nil {
		zap_logger.Logger.Info("Error updating district =>", zap.Error(err))
		return 500, ErrUpdate, err
	}

	row, err := result.RowsAffected()
	switch {
	case row == 0:
		return 404, NotFound, nil
	case err != nil:
		zap_logger.Logger.Info("Error updating district =>", zap.Error(err))
		return 500, ErrUpdate, err
	}

	return 200, SuccessUpdate, nil
}
