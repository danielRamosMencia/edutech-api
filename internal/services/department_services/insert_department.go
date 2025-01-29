package department_services

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/models/department_models"
	"github.com/danielRamosMencia/edutech-api/internal/utils"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func InsertDepartment(ctx context.Context, input department_models.CreateDepartment, author string) (int, string, error) {
	const query = `
	INSERT INTO "Department" (
		"id", 
		"name", 
		"code",
		"active",
		"country_id",
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
		input.Active,
		input.CountryId,
		author,
	)

	if err != nil {
		zap_logger.Logger.Info("Error inserting department =>", zap.Error(err))
		return 500, ErrInsert, err
	}

	return 201, SuccessInsert, nil
}
