package signature_services

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/models"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func SelectSignatureOptions(ctx context.Context) ([]models.CatalogOption, int, string, error) {
	var signatures []models.CatalogOption

	const query = `
	SELECT
		"id", 
		"name",
		"code"
	FROM 
		"Signature"
	WHERE
		"active" = true;
	`

	rows, err := db.Connx.QueryContext(ctx, query)
	if err != nil {
		zap_logger.Logger.Info("Error selecting signatures =>", zap.Error(err))
		return signatures, 500, ErrSelectSignatures, err
	}
	defer rows.Close()

	for rows.Next() {
		var signature models.CatalogOption
		err := rows.Scan(
			&signature.Id,
			&signature.Name,
			&signature.Code,
		)
		if err != nil {
			zap_logger.Logger.Info("Error scanning signatures =>", zap.Error(err))
			return signatures, 500, ErrSelectSignatures, err
		}
		signatures = append(signatures, signature)
	}

	return signatures, 200, Success, nil
}
