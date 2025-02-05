package employee_services

import (
	"context"
	"errors"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func ActiveEmployee(ctx context.Context, active bool, employeeId string, author string) (int, string, error) {
	var message string
	const query = `
	UPDATE "Employee" SET
		"active" = $1,
		"modified_by" = $2,
		"updated_at" = CURRENT_TIMESTAMP
	WHERE
		"id" = $3;
	`

	result, err := db.Connx.ExecContext(
		ctx,
		query,
		active,
		author,
		employeeId,
	)

	if err != nil {
		zap_logger.Logger.Info("Error updating employee =>", zap.Error(err))
		return 500, ErrUpdate, err
	}

	row, err := result.RowsAffected()
	switch {
	case row == 0:
		return 404, NotFound, errors.New("employee not found")
	case err != nil:
		zap_logger.Logger.Info("Error updating employee =>", zap.Error(err))
		return 500, ErrUpdate, err
	}

	if active {
		message = SuccessActivated
	} else {
		message = SuccessDesactivated
	}

	return 200, message, nil
}
