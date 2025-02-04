package grade_models

import "time"

type GradeSignatures struct {
	Id            string    `json:"id" db:"id"`
	GradeId       string    `json:"grade_id" db:"grade_id"`
	GradeCode     string    `json:"grade_code" db:"grade_code"`
	GradeName     string    `json:"grade_name" db:"grade_name"`
	SignatureId   string    `json:"signature_id" db:"signature_id"`
	SignatureCode string    `json:"signature_code" db:"signature_code"`
	SignatureName string    `json:"signature_name" db:"signature_name"`
	AssignedBy    string    `json:"assigned_by" db:"assigned_by"`
	AssignedAt    time.Time `json:"assigned_at" db:"assigned_at"`
}

type AssignSignature struct {
	SignatureId string `json:"signature_id" db:"signature_id" validate:"required"`
}
