package models

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type Store interface {
	CreateEvent(name string, uid string) error
	// Administrative: read all events
	ReadAllEvents() ([]*Event, error)
	// Search through events
	// SearchEvents(q string) ([]*Event, error)
	// Read a single event
	ReadEvent(name string) (*Event, error)

	CreateUser(suri string, name string) error
	ReadUser(suri string) (*User, error)
	ReadAttendees(evid string) ([]*User, error)
}

type DB struct {
	*sqlx.DB
}

func NewDB(dataSourceName string) (*DB, error) {
	db, err := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		log.Println("Data source: ", dataSourceName)
		log.Println("Could not create a new database in NewDB()")
		return nil, err
	}
	if err := db.Ping(); err != nil {
		log.Println("Could not ping new database in NewDB()")
		return nil, err
	}
	return &DB{db}, nil
}
