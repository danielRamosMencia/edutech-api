package signature_services

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/models"
	"github.com/danielRamosMencia/edutech-api/internal/models/signature_models"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func SelectSignatures(ctx context.Context, pagination models.PaginationParams) ([]signature_models.Signature, int, string, error) {
	var signatures []signature_models.Signature

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
	LIMIT $1 OFFSET $2;
	`

	rows, err := db.Connx.QueryContext(
		ctx,
		query,
		pagination.Limit,
		pagination.Offset,
	)
	if err != nil {
		zap_logger.Logger.Info("Error selecting signatures =>", zap.Error(err))
		return signatures, 500, ErrSelectSignatures, err
	}
	defer rows.Close()

	for rows.Next() {
		var signature signature_models.Signature
		err := rows.Scan(
			&signature.Id,
			&signature.Name,
			&signature.Code,
			&signature.Active,
			&signature.CreatedBy,
			&signature.ModifiedBy,
			&signature.CreatedAt,
			&signature.UpdatedAt,
		)
		if err != nil {
			zap_logger.Logger.Info("Error scanning signatures =>", zap.Error(err))
			return signatures, 500, ErrSelectSignatures, err
		}
		signatures = append(signatures, signature)
	}

	return signatures, 200, Success, nil
}
