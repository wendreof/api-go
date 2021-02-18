package models

import "database/sql"

//Local represents the Place table on DB
type Local struct {
	Country   string         `json:"country"`
	City      sql.NullString `json:"city"`
	PhoneCode int            `json:phonecode"`
}
