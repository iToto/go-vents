package models

import "time"

// Event is the definition of an event that we want to track
type Event struct {
	ID         string                 `json:"id,omitempty"`
	Name       string                 `json:"name,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty"`
	CreatedOn  time.Time              `json:"created_on,omitempty"`
	UpdatedOn  time.Time              `json:"updated_on,omitempty"`
	DeletedOn  time.Time              `json:"deleted_on,omitempty"`
}
