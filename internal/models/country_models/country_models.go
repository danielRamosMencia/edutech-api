package country_models

import "time"

type Country struct {
	Id         string    `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	Active     bool      `json:"active" db:"active"`
	A2         string    `json:"A2" db:"A2"`
	A3         string    `json:"A3" db:"A3"`
	Code       string    `json:"code" db:"code"`
	CreatedBy  string    `json:"created_by" db:"created_by"`
	ModifiedBy *string   `json:"modified_by" db:"modified_by"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}
