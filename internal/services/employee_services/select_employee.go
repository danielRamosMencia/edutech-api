package employee_services

import (
	"context"
	"database/sql"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/models/employee_models"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func SelectEmployee(ctx context.Context, employeeId string) (employee_models.Employee, int, string, error) {
	var employee employee_models.Employee

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
	WHERE
		"E"."id" = $1
	LIMIT 1;
	`

	row := db.Connx.QueryRowContext(ctx, query, employeeId)
	err := row.Scan(
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

	switch {
	case err == sql.ErrNoRows:
		return employee, 404, NotFound, err
	case err != nil:
		zap_logger.Logger.Info("Error getting employee =>", zap.Error(err))
		return employee, 500, ErrSelectEmployee, err
	}

	return employee, 200, Success, nil
}
