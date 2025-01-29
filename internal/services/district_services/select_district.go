package district_services

import (
	"context"
	"database/sql"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/models/district_models"
)

func SelectDistrict(ctx context.Context, districtId string) (district_models.District, int, string, error) {
	var district district_models.District
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
	WHERE
		"id" = $1
	LIMIT 1;
	`

	row := db.Connx.QueryRowContext(ctx, query, districtId)
	err := row.Scan(
		&district.Id,
		&district.Name,
		&district.Code,
		&district.Active,
		&district.CreatedBy,
		&district.ModifiedBy,
		&district.CreatedAt,
		&district.UpdatedAt,
	)

	switch {
	case err == sql.ErrNoRows:
		return district, 404, NotFound, err
	case err != nil:
		return district, 500, ErrSelectDistrict, err
	}

	return district, 200, Success, nil

}
