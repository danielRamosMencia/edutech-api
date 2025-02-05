package registration_status_services

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/models"
	"github.com/danielRamosMencia/edutech-api/internal/models/registration_status_models"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func SelectRegistrationStatuses(ctx context.Context, pagination models.PaginationParams) ([]registration_status_models.RegistrationStatus, int, string, error) {
	var registrationStatuses []registration_status_models.RegistrationStatus

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
	LIMIT $1 OFFSET $2;
	`

	rows, err := db.Connx.QueryContext(ctx, query, pagination.Limit, pagination.Offset)
	if err != nil {
		zap_logger.Logger.Info("Error selecting registration statuses =>", zap.Error(err))
		return registrationStatuses, 500, ErrSelectRegistrationStatuses, err
	}
	defer rows.Close()

	for rows.Next() {
		var registrationStatus registration_status_models.RegistrationStatus

		err := rows.Scan(
			&registrationStatus.Id,
			&registrationStatus.Name,
			&registrationStatus.Code,
			&registrationStatus.Active,
			&registrationStatus.CreatedBy,
			&registrationStatus.ModifiedBy,
			&registrationStatus.CreatedAt,
			&registrationStatus.UpdatedAt,
		)
		if err != nil {
			zap_logger.Logger.Info("Error scanning registration statuses =>", zap.Error(err))
			return registrationStatuses, 500, ErrSelectRegistrationStatuses, err
		}

		registrationStatuses = append(registrationStatuses, registrationStatus)
	}

	return registrationStatuses, 200, Success, nil
}
