package models

// Pagination model that has page and limit
type Pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}
