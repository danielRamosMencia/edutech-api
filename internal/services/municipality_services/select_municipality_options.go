package municipality_services

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/models"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func SelectMunicipalityOptions(ctx context.Context) ([]models.CatalogOption, int, string, error) {
	var municipalities []models.CatalogOption
	const query = `
	SELECT 
		"id", 
		"name",
		"code"
	FROM 
		"Municipality"
	WHERE
		"active" = true;
	`

	rows, err := db.Connx.QueryContext(ctx, query)
	if err != nil {
		zap_logger.Logger.Info("Error selecting municipalities =>", zap.Error(err))
		return municipalities, 500, ErrSelectMunicipalities, err
	}
	defer rows.Close()

	for rows.Next() {
		var municipality models.CatalogOption
		err := rows.Scan(
			&municipality.Id,
			&municipality.Name,
			&municipality.Code,
		)
		if err != nil {
			zap_logger.Logger.Info("Error scanning municipalities =>", zap.Error(err))
			return municipalities, 500, ErrSelectMunicipalities, err
		}
		municipalities = append(municipalities, municipality)
	}

	return municipalities, 200, Success, nil
}
