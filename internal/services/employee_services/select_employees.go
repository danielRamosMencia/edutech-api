package employee_services

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/models"
	"github.com/danielRamosMencia/edutech-api/internal/models/employee_models"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func SelectEmployees(ctx context.Context, pagination models.PaginationParams) ([]employee_models.Employee, int, string, error) {
	var employees []employee_models.Employee

	const query = `
	SELECT
		"E"."id",
		"E"."name",
		"E"."middle_name",
		"E"."last_name",
		"E"."middle_last_name",
		"E"."dni",
		"E"."rtn",
		"E"."address",
		"E"."email",
		"E"."phone",
		"E"."birthdate",
		"E"."active",
		"M"."id" AS "municipality_id",
		"M"."name" AS "municipality",
		"D"."name" AS "department",
		"I"."id" AS "institution_id",
		"I"."name" AS "institution",
		"PU"."id" AS "portal_user"
		"E"."created_by",
		"E"."modified_by",
		"E"."created_at",
		"E"."modified_at"
	FROM
		"Employee" AS "E"
	INNER JOIN 
		"Municipality" AS "M" ON "E"."municipality_id" = "M"."id"
	INNER JOIN 
		"Department" AS "D" ON "M"."department_id" = "D"."id"
	INNER JOIN 
		"Institution" AS "I" ON "E"."institution_id" = "I"."id"
	INNER JOIN 
		"PortalUser" AS "PU" ON "E"."id" = "PU"."employee_id"
	LIMIT $1 OFFSET $2;
	`

	rows, err := db.Connx.QueryContext(ctx, query, pagination.Limit, pagination.Offset)
	if err != nil {
		zap_logger.Logger.Info("Error selecting employees =>", zap.Error(err))
		return employees, 500, ErrSelectEmployees, err
	}
	defer rows.Close()

	for rows.Next() {
		var employee employee_models.Employee
		err := rows.Scan(
			&employee.Id,
			&employee.Name,
			&employee.MiddleName,
			&employee.LastName,
			&employee.MiddleLastName,
			&employee.Dni,
			&employee.Rtn,
			&employee.Address,
			&employee.Email,
			&employee.Phone,
			&employee.Birthdate,
			&employee.Active,
			&employee.MunicipalityId,
			&employee.Municipality,
			&employee.Department,
			&employee.InstitutionId,
			&employee.Institution,
			&employee.PortalUser,
			&employee.CreatedBy,
			&employee.ModifiedBy,
			&employee.CreatedAt,
			&employee.UpdatedAt,
		)
		if err != nil {
			zap_logger.Logger.Info("Error scanning employee =>", zap.Error(err))
			return employees, 500, ErrSelectEmployees, err
		}

		employees = append(employees, employee)
	}

	return employees, 200, Success, nil
}
