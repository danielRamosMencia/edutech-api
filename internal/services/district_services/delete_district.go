package district_services

import (
	"context"
	"errors"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func DeleteDistrict(ctx context.Context, districtId string) (int, string, error) {
	const query = `
	DELETE FROM "District"
	WHERE
		"id" = $1;
	`

	result, err := db.Connx.ExecContext(ctx, query, districtId)
	if err != nil {
		zap_logger.Logger.Info("Error deleting district =>", zap.Error(err))
		return 500, ErrDelete, err
	}

	row, err := result.RowsAffected()
	switch {
	case row == 0:
		return 404, NotFound, errors.New("district not found")
	case err != nil:
		zap_logger.Logger.Info("Error deleting district =>", zap.Error(err))
		return 500, ErrDelete, err
	}

	return 200, SuccessDelete, nil
}
