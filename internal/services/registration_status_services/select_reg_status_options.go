package registration_status_services

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/models"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func SelectRegStatusOptions(ctx context.Context) ([]models.CatalogOption, int, string, error) {
	var regStatus []models.CatalogOption

	const query = `
	SELECT 
		"id", 
		"name", 
		"code"
	FROM 
		"RegistrationStatus"
	WHERE
		"active" = true;
	`

	rows, err := db.Connx.QueryContext(ctx, query)
	if err != nil {
		zap_logger.Logger.Info("Error selecting registration statuses =>", zap.Error(err))
		return regStatus, 500, ErrSelectRegistrationStatuses, err
	}
	defer rows.Close()

	for rows.Next() {
		var regStat models.CatalogOption

		err := rows.Scan(
			&regStat.Id,
			&regStat.Name,
			&regStat.Code,
		)
		if err != nil {
			zap_logger.Logger.Info("Error scanning registration statuses =>", zap.Error(err))
			return regStatus, 500, ErrSelectRegistrationStatuses, err
		}

		regStatus = append(regStatus, regStat)
	}

	return regStatus, 200, Success, nil
}
