package municipality_services

import (
	"context"
	"errors"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func DeleteMunicipality(ctx context.Context, municipalityId string) (int, string, error) {
	const query = `
	DELETE FROM "Municipality"
	WHERE
		"id" = $1;
	`

	result, err := db.Connx.ExecContext(ctx, query, municipalityId)
	if err != nil {
		zap_logger.Logger.Info("Error deleting municipality =>", zap.Error(err))
		return 500, ErrDelete, err
	}

	row, err := result.RowsAffected()
	switch {
	case row == 0:
		return 404, NotFound, errors.New("municipality not found")
	case err != nil:
		zap_logger.Logger.Info("Error deleting municipality =>", zap.Error(err))
		return 500, ErrDelete, err
	}

	return 200, SuccessDelete, nil
}
