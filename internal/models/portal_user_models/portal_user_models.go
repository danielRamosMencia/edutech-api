package portal_user_models

import "time"

type PortalUser struct {
	Id           string    `json:"id" db:"id"`
	Username     string    `json:"username" db:"username"`
	Code         string    `json:"code" db:"code"`
	Active       bool      `json:"active" db:"active"`
	Email        string    `json:"email" db:"email"`
	EmployeeId   string    `json:"employee_id" db:"employee_id"`
	EmployeeName string    `json:"employee_name" db:"employee_name"`
	EmployeeDni  string    `json:"employee_dni" db:"employee_dni"`
	RoleId       string    `json:"role_id" db:"role_id"`
	Role         string    `json:"role" db:"role"`
	CreatedBy    string    `json:"created_by" db:"created_by"`
	ModifiedBy   *string   `json:"modified_by" db:"modified_by"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}
