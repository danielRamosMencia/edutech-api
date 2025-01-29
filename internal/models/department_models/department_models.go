package department_models

import "time"

type Department struct {
	Id         string    `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	Code       string    `json:"code" db:"code"`
	Active     bool      `json:"active" db:"active"`
	CountryId  string    `json:"country_id" db:"country_id"`
	Country    string    `json:"country" db:"country"`
	CreatedBy  string    `json:"created_by" db:"created_by"`
	ModifiedBy *string   `json:"modified_by" db:"modified_by"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

type CreateDepartment struct {
	Name      string `json:"name" validate:"required"`
	Code      string `json:"code" validate:"required,len=2"`
	Active    bool   `json:"active"`
	CountryId string `json:"country_id" validate:"required"`
}

type UpdateDeparment struct {
	Name      string `json:"name" validate:"required"`
	Code      string `json:"code" validate:"required,len=2"`
	Active    bool   `json:"active"`
	CountryId string `json:"country_id" validate:"required"`
}
