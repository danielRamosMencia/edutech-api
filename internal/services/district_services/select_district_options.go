package district_services

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/models"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func SelectDistrictOptions(ctx context.Context) ([]models.CatalogOption, int, string, error) {
	var districts []models.CatalogOption

	const query = `
	SELECT
		"id",
		"name",
		"code"
	FROM
		"District"
	WHERE 
		"active" = true;
	`

	rows, err := db.Connx.QueryContext(ctx, query)
	if err != nil {
		zap_logger.Logger.Info("Error selecting districts =>", zap.Error(err))
		return districts, 500, ErrSelectDistricts, err
	}
	defer rows.Close()

	for rows.Next() {
		var district models.CatalogOption
		err := rows.Scan(
			&district.Id,
			&district.Name,
			&district.Code,
		)
		if err != nil {
			zap_logger.Logger.Info("Error scanning districts =>", zap.Error(err))
			return districts, 500, ErrSelectDistricts, err
		}
		districts = append(districts, district)
	}

	return districts, 200, Success, nil
}
