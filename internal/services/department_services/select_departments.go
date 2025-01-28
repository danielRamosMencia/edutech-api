package department_services

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/models"
	"github.com/danielRamosMencia/edutech-api/internal/models/department_models"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func SelectDepartments(ctx context.Context, pagination models.PaginationParams) ([]department_models.Department, int, string, error) {
	var departments []department_models.Department

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
	LIMIT $1 OFFSET $2;
	`

	rows, err := db.Connx.QueryContext(
		ctx,
		query,
		pagination.Limit,
		pagination.Offset,
	)
	if err != nil {
		zap_logger.Logger.Info("Error selecting departments =>", zap.Error(err))
		return departments, 500, ErrSelectDepartments, err
	}
	defer rows.Close()

	for rows.Next() {
		var department department_models.Department
		err := rows.Scan(
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
		if err != nil {
			zap_logger.Logger.Info("Error scanning departments =>", zap.Error(err))
			return departments, 500, ErrSelectDepartments, err
		}
		departments = append(departments, department)
	}

	return departments, 200, Success, nil
}
