package municipality_services

import (
	"context"
	"database/sql"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/models/municipality_models"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func SelectMunicipality(ctx context.Context, municipalityId string) (municipality_models.Municipality, int, string, error) {
	var municipality municipality_models.Municipality

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
	WHERE "M"."id" = $1
	LIMIT 1;
	`

	row := db.Connx.QueryRowContext(ctx, query, municipalityId)
	err := row.Scan(
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

	switch {
	case err == sql.ErrNoRows:
		return municipality, 404, NotFound, err
	case err != nil:
		zap_logger.Logger.Info("Error selecting municipality =>", zap.Error(err))
		return municipality, 500, ErrSelectMunicipality, err
	}

	return municipality, 200, Success, nil
}
