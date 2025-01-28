package department_services

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/models"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func SelectDepartmentOptions(ctx context.Context) ([]models.CatalogOption, int, string, error) {
	var departments []models.CatalogOption

	const query = `
	SELECT
		"id", 
		"name",
		"code"
	FROM 
		"Department"
	WHERE
		"active" = true;
	`

	rows, err := db.Connx.QueryContext(ctx, query)
	if err != nil {
		zap_logger.Logger.Info("Error selecting departments =>", zap.Error(err))
		return departments, 500, ErrSelectDepartments, err
	}
	defer rows.Close()

	for rows.Next() {
		var department models.CatalogOption
		err := rows.Scan(
			&department.Id,
			&department.Name,
			&department.Code,
		)
		if err != nil {
			zap_logger.Logger.Info("Error scanning departments =>", zap.Error(err))
			return departments, 500, ErrSelectDepartments, err
		}
		departments = append(departments, department)
	}

	return departments, 200, Success, nil
}
