package grade_signature_services

import (
	"context"
	"errors"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func UnassignSignature(ctx context.Context, recordId string, gradeId string) (int, string, error) {
	const query = `
	DELETE FROM "GradeSignatures"
	WHERE "id" = $1
	AND "grade_id" = $2;
	`

	result, err := db.Connx.ExecContext(ctx, query, recordId, gradeId)
	if err != nil {
		zap_logger.Logger.Info("Error unassigning signature =>", zap.Error(err))
		return 500, ErrUnassign, err
	}

	row, err := result.RowsAffected()
	switch {
	case row == 0:
		return 404, NotFound, errors.New("grade signature not found")
	case err != nil:
		zap_logger.Logger.Info("Error unassigning signature =>", zap.Error(err))
		return 500, ErrUnassign, err
	}

	return 200, SuccessUnassign, nil
}
