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
	var storedPassword string

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

	row := db.Connx.QueryRowContext(
		ctx,
		query,
		input.Username,
		input.Email,
	)
	err := row.Scan(
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

	return sessionData, 200, Sucess, nil
}
