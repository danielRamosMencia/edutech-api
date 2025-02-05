package registration_status_services

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/models/registration_status_models"
	"github.com/danielRamosMencia/edutech-api/internal/utils"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func InsertRegistrationStatus(ctx context.Context, input registration_status_models.CreateRegistrationStatus, author string) (int, string, error) {
	const query = `
	INSERT INTO "RegistrationStatus" (
		"id", 
		"name", 
		"code",
		"active",
		"created_by",
		"created_at",
		"updated_at"
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		CURRENT_TIMESTAMP(),
		CURRENT_TIMESTAMP()
	);
	`

	id := utils.GenerateId()

	_, err := db.Connx.ExecContext(
		ctx,
		query,
		id,
		input.Name,
		input.Code,
		input.Active,
		author,
	)
	if err != nil {
		zap_logger.Logger.Info("Error inserting registration status =>", zap.Error(err))
		return 500, ErrInsert, err
	}

	return 201, SuccessInsert, nil
}
