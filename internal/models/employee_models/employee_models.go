package employee_models

import "time"

type Employee struct {
	Id             string     `json:"id" db:"id"`
	Name           string     `json:"name" db:"name"`
	MiddleName     *string    `json:"middle_name" db:"middle_name"`
	LastName       string     `json:"last_name" db:"last_name"`
	MiddleLastName *string    `json:"middle_last_name" db:"middle_last_name"`
	Dni            string     `json:"dni" db:"dni"`
	Rtn            *string    `json:"rtn" db:"rtn"`
	Address        *string    `json:"address" db:"address"`
	Email          *string    `json:"email" db:"email"`
	Phone          *string    `json:"phone" db:"phone"`
	Birthdate      *time.Time `json:"birthdate" db:"birthdate"`
	Active         bool       `json:"active" db:"active"`
	MunicipalityId string     `json:"municipality_id" db:"municipality_id"`
	Municipality   string     `json:"municipality" db:"municipality"`
	Department     string     `json:"department" db:"department"`
	InstitutionId  string     `json:"institution_id" db:"institution_id"`
	Institution    string     `json:"institution" db:"institution"`
	PortalUser     string     `json:"portal_user" db:"portal_user"`
	CreatedBy      string     `json:"created_by" db:"created_by"`
	ModifiedBy     *string    `json:"modified_by" db:"modified_by"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at" db:"updated_at"`
}

type CreateEmployee struct {
	Name           string  `json:"name" db:"name" validate:"required"`
	MiddleName     *string `json:"middle_name" db:"middle_name"`
	LastName       string  `json:"last_name" db:"last_name" validate:"required"`
	MiddleLastName *string `json:"middle_last_name" db:"middle_last_name"`
	Dni            string  `json:"dni" db:"dni" validate:"required"`
	Rtn            *string `json:"rtn" db:"rtn"`
	Address        *string `json:"address" db:"address"`
	Email          *string `json:"email" db:"email"`
	Phone          *string `json:"phone" db:"phone"`
	Birthdate      *string `json:"birthdate" db:"birthdate"`
	Active         bool    `json:"active" db:"active"`
	MunicipalityId string  `json:"municipality_id" db:"municipality_id" validate:"required"`
	InstitutionId  string  `json:"institution_id" db:"institution_id" validate:"required"`
}

type UpdateEmployee struct {
	Name           string  `json:"name" db:"name" validate:"required"`
	MiddleName     *string `json:"middle_name" db:"middle_name"`
	LastName       string  `json:"last_name" db:"last_name" validate:"required"`
	MiddleLastName *string `json:"middle_last_name" db:"middle_last_name"`
	Dni            string  `json:"dni" db:"dni" validate:"required"`
	Rtn            *string `json:"rtn" db:"rtn"`
	Address        *string `json:"address" db:"address"`
	Email          *string `json:"email" db:"email"`
	Phone          *string `json:"phone" db:"phone"`
	Birthdate      *string `json:"birthdate" db:"birthdate"`
	Active         bool    `json:"active" db:"active"`
	MunicipalityId string  `json:"municipality_id" db:"municipality_id" validate:"required"`
	InstitutionId  string  `json:"institution_id" db:"institution_id" validate:"required"`
}
