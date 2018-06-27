package models

import (
	"encoding/json"

	"github.com/go-sql-driver/mysql"
)

// Event is the definition of an event that we want to track
type Event struct {
	ID         string          `json:"id,omitempty"`
	Name       string          `json:"name,omitempty"`
	Properties json.RawMessage `json:"properties,omitempty"`
	CreatedOn  mysql.NullTime  `db:"created_on" json:"created_on,omitempty"`
	UpdatedOn  mysql.NullTime  `db:"updated_on" json:"updated_on,omitempty"`
	DeletedOn  mysql.NullTime  `db:"deleted_on" json:"deleted_on,omitempty"`
}
