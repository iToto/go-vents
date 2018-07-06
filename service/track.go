package service

import (
	"fmt"
	"time"

	"github.com/guregu/null"
	"github.com/iToto/go-vents/models"
	"github.com/jmoiron/sqlx"
)

// TrackService is the service that will capture triggered events
type TrackService struct {
	eventService *EventService
	db           *sqlx.DB
}

// NewTrackService ...
func NewTrackService(es *EventService, db *sqlx.DB) *TrackService {
	ts := &TrackService{
		eventService: es,
		db:           db,
	}

	return ts
}

// TrackEvent will capture a tracked event after validating it first
func (ts TrackService) TrackEvent(event models.TrackEvent) error {
	// Validate if event exists
	err := ts.validateEvent(event)

	if err != nil {
		return err
	}

	event.CreatedOn = null.NewTime(time.Now(), true)
	// Event exists, lets capture tracked event
	query := `INSERT INTO public.tracked_events 
	(name, properties, created_on, tracked_on) 
	VALUES (:name, :properties, :created_on, :tracked_on)`

	_, err = ts.db.NamedQuery(query, &event)

	if err != nil {
		// TODO: Add to retry queue
		err = fmt.Errorf(
			"Create track event with error %s",
			err.Error(),
		)
		return err
	}

	return nil
}

// validateEvent will ensure that the event name is on our whitelist
func (ts TrackService) validateEvent(event models.TrackEvent) error {
	// Check if the event name is one that we care about
	_, err := ts.eventService.GetByName(event.Name)

	if err != nil {
		// Not found, ignore
		fmt.Errorf(
			"could not find set event with name %s, with error %s",
			event.Name,
			err.Error(),
		)
		return err
	}

	return nil
}
