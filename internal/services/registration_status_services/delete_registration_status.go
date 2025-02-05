package registration_status_services

import (
	"context"
	"errors"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func DeleteRegistrationStatus(ctx context.Context, regStatusId string) (int, string, error) {
	const query = `
	DELETE FROM "RegistrationStatus"
	WHERE
		"id" = $1;
	`

	result, err := db.Connx.ExecContext(ctx, query, regStatusId)
	if err != nil {
		zap_logger.Logger.Info("Error deleting registrationStatus =>", zap.Error(err))
		return 500, ErrDelete, err
	}

	row, err := result.RowsAffected()

	switch {
	case row == 0:
		return 404, NotFound, errors.New("registrationStatus not found")
	case err != nil:
		zap_logger.Logger.Info("Error deleting registrationStatus =>", zap.Error(err))
		return 500, ErrDelete, err
	}

	return 200, SuccessDelete, nil
}
