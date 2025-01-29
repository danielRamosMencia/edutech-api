package district_services

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/models"
	"github.com/danielRamosMencia/edutech-api/internal/models/district_models"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func SelectDistricts(ctx context.Context, pagination models.PaginationParams) ([]district_models.District, int, string, error) {
	var districts []district_models.District

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
		"District"
	LIMIT $1 OFFSET $2;
	`

	rows, err := db.Connx.QueryContext(
		ctx,
		query,
		pagination.Limit,
		pagination.Offset,
	)
	if err != nil {
		zap_logger.Logger.Info("Error selecting districts =>", zap.Error(err))
		return districts, 500, ErrSelectDistricts, err
	}
	defer rows.Close()

	for rows.Next() {
		var district district_models.District
		err := rows.Scan(
			&district.Id,
			&district.Name,
			&district.Code,
			&district.Active,
			&district.CreatedBy,
			&district.ModifiedBy,
			&district.CreatedAt,
			&district.UpdatedAt,
		)
		if err != nil {
			zap_logger.Logger.Info("Error scanning districts =>", zap.Error(err))
			return districts, 500, ErrSelectDistricts, err
		}
		districts = append(districts, district)
	}

	return districts, 200, Success, nil
}
