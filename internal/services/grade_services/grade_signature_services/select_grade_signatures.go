package grade_signature_services

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/models/grade_models"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"go.uber.org/zap"
)

func SelectGradeSignatures(ctx context.Context, gradeId string) ([]grade_models.GradeSignatures, int, string, error) {
	var gradeSignatures []grade_models.GradeSignatures

	const query = `
	SELECT
		"GS"."id",
		"GS"."grade_id",
		"G"."code" AS "grade_code",
		"G"."name" AS "grade_name",
		"GS"."signature_id",
		"S"."code" AS "signature_code",
		"S"."name" AS "signature_name",
		"GS"."assigned_by",
		"GS"."assigned_at"
	FROM
		"GradeSignatures" AS "GS"
	INNER JOIN 
		"Grade" AS "G" ON "GS"."grade_id" = "G"."id"
	INNER JOIN 
		"Signature" AS "S" ON "GS"."signature_id" = "S"."id"
	WHERE 
		"GS"."grade_id" = $1;
	`

	rows, err := db.Connx.QueryContext(ctx, query, gradeId)
	if err != nil {
		zap_logger.Logger.Info("Error selecting grade signatures =>", zap.Error(err))
		return gradeSignatures, 500, ErrSelectAssigns, err
	}
	defer rows.Close()

	for rows.Next() {
		var gradeSignature grade_models.GradeSignatures
		err := rows.Scan(
			&gradeSignature.Id,
			&gradeSignature.GradeId,
			&gradeSignature.GradeCode,
			&gradeSignature.GradeName,
			&gradeSignature.SignatureId,
			&gradeSignature.SignatureCode,
			&gradeSignature.SignatureName,
			&gradeSignature.AssignedBy,
			&gradeSignature.AssignedAt,
		)
		if err != nil {
			zap_logger.Logger.Info("Error selecting grade signatures =>", zap.Error(err))
			return gradeSignatures, 500, ErrSelectAssigns, err
		}
		gradeSignatures = append(gradeSignatures, gradeSignature)
	}

	return gradeSignatures, 200, Success, nil
}
