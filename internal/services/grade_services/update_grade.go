package grade_services

import (
	"context"
	"errors"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/models/grade_models"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func UpdateGrade(ctx context.Context, input grade_models.UpdateGrade, gradeId string, author string) (int, string, error) {
	const query = `
	UPDATE "Grade" SET
		"name" = $1,
		"code" = $2,
		"active" = $3,
		"grade_number" = $4,
		"modified_by" = $5,
		"updated_at" = CURRENT_TIMESTAMP
	WHERE
		"id" = $6;
	`

	result, err := db.Connx.ExecContext(
		ctx,
		query,
		input.Name,
		input.Code,
		input.Active,
		input.GradeNumber,
		author,
		gradeId,
	)

	if err != nil {
		zap_logger.Logger.Info("Error updating grade =>", zap.Error(err))
		return 500, ErrUpdate, err
	}

	row, err := result.RowsAffected()
	switch {
	case row == 0:
		return 404, NotFound, errors.New("grade not found")
	case err != nil:
		zap_logger.Logger.Info("Error updating grade =>", zap.Error(err))
		return 500, ErrUpdate, err
	}

	return 200, SuccessUpdate, nil
}
