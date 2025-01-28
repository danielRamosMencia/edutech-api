package auth_services

import (
	"context"
	"database/sql"
	"errors"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/models"
	"github.com/danielRamosMencia/edutech-api/internal/models/auth_models"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func SelectSessionData(ctx context.Context, input auth_models.Login) (models.SessionData, int, string, error) {
	var sessionData models.SessionData
	var permissions []models.SessionPermissions
	var storedPassword string

	trx, err := db.Connx.BeginTx(ctx, nil)
	if err != nil {
		zap_logger.Logger.Info("Error initializing transaction", zap.Error(err))
		return sessionData, 500, ErrSelectUser, err
	}

	defer func() {
		if p := recover(); p != nil {
			trx.Rollback()
			panic(p)
		} else if err != nil {
			trx.Rollback()
		} else {
			err = trx.Commit()
		}
	}()

	const query = `
	SELECT
		"PU"."id",
		"PU"."username",
		"PU"."email",
		"PU"."active",
		"PU"."password",
		"R"."id" AS "role_id",
		"R"."name" AS "role"
	FROM 
		"PortalUser" as "PU"
	INNER JOIN 
		"Role" AS "R" ON "PU"."role_id" = "R"."id"
	WHERE
		"PU"."username" = $1 AND "PU"."email" = $2
	LIMIT 1;
	`

	row := trx.QueryRowContext(
		ctx,
		query,
		input.Username,
		input.Email,
	)
	err = row.Scan(
		&sessionData.Id,
		&sessionData.Username,
		&sessionData.Email,
		&sessionData.Active,
		&storedPassword,
		&sessionData.RoleId,
		&sessionData.Role,
	)

	switch {
	case err == sql.ErrNoRows:
		return sessionData, 404, BadCredentials, err
	case err != nil:
		zap_logger.Logger.Info("Error selecting session data =>", zap.Error(err))
		return sessionData, 500, ErrSelectUser, err
	}

	if !sessionData.Active {
		return sessionData, 401, InactiveUser, errors.New("inactive user")
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(input.Password))
	if err != nil {
		zap_logger.Logger.Info("Error comparing password =>", zap.Error(err))
		return sessionData, 401, BadPassword, err
	}

	const permissionsQuery = `
	SELECT
		"P"."id",
		"P"."name",
		"P"."code"
	FROM 
		"RolePermissions" AS "RP"
	JOIN 
		"Permission" AS "P" ON "RP"."permission_id" = "P"."id"
	WHERE
		"RP"."role_id" = $1;
	`

	rows, err := trx.QueryContext(ctx, permissionsQuery, sessionData.RoleId)
	if err != nil {
		zap_logger.Logger.Info("Error selecting permissions =>", zap.Error(err))
		return sessionData, 500, ErrSelectUser, err
	}

	for rows.Next() {
		var permission models.SessionPermissions
		err = rows.Scan(
			&permission.PermissionId,
			&permission.Permission,
			&permission.Code,
		)
		if err != nil {
			zap_logger.Logger.Info("Error scanning permissions =>", zap.Error(err))
			return sessionData, 500, ErrSelectUser, err
		}
		permissions = append(permissions, permission)
	}

	sessionData.Permissions = &permissions

	return sessionData, 200, Sucess, nil
}
