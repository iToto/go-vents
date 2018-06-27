package main

import (
	"encoding/json"

	"fmt"
	"log"
	"net/http"
	"os"

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
	router.HandleFunc("/events/{id}", GetEvent).Methods("GET")
	// router.HandleFunc("/foo/{id}", PostFoo).Methods("POST")
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello there and welcome to your service! 111")
}

func GetEvent(w http.ResponseWriter, r *http.Request) {
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

	jstring, err := json.Marshal(event)
	if err != nil {
		log.Printf(
			"error marshalling event to json with error: %s",
			err.Error(),
		)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(``))
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jstring)
}

func PostFoo(w http.ResponseWriter, r *http.Request) {
	// TODO
}
