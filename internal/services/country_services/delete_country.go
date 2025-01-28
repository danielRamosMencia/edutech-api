package country_services

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func DeleteCountry(ctx context.Context, countryId string) (int, string, error) {
	const query = `
	DELETE FROM "Country"
	WHERE
		"id" = $1;
	`

	result, err := db.Connx.ExecContext(ctx, query, countryId)
	if err != nil {
		zap_logger.Logger.Info("Error deleting country =>", zap.Error(err))
		return 500, ErrDelete, err
	}

	row, err := result.RowsAffected()
	switch {
	case row == 0:
		return 404, NotFound, nil
	case err != nil:
		zap_logger.Logger.Info("Error deleting country =>", zap.Error(err))
		return 500, ErrDelete, err
	}

	return 200, SuccessDelete, nil
}
