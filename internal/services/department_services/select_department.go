package department_services

import (
	"context"
	"database/sql"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/models/department_models"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func SelectDepartment(ctx context.Context, departmentId string) (department_models.Department, int, string, error) {
	var department department_models.Department

	const query = `
	SELECT
		"D"."id",
		"D"."name",
		"D"."code",
		"D"."active",
		"C"."id" AS "country_id",
		"C"."name" AS "country",
		"D"."created_by",
		"D"."modified_by",
		"D"."created_at",
		"D"."updated_at"
	FROM
		"Department" AS "D"
	INNER JOIN 
		"Country" AS "C" ON "D"."country_id" = "C"."id"
	WHERE
		"D"."id" = $1;
	`

	row := db.Connx.QueryRowContext(ctx, query, departmentId)
	err := row.Scan(
		&department.Id,
		&department.Name,
		&department.Code,
		&department.Active,
		&department.CountryId,
		&department.Country,
		&department.CreatedBy,
		&department.ModifiedBy,
		&department.CreatedAt,
		&department.UpdatedAt,
	)

	switch {
	case err == sql.ErrNoRows:
		return department, 404, NotFound, err
	case err != nil:
		zap_logger.Logger.Info("Error selecting department =>", zap.Error(err))
		return department, 500, ErrSelectDepartment, err
	}

	return department, 200, Success, nil
}
