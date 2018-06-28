package service

import (
	"fmt"

	"github.com/iToto/go-vents/models"
	"github.com/jmoiron/sqlx"
	"github.com/satori/go.uuid"
)

// EventService is the service that will manage the events SoR
type EventService struct {
	db *sqlx.DB
}

// NewEventService ...
func NewEventService(db *sqlx.DB) *EventService {
	es := &EventService{
		db: db,
	}

	return es
}

// Get will retrieve an event by its ID
func (es EventService) Get(id string) (*models.SetEvent, error) {
	var event models.SetEvent
	query := "SELECT * FROM events WHERE id = $1"
	err := es.db.Get(&event, query, id)

	if err != nil {
		err = fmt.Errorf(
			"could not select event by id: %s with error %s",
			id,
			err.Error(),
		)
		return nil, err
	}

	return &event, nil
}

// List will list all events that exist in the SoR
func (es EventService) List() ([]models.SetEvent, error) {
	var events []models.SetEvent
	query := "SELECT * FROM events"
	rows, err := es.db.Query(query)
	if err != nil {
		err = fmt.Errorf(
			"could not select events with error %s",
			err.Error(),
		)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var event models.SetEvent
		err := rows.Scan(&event)
		if err != nil {
			err = fmt.Errorf(
				"could not parse results with error %s",
				err.Error(),
			)
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

// Create will create a new event in the SoR
func (es EventService) Create(event models.SetEvent) (*models.SetEvent, error) {
	// Populate UUID if not already set
	if event.ID == "" {
		event.ID = uuid.Must(uuid.NewV4()).String()
	}
	query := "INSERT INTO public.events (id, name, properties, created_on) VALUES (:id, :name, :properties, :createdon)"
	_, err := es.db.NamedExec(query, &event)

	if err != nil {
		err = fmt.Errorf(
			"Create event with error %s",
			err.Error(),
		)
		return nil, err
	}

	return &event, nil
}

// Update will update an existing event in the SoR
func (es EventService) Update(event models.SetEvent) (*models.SetEvent, error) {
	return nil, nil
}

// Delete will delete an existing event by its ID
func (es EventService) Delete(id string) error {
	return nil
}
