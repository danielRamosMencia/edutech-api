package signature_services

import (
	"context"
	"errors"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/models/signature_models"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func UpdateSignature(ctx context.Context, signatureId string, input signature_models.UpdateSignature, author string) (int, string, error) {
	const query = `
	UPDATE "Signature" SET
		"name" = $1,
		"code" = $2,
		"active" = $3,
		"modified_by" = $4,
		"updated_at" = CURRENT_TIMESTAMP
	WHERE
		"id" = $5;
	`

	result, err := db.Connx.ExecContext(
		ctx,
		query,
		input.Name,
		input.Code,
		input.Active,
		author,
		signatureId,
	)

	if err != nil {
		zap_logger.Logger.Info("Error updating signature =>", zap.Error(err))
		return 500, ErrUpdate, err
	}

	row, err := result.RowsAffected()
	switch {
	case row == 0:
		return 404, NotFound, errors.New("signature not found")
	case err != nil:
		zap_logger.Logger.Info("Error updating signature =>", zap.Error(err))
		return 500, ErrUpdate, err
	}

	return 200, SuccessUpdate, nil
}
