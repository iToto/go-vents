package main

import (
	"encoding/json"

	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/iToto/go-vents/models"
	"github.com/iToto/go-vents/service"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
)

var db *sqlx.DB

type Foo struct {
	Bar string
}

func DBConnection() *sqlx.DB {
	dbURL := os.Getenv("DATABASE_URL")

	if dbURL == "" {
		log.Fatal("$DATABASE_URL environment variable must be set")
	} else {
		fmt.Println("Connected to database: " + dbURL)
	}

	db, err := sqlx.Open("postgres", dbURL)

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT environment variable must be set")
	} else {
		fmt.Println("Running on port: " + port)
	}

	db = DBConnection()
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", Index)
	router.HandleFunc("/events", CreateEvent).Methods("POST")
	router.HandleFunc("/events", ListEvents).Methods("GET")
	router.HandleFunc("/events/{id}", GetEvent).Methods("GET")
	router.HandleFunc("/events", UpdateEvent).Methods("PUT")
	router.HandleFunc("/events/{id}", DeleteEvent).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":"+port, router))
}

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello there and welcome to your service! 111")
}

// GetEvent ...
func GetEvent(w http.ResponseWriter, r *http.Request) {
	log.Println("Get Event Requested")
	vars := mux.Vars(r)
	id := vars["id"]
	eventService := service.NewEventService(db)

	event, err := eventService.Get(id)

	if err != nil {
		log.Printf(
			"error getting event with id: %s with error: %s",
			id,
			err.Error(),
		)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(``))
	}

	sendJSON(event, w)
}

// CreateEvent ...
func CreateEvent(w http.ResponseWriter, r *http.Request) {
	log.Println("Create Event Requested")
	var event models.SetEvent

	eventService := service.NewEventService(db)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&event)

	if err != nil {
		log.Printf(
			"error parsing payload with error: %s",
			err.Error(),
		)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(``))
	}

	newEvent, err := eventService.Create(event)

	if err != nil {
		log.Printf(
			"error creating event with error: %s",
			err.Error(),
		)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(``))
	}

	sendJSON(newEvent, w)
}

// ListEvents ...
func ListEvents(w http.ResponseWriter, r *http.Request) {
	log.Println("List Events Requested")
	eventService := service.NewEventService(db)
	events, err := eventService.List()

	if err != nil {
		log.Printf(
			"error getting events with error: %s",
			err.Error(),
		)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(``))
	}

	sendJSON(events, w)
}

// UpdateEvent ...
func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	log.Println("Update Event Requested")
	var event models.SetEvent

	eventService := service.NewEventService(db)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&event)

	if err != nil {
		log.Printf(
			"error parsing payload with error: %s",
			err.Error(),
		)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(``))
	}

	updatedEvent, err := eventService.Update(event)

	if err != nil {
		log.Printf(
			"error updating event with error: %s",
			err.Error(),
		)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(``))
	}

	sendJSON(updatedEvent, w)
}

// DeleteEvent ...
func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete Event Requested")
	vars := mux.Vars(r)
	id := vars["id"]
	eventService := service.NewEventService(db)
	event, err := eventService.Get(id)

	if err != nil {
		log.Printf(
			"error getting event with id: %s with error: %s",
			id,
			err.Error(),
		)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(``))
	}
	deletedEvent, err := eventService.Delete(*event)

	if err != nil {
		log.Printf(
			"error deleting event with id: %s with error: %s",
			id,
			err.Error(),
		)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(``))
	}

	sendJSON(deletedEvent, w)
}

func sendJSON(payload interface{}, w http.ResponseWriter) {
	jstring, err := json.Marshal(payload)
	if err != nil {
		log.Printf(
			"error marshalling payload to json with error: %s",
			err.Error(),
		)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(``))
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jstring)
}
