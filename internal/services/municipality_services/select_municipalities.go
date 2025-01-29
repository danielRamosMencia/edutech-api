package municipality_services

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/models"
	"github.com/danielRamosMencia/edutech-api/internal/models/municipality_models"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func SelectMunicipalities(ctx context.Context, pagination models.PaginationParams) ([]municipality_models.Municipality, int, string, error) {
	var municipalities []municipality_models.Municipality
	const query = `
	SELECT
		"M"."id", 
		"M"."name", 
		"M"."code",
		"M"."active",
		"D"."id" AS "department_id",
		"D"."name" AS "department",
		"M"."created_by",
		"M"."modified_by",
		"M"."created_at",
		"M"."updated_at"
	FROM 
		"Municipality" AS "M"
	INNER JOIN 
	"Department" AS "D" ON "M"."department_id" = "D"."id"
	LIMIT $1 OFFSET $2;
	`

	rows, err := db.Connx.QueryContext(
		ctx,
		query,
		pagination.Limit,
		pagination.Offset,
	)
	if err != nil {
		zap_logger.Logger.Info("Error selecting municipalities =>", zap.Error(err))
		return municipalities, 500, ErrSelectMunicipalities, err
	}
	defer rows.Close()

	for rows.Next() {
		var municipality municipality_models.Municipality
		err := rows.Scan(
			&municipality.Id,
			&municipality.Name,
			&municipality.Code,
			&municipality.Active,
			&municipality.DepartmentId,
			&municipality.Department,
			&municipality.CreatedBy,
			&municipality.ModifiedBy,
			&municipality.CreatedAt,
			&municipality.UpdatedAt,
		)

		if err != nil {
			zap_logger.Logger.Info("Error scanning municipalities =>", zap.Error(err))
			return municipalities, 500, ErrSelectMunicipalities, err
		}
		municipalities = append(municipalities, municipality)
	}

	return municipalities, 200, Success, nil
}
