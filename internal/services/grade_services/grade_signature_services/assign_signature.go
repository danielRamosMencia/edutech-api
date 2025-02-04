package grade_signature_services

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/models/grade_models"
	"github.com/danielRamosMencia/edutech-api/internal/utils"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func AssignSignature(ctx context.Context, input grade_models.AssignSignature, gradeId string, author string) (int, string, error) {
	const query = `
	INSERT INTO "GradeSignatures" (
		"id", 
		"grade_id", 
		"signature_id",
		"assigned_by",
		"assigned_at"
	) VALUES (
		$1, 
		$2, 
		$3, 
		$4, 
		CURRENT_TIMESTAMP
	);
	`

	id := utils.GenerateId()

	zap_logger.Logger.Info("Assigning signature =>", zap.String("gradeId", gradeId))

	_, err := db.Connx.ExecContext(
		ctx,
		query,
		id,
		gradeId,
		input.SignatureId,
		author,
	)
	if err != nil {
		zap_logger.Logger.Info("Error assigning signature =>", zap.Error(err))
		return 500, ErrAssign, err
	}

	return 200, SuccessAssign, nil
}
