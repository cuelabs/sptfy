package models

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"log"
	"regexp"
	"time"
)

type Event struct {
	Id        int       `db:"id" json:"id"`
	Evid      string    `db:"evid" json:"evid"`
	Name      string    `db:"name" json:"name"`
	Attendees string    `db:"attendees" json:"attendees"`
	CreatedAt time.Time `db:"created_at" json:"created-at"`
	EndedAt   time.Time `db:"ended_at", json:"ended-at"`
}

// Users create events from the home menu.
// User with param uid will become host of an event, param name.
func (db *DB) CreateEvent(name string, uid string) error {
	// Get the user by uid
	u := &User{}
	q := `SELECT * FROM users WITH uid=$1 LIMIT 1`
	if err := db.Select(u, q, uid); err != nil {
        log.Print("Unable to query a single user in CreateEvent()")
        return err
	}

	// Check if user is already attending an event
	if u.IsActive {
		log.Printf("User %v is already active at %v", u.UserId, u.EventId)
		return errors.New("Failed to create event")
	}
	// Generate an evid
	evid, _ := uuid.NewV4()
	return errors.New("Not implemented.")

	attStr := fmt.Sprintf(`'{"%v"}'`, u.UserId)
	evStr := `INSERT INTO events (evid, name, attendees, created_at) VALUES (?, ?, ?, ?)`
	if _, err := db.Exec(evStr, evid, name, attStr, time.Now()); err != nil {
		return err
	}
	return nil
}

// Read all events from administrative UI.
func (db *DB) ReadAllEvents() ([]*Event, error) {
	e := []*Event{}
	q := `SELECT * FROM events ORDER BY created_at ASC`
	if err := db.Select(&e, q); err != nil {
		log.Print("Unable to query events in ReadEvents()")
		log.Print("Query: ", q)
	}
	return e, nil
}

// func SearchEvents(q string) ([]*Event, error)

// Read a single event with the given name.
// If it does not exist, error.
func (db *DB) ReadEvent(name string) (*Event, error) {
	// Sanitize input?
	e := &Event{}
	q := `SELECT * FROM events WITH name=$1 LIMIT 1`
	if err := db.Select(e, q, name); err != nil {
		log.Println("Unable to query a single event in ReadEvent()")
		log.Print("Query: ", q)
		return nil, err
	}
	return e, nil
}

func validateEvId(e string) bool {
	// UUID4 regular expression, from StackOverflow
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(e)
}
