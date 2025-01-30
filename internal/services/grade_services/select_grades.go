package grade_services

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/models"
	"github.com/danielRamosMencia/edutech-api/internal/models/grade_models"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func SelectGrades(ctx context.Context, pagination models.PaginationParams) ([]grade_models.Grade, int, string, error) {
	var grades []grade_models.Grade

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
	LIMIT $1 OFFSET $2;
	`

	rows, err := db.Connx.QueryContext(
		ctx,
		query,
		pagination.Limit,
		pagination.Offset,
	)
	if err != nil {
		zap_logger.Logger.Info("Error selecting grades =>", zap.Error(err))
		return grades, 500, ErrSelectGrades, err
	}
	defer rows.Close()

	for rows.Next() {
		var grade grade_models.Grade
		err := rows.Scan(
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
		if err != nil {
			zap_logger.Logger.Info("Error scanning grade =>", zap.Error(err))
			return grades, 500, ErrSelectGrades, err
		}
		grades = append(grades, grade)
	}

	return grades, 200, Success, nil
}
