package registration_status_services

import (
	"context"
	"database/sql"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/models/registration_status_models"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func SelectRegistrationStatus(ctx context.Context, regStatusId string) (registration_status_models.RegistrationStatus, int, string, error) {
	var registrationStatus registration_status_models.RegistrationStatus

	const query = `
	SELECT
		"id",
		"name",
		"code",
		"active",
		"created_by",
		"modified_by",
		"created_at",
		"updated_at"
	FROM
		"RegistrationStatus"
	WHERE
		"id" = $1
	LIMIT 1;
	`

	row := db.Connx.QueryRowContext(ctx, query, regStatusId)
	err := row.Scan(
		&registrationStatus.Id,
		&registrationStatus.Name,
		&registrationStatus.Code,
		&registrationStatus.Active,
		&registrationStatus.CreatedBy,
		&registrationStatus.ModifiedBy,
		&registrationStatus.CreatedAt,
		&registrationStatus.UpdatedAt,
	)

	switch {
	case err == sql.ErrNoRows:
		return registrationStatus, 404, NotFound, err
	case err != nil:
		zap_logger.Logger.Info("Error selecting registration status =>", zap.Error(err))
		return registrationStatus, 500, ErrSelectRegistrationStatus, err
	}

	return registrationStatus, 200, Success, nil
}
