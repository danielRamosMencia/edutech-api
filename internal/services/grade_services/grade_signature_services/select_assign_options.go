package grade_signature_services

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/models"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func SelectAssignOptions(ctx context.Context, gradeId string) ([]models.CatalogOption, int, string, error) {
	var availableSignatures []models.CatalogOption

	const query = `
	SELECT 
		"S"."id",
		"S"."code",
		"S"."name"
	FROM 
		"Signature" AS "S"
	LEFT JOIN
		"GradeSignatures" AS "GS" ON "GS"."signature_id" = "S"."id"
	AND "GS"."grade_id" = $1
	WHERE "GS"."id" IS NULL;
	`

	rows, err := db.Connx.QueryContext(ctx, query, gradeId)
	if err != nil {
		zap_logger.Logger.Info("Error selecting available signatures =>", zap.Error(err))
		return availableSignatures, 500, ErrSelectAssigns, err
	}
	defer rows.Close()

	for rows.Next() {
		var signature models.CatalogOption

		err := rows.Scan(
			&signature.Id,
			&signature.Code,
			&signature.Name,
		)
		if err != nil {
			zap_logger.Logger.Info("Error scanning available signatures =>", zap.Error(err))
			return availableSignatures, 500, ErrSelectAssigns, err
		}

		availableSignatures = append(availableSignatures, signature)
	}

	return availableSignatures, 200, Success, nil
}
