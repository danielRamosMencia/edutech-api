package portal_user_services

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/models"
	"github.com/danielRamosMencia/edutech-api/internal/models/portal_user_models"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func SelectPortalUsers(ctx context.Context, pagination models.PaginationParams) ([]portal_user_models.PortalUser, int, string, error) {
	var portalUsers []portal_user_models.PortalUser

	const query = `
	SELECT
		"PU"."id",
		"PU"."username",
		"PU"."code",
		"PU"."active",
		"PU"."email",
		"PU"."employee_id",
		"E"."name" AS "employee_name",
		"E"."dni" AS "employee_dni",
		"PU"."role_id",
		"R"."name" AS "role",
		"PU"."created_by",
		"PU"."modified_by",
		"PU"."created_at",
		"PU"."updated_at"
	FROM
		"PortalUser" AS "PU"
	LEFT JOIN 
		"Employee" AS "E" ON "PU"."employee_id" = "E"."id"
	INNER JOIN
		"Role" AS "R" ON "PU"."role_id" = "R"."id"
	ORDER BY
		"PU"."username" ASC
	LIMIT $1 OFFSET $2;
	`

	rows, err := db.Connx.QueryContext(
		ctx,
		query,
		pagination.Limit,
		pagination.Offset,
	)
	if err != nil {
		zap_logger.Logger.Info("Error selecting portal users =>", zap.Error(err))
		return portalUsers, 500, ErrSelectPortalUsers, err
	}

	defer rows.Close()

	for rows.Next() {
		var portalUser portal_user_models.PortalUser
		err := rows.Scan(
			&portalUser.Id,
			&portalUser.Username,
			&portalUser.Code,
			&portalUser.Active,
			&portalUser.Email,
			&portalUser.EmployeeId,
			&portalUser.EmployeeName,
			&portalUser.EmployeeDni,
			&portalUser.RoleId,
			&portalUser.Role,
			&portalUser.CreatedBy,
			&portalUser.ModifiedBy,
			&portalUser.CreatedAt,
			&portalUser.UpdatedAt,
		)
		if err != nil {
			zap_logger.Logger.Info("Error scanning portal user =>", zap.Error(err))
			return portalUsers, 500, ErrSelectPortalUsers, err
		}

		portalUsers = append(portalUsers, portalUser)
	}

	return portalUsers, 200, Success, nil
}
