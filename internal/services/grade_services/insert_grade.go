package grade_services

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/models/grade_models"
	"github.com/danielRamosMencia/edutech-api/internal/utils"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func InsertGrade(ctx context.Context, input grade_models.CreateGrade, author string) (int, string, error) {
	const query = `
	INSERT INTO "Grade" (
		"id", 
		"name", 
		"code",
		"active",
		"grade_number",
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
	)
	`

	id := utils.GenerateId()
	_, err := db.Connx.ExecContext(
		ctx,
		query,
		id,
		input.Name,
		input.Code,
		input.Active,
		input.GradeNumber,
		author,
	)

	if err != nil {
		zap_logger.Logger.Info("Error inserting grade =>", zap.Error(err))
		return 500, ErrInsert, err
	}

	return 200, SuccessInsert, nil
}
