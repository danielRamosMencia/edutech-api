package grade_models

import "time"

type Grade struct {
	Id          string    `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Code        string    `json:"code" db:"code"`
	Active      bool      `json:"active" db:"active"`
	GradeNumber int       `json:"grade_number" db:"grade_number"`
	CreatedBy   string    `json:"created_by" db:"created_by"`
	ModifiedBy  *string   `json:"modified_by" db:"modified_by"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type CreateGrade struct {
	Name        string `json:"name" db:"name"`
	Code        string `json:"code" db:"code"`
	Active      bool   `json:"active" db:"active"`
	GradeNumber int    `json:"grade_number" db:"grade_number"`
}

type UpdateGrade struct {
	Name        string `json:"name" db:"name"`
	Code        string `json:"code" db:"code"`
	Active      bool   `json:"active" db:"active"`
	GradeNumber int    `json:"grade_number" db:"grade_number"`
}
