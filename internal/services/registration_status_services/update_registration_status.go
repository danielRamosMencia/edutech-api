package registration_status_services

import (
	"context"
	"errors"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/models/registration_status_models"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func UpdateRegistrationStatus(ctx context.Context, input registration_status_models.UpdateRegistrationStatus, regStatusId string, author string) (int, string, error) {
	const query = `
	UPDATE "RegistrationStatus" SET
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
		regStatusId,
	)
	if err != nil {
		zap_logger.Logger.Info("Error updating registrationStatus =>", zap.Error(err))
		return 500, ErrUpdate, err
	}

	row, err := result.RowsAffected()

	switch {
	case row == 0:
		return 404, NotFound, errors.New("registrationStatus not found")
	case err != nil:
		zap_logger.Logger.Info("Error updating registrationStatus =>", zap.Error(err))
		return 500, ErrUpdate, err
	}

	return 200, SuccessUpdate, nil
}
