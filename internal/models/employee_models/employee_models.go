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
	PortalUser     *string    `json:"portal_user" db:"portal_user"`
	CreatedBy      string     `json:"created_by" db:"created_by"`
	ModifiedBy     *string    `json:"modified_by" db:"modified_by"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at" db:"updated_at"`
}

type CreateEmployee struct {
	Name           string  `json:"name" db:"name" validate:"required,max=255"`
	MiddleName     *string `json:"middle_name" db:"middle_name" validate:"max=255"`
	LastName       string  `json:"last_name" db:"last_name" validate:"required,max=255"`
	MiddleLastName *string `json:"middle_last_name" db:"middle_last_name" validate:"max=255"`
	Dni            string  `json:"dni" db:"dni" validate:"required,len=15"`
	Rtn            *string `json:"rtn" db:"rtn" validate:"len=16"`
	Address        *string `json:"address" db:"address" validate:"max=255"`
	Email          *string `json:"email" db:"email" validate:"email,max=255"`
	Phone          *string `json:"phone" db:"phone" validate:"max=30"`
	Birthdate      *string `json:"birthdate" db:"birthdate"`
	Active         bool    `json:"active" db:"active"`
	MunicipalityId string  `json:"municipality_id" db:"municipality_id" validate:"required"`
	InstitutionId  string  `json:"institution_id" db:"institution_id" validate:"required"`
}

type UpdateEmployee struct {
	Name           string  `json:"name" db:"name" validate:"required,max=255"`
	MiddleName     *string `json:"middle_name" db:"middle_name" validate:"max=255"`
	LastName       string  `json:"last_name" db:"last_name" validate:"required,max=255"`
	MiddleLastName *string `json:"middle_last_name" db:"middle_last_name" validate:"max=255"`
	Dni            string  `json:"dni" db:"dni" validate:"required,len=15"`
	Rtn            *string `json:"rtn" db:"rtn" validate:"len=16"`
	Address        *string `json:"address" db:"address" validate:"max=255"`
	Email          *string `json:"email" db:"email" validate:"email,max=255"`
	Phone          *string `json:"phone" db:"phone" validate:"max=30"`
	Birthdate      *string `json:"birthdate" db:"birthdate"`
	Active         bool    `json:"active" db:"active"`
	MunicipalityId string  `json:"municipality_id" db:"municipality_id" validate:"required"`
	InstitutionId  string  `json:"institution_id" db:"institution_id" validate:"required"`
}
