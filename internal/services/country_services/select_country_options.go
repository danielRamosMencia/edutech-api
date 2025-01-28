package country_services

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/models"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func SelectCountryOptions(ctx context.Context) ([]models.CatalogOption, int, string, error) {
	var countries []models.CatalogOption

	const query = `
	SELECT
		"id", 
		"name",
		"code"
	FROM 
		"Country"
	WHERE
		"active" = true;
	`

	rows, err := db.Connx.QueryContext(ctx, query)
	if err != nil {
		zap_logger.Logger.Info("Error selecting countries =>", zap.Error(err))
		return countries, 500, ErrSelectCountries, err
	}
	defer rows.Close()

	for rows.Next() {
		var country models.CatalogOption
		err := rows.Scan(
			&country.Id,
			&country.Name,
			&country.Code,
		)
		if err != nil {
			zap_logger.Logger.Info("Error scanning countries =>", zap.Error(err))
			return countries, 500, ErrSelectCountries, err
		}

		countries = append(countries, country)
	}

	return countries, 200, Success, nil
}
