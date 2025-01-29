package country_services

import (
	"context"
	"database/sql"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/models/country_models"
)

var (
	country country_models.Country
)

func SelectCountry(ctx context.Context, countryId string) (country_models.Country, int, string, error) {
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
	WHERE
		"id" = $1
	LIMIT 1;
	`

	row := db.Connx.QueryRowContext(ctx, query, countryId)
	err := row.Scan(
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

	switch {
	case err == sql.ErrNoRows:
		return country, 404, NotFound, err
	case err != nil:
		return country, 500, ErrSelectCountry, err
	}

	return country, 200, Success, nil
}
