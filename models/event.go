package models

import (
	"encoding/json"

	"github.com/guregu/null"
)

// SetEvent is the definition of an event that we want to track
type SetEvent struct {
	ID         string          `db:"id" json:"id,omitempty"`
	Name       string          `db:"name" json:"name,omitempty"`
	Properties json.RawMessage `db:"properties" json:"properties,omitempty"`
	CreatedOn  null.Time       `db:"created_on" json:"created_on,omitempty"`
	UpdatedOn  null.Time       `db:"updated_on" json:"updated_on,omitempty"`
	DeletedOn  null.Time       `db:"deleted_on" json:"deleted_on,omitempty"`
}
