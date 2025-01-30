package grade_services

import (
	"context"
	"errors"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func DeleteGrade(ctx context.Context, gradeId string) (int, string, error) {
	const query = `
	DELETE FROM "Grade"
	WHERE
		"id" = $1;
	`

	result, err := db.Connx.ExecContext(ctx, query, gradeId)
	if err != nil {
		zap_logger.Logger.Info("Error deleting grade =>", zap.Error(err))
		return 500, ErrDelete, err
	}

	row, err := result.RowsAffected()
	switch {
	case row == 0:
		return 404, NotFound, errors.New("grade not found")
	case err != nil:
		zap_logger.Logger.Info("Error deleting grade =>", zap.Error(err))
		return 500, ErrDelete, err
	}

	return 200, SuccessDelete, nil
}
