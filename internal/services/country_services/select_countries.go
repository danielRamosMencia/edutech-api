package country_services

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/models"
	"github.com/danielRamosMencia/edutech-api/internal/models/country_models"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func SelectCountries(ctx context.Context, pagination models.PaginationParams) ([]country_models.Country, int, string, error) {
	var countries []country_models.Country

	const query = `
	SELECT
		"id", 
		"name", 
		"active",
		"A2",
		"A3",
		"code",
		"created_by",
		"modified_by",
		"created_at",
		"updated_at"
	FROM 
		"Country"
	LIMIT $1 OFFSET $2;
	`

	rows, err := db.Connx.QueryContext(
		ctx,
		query,
		pagination.Limit,
		pagination.Offset,
	)
	if err != nil {
		zap_logger.Logger.Info("Error selecting countries =>", zap.Error(err))
		return countries, 500, ErrSelectCountries, err
	}
	defer rows.Close()

	for rows.Next() {
		var country country_models.Country
		err := rows.Scan(
			&country.Id,
			&country.Name,
			&country.Active,
			&country.A2,
			&country.A3,
			&country.Code,
			&country.CreatedBy,
			&country.ModifiedBy,
			&country.CreatedAt,
			&country.UpdatedAt,
		)
		if err != nil {
			zap_logger.Logger.Info("Error selecting countries =>", zap.Error(err))
			return countries, 500, ErrSelectCountries, err
		}

		countries = append(countries, country)
	}

	return countries, 200, Success, nil
}
