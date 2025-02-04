package signature_services

import (
	"context"
	"database/sql"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/models/signature_models"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func SelectSignature(ctx context.Context, signatureId string) (signature_models.Signature, int, string, error) {
	var signature signature_models.Signature

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
		"Signature"
	WHERE
		"id" = $1
	LIMIT 1;
	`

	row := db.Connx.QueryRowContext(ctx, query, signatureId)
	err := row.Scan(
		&signature.Id,
		&signature.Name,
		&signature.Code,
		&signature.Active,
		&signature.CreatedBy,
		&signature.ModifiedBy,
		&signature.CreatedAt,
		&signature.UpdatedAt,
	)

	switch {
	case err == sql.ErrNoRows:
		return signature, 404, NotFound, err
	case err != nil:
		zap_logger.Logger.Info("Error selecting signature =>", zap.Error(err))
		return signature, 500, ErrSelectSignature, err
	}

	return signature, 200, Success, nil
}
