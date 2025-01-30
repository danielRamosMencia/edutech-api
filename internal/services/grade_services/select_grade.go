package grade_services

import (
	"context"
	"database/sql"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/models/grade_models"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func SelectGrade(ctx context.Context, gradeId string) (grade_models.Grade, int, string, error) {
	var grade grade_models.Grade
	const query = `
	SELECT
		"id", 
		"name",
		"code",
		"active",
		"grade_number",
		"created_by",
		"modified_by",
		"created_at",
		"updated_at"
	FROM 
		"Grade"
	WHERE
		"id" = $1
	LIMIT 1;
	`

	row := db.Connx.QueryRowContext(ctx, query, gradeId)
	err := row.Scan(
		&grade.Id,
		&grade.Name,
		&grade.Code,
		&grade.Active,
		&grade.GradeNumber,
		&grade.CreatedBy,
		&grade.ModifiedBy,
		&grade.CreatedAt,
		&grade.UpdatedAt,
	)

	switch {
	case err == sql.ErrNoRows:
		return grade, 404, NotFound, err
	case err != nil:
		zap_logger.Logger.Info("Error selecting grade =>", zap.Error(err))
		return grade, 500, ErrSelectGrade, err
	}

	return grade, 200, Success, nil
}
