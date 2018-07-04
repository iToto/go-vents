package models

import (
	"encoding/json"

	"github.com/lib/pq"
)

// SetEvent is the definition of an event that we want to track
type SetEvent struct {
	ID         string          `db:"id" json:"id,omitempty"`
	Name       string          `db:"name" json:"name,omitempty"`
	Properties json.RawMessage `db:"properties" json:"properties,omitempty"`
	CreatedOn  pq.NullTime     `db:"created_on" json:"created_on,omitempty"`
	UpdatedOn  pq.NullTime     `db:"updated_on" json:"updated_on,omitempty"`
	DeletedOn  pq.NullTime     `db:"deleted_on" json:"deleted_on,omitempty"`
}
