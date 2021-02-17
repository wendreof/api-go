package models

import "database/sql"

//Local represents the Place table on DB
type Local struct {
	Country   string         `json:"country" db:country`
	City      sql.NullString `json:"city" db:city`
	PhoneCode int            `json:phonecode db:phonecode`
}
