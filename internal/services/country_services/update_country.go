package country_services

import (
	"context"
	"errors"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/models/country_models"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func UpdateCountry(ctx context.Context, input country_models.UpdateCountry, countryId string, author string) (int, string, error) {
	const query = `
	UPDATE "Country" SET
		"name" = $1,
		"active" = $2,
		"A2" = $3,
		"A3" = $4,
		"code" = $5,
		"modified_by" = $6,
		"updated_at" = CURRENT_TIMESTAMP
	WHERE
		"id" = $7;
	`

	result, err := db.Connx.ExecContext(
		ctx,
		query,
		input.Name,
		input.Active,
		input.A2,
		input.A3,
		input.Code,
		author,
		countryId,
	)

	if err != nil {
		zap_logger.Logger.Info("Error updating country =>", zap.Error(err))
		return 500, ErrUpdate, err
	}

	row, err := result.RowsAffected()
	switch {
	case row == 0:
		return 404, NotFound, errors.New("country not found")
	case err != nil:
		zap_logger.Logger.Info("Error updating country =>", zap.Error(err))
		return 500, ErrUpdate, err
	}

	return 200, SuccessUpdate, nil
}
