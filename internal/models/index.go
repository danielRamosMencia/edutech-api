package models

type SwicthActive struct {
	Active bool `json:"active" db:"active"`
}

type PaginationParams struct {
	Offset int
	Limit  int
}
