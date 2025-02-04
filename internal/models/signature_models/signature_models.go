package signature_models

import "time"

type Signature struct {
	Id         string    `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	Code       string    `json:"code" db:"code"`
	Active     bool      `json:"active" db:"active"`
	CreatedBy  string    `json:"created_by" db:"created_by"`
	ModifiedBy *string   `json:"modified_by" db:"modified_by"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

type CreateSignature struct {
	Name   string `json:"name" db:"name" validate:"required"`
	Code   string `json:"code" db:"code" validate:"required"`
	Active bool   `json:"active" db:"active"`
}

type UpdateSignature struct {
	Name   string `json:"name" db:"name" validate:"required"`
	Code   string `json:"code" db:"code" validate:"required"`
	Active bool   `json:"active" db:"active"`
}
