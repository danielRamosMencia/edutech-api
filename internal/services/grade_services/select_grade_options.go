package grade_services

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/models"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func SelectGradeOptions(ctx context.Context) ([]models.CatalogOption, int, string, error) {
	var grades []models.CatalogOption

	const query = `
	SELECT
		"id",
		"name",
		"code"
	FROM
		"Grade"
	WHERE
		"active" = true;
	`

	rows, err := db.Connx.QueryContext(ctx, query)
	if err != nil {
		zap_logger.Logger.Info("Error selecting grades =>", zap.Error(err))
		return grades, 500, ErrSelectGrades, err
	}
	defer rows.Close()

	for rows.Next() {
		var grade models.CatalogOption
		err := rows.Scan(
			&grade.Id,
			&grade.Name,
			&grade.Code,
		)
		if err != nil {
			zap_logger.Logger.Info("Error scanning grades =>", zap.Error(err))
			return grades, 500, ErrSelectGrades, err
		}
		grades = append(grades, grade)
	}

	return grades, 200, Success, nil
}
