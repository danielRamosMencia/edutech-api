package models

type SwicthActive struct {
	Active bool `json:"active" db:"active"`
}

type PaginationParams struct {
	Offset int
	Limit  int
}

type SessionPermissions struct {
	PermissionId string `json:"permission_id" db:"permission_id"`
	Permission   string `json:"permission" db:"permission"`
	Code         string `json:"code" db:"code"`
}

type SessionData struct {
	Id          string                `json:"id" db:"id"`
	Username    string                `json:"username" db:"username"`
	Email       string                `json:"email" db:"email"`
	Active      bool                  `json:"active" db:"active"`
	RoleId      string                `json:"role_id" db:"role_id"`
	Role        string                `json:"role" db:"role"`
	Permissions *[]SessionPermissions `json:"permissions" db:"permissions"`
}
