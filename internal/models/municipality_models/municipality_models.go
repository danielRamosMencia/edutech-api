package municipality_models

import "time"

type Municipality struct {
	Id           string    `json:"id" db:"id"`
	Name         string    `json:"name" db:"name"`
	Code         string    `json:"code" db:"code"`
	Active       bool      `json:"active" db:"active"`
	DepartmentId string    `json:"department_id" db:"department_id"`
	Department   string    `json:"department" db:"department"`
	CreatedBy    string    `json:"created_by" db:"created_by"`
	ModifiedBy   *string   `json:"modified_by" db:"modified_by"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

type CreateMunicipality struct {
	Name         string `json:"name" db:"name" validate:"required"`
	Code         string `json:"code" db:"code" validate:"required,len=4"`
	Active       bool   `json:"active" db:"active"`
	DepartmentId string `json:"department_id" db:"department_id" validate:"required"`
}

type UpdateMunicipality struct {
	Name         string `json:"name" db:"name" validate:"required"`
	Code         string `json:"code" db:"code" validate:"required,len=4"`
	Active       bool   `json:"active" db:"active"`
	DepartmentId string `json:"department_id" db:"department_id" validate:"required"`
}
