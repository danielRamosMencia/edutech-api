package district_services

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/models/district_models"
	"github.com/danielRamosMencia/edutech-api/internal/utils"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func InsertDistrict(ctx context.Context, input district_models.CreateDistrict, author string) (int, string, error) {
	const query = `
	INSERT INTO "District" (
		"id", 
		"name", 
		"code",
		"created_by",
		"created_at",
		"updated_at"
	) VALUES (
		$1, 
		$2, 
		$3,
		$4,
		CURRENT_TIMESTAMP,
		CURRENT_TIMESTAMP
	);
	`

	id := utils.GenerateId()
	_, err := db.Connx.ExecContext(
		ctx,
		query,
		id,
		input.Name,
		input.Code,
		author,
	)

	if err != nil {
		zap_logger.Logger.Info("Error inserting district =>", zap.Error(err))
		return 500, ErrInsert, err
	}

	return 200, SuccessInsert, nil
}
