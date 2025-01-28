package country_services

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/models/country_models"
	"github.com/danielRamosMencia/edutech-api/internal/utils"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func InsertCountry(ctx context.Context, input country_models.CreateCoutry, author string) (int, string, error) {
	const query = `
	INSERT INTO "Country" (
		"id", 
		"name", 
		"active",
		"A2",
		"A3",
		"code",
		"created_by",
		"created_at",
		"updated_at"
	) VALUES (
		$1, 
		$2, 
		$3,
		$4,
		$5,
		$6,
		$7,
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
		input.Active,
		input.A2,
		input.A3,
		input.Code,
		author,
	)

	if err != nil {
		zap_logger.Logger.Info("Error inserting country =>", zap.Error(err))
		return 500, ErrInsert, err
	}

	return 201, SuccessInsert, nil
}
