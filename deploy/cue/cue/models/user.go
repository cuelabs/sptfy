package models

import (
	"errors"
	"time"
	"log"
	"github.com/satori/go.uuid"
	"regexp"
	"strings"
)

type User struct {
	Id          int       `db:"id" json:"id"`
	UserId      string    `db:"uid" json:"user-id"`
	SpotifyURI  string    `db:"suri" json:"spotify-uri"`
	DisplayName string    `db:"display_name" json:"display-name"`
	IsActive    bool      `db:"is_active" json:"is-active"`
	EventId     string    `db:"evid" json:"event-id"`
	IsHost      string    `db:"is_host" json:"is-host"`
	JoinedAt    time.Time `db:"joined_at" json:"joined-at"`
}

// Creates a user upon initial log-in irregardless of joining an event
func (db *DB) CreateUser(suri string, name string) error {
	// Check if user already exists with that suri
	u := &User{}
	q := `SELECT * FROM users WITH suri=$1 LIMIT 1`
	if err := db.Select(u, q, suri); err != nil {
		if strings.Contains(err.Error(), "Did not find a result") {
			// No user exists with this suri
			// Continue to create the user
			// Generate a uid
			uid, _ := uuid.NewV4()
			uStr := `INSERT INTO users (uid, suri, display_name, is_active, joined_at) VALUES (?, ?, ?, ?, ?)`
			if _, err := db.Exec(uStr, uid, suri, name, false, time.Now()); err != nil {
				return err
			}
		}
		log.Println("Unable to query users in CreateUser()")
		log.Print("Query: ", q)
		log.Print("Suri: ", u.SpotifyURI)
		return err
	}
	log.Println("User already exists with provided Spotify URI in CreateUser()")
	log.Print("Suri: ", u.SpotifyURI)
	return errors.New("USERALREADYEXISTSERROR")
}

func (db *DB) ReadUser (suri string) (*User, error) {
	if ok := validateSuri(suri); !ok {
         log.Println("Invalid suri entered on ReadUser()")
         log.Print("Suri: ", suri)
         return nil, errors.New("INVALIDSURIERROR")
	}
	u := &User{}
	q := `SELECT * FROM users WITH suri=$ LIMIT 1`
	if err := db.Select(u, q, suri); err != nil {
		log.Println("Unable to query a single user in ReadUser()")
		log.Print("Query: ", q)
		log.Print("Suri: ", suri)
		return nil, errors.New("FAILEDUSERQUERYERROR")
	}
	return u, nil
}

func (db *DB) ReadAttendees(evid string) ([]*User, error) {
	// Validate input string is a UUID4
	if ok := validateEvId(evid); !ok {
		log.Println("Invalid event id entered on ReadUsersFromEvent()")
		return nil, errors.New("")
	}

	// Get the attendees from the event input string
	var ev string
	if err := db.Select(ev, "SELECT (attendees) from events WITH evid=$1 LIMIT 1", evid); err != nil {
		log.Println("Unable to query events in ReadUsersFromEvent()")
		return nil, err
	}

	idRegexp := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
    attendeesList := idRegexp.FindAllString(ev, -1)
    attendees := []*User{}
    for _, uid := range attendeesList {
    	// Get the user
    	user := &User{}
    	if err := db.Select(user, "SELECT * FROM users WITH uid=$1 LIMIT 1", uid); err != nil {
    		log.Println("Unable to query a single user in ReadUsersFromEvent")
    		return nil, err
		}
    	// Add user to output
    	attendees = append(attendees, user)
	}
	return attendees, nil
}

// From the Spotify Web API
// "The base-62 identifier that you can find at the end of the Spotify URI (see above) for an artist, track, album, playlist, etc.
// Unlike a Spotify URI, a Spotify ID does not clearly identify the type of resource; that information is provided elsewhere in the call."
//
// We store the Spotify ID as 'suri', and prepend the resource identifier.
func validateSuri(s string) bool {
	return true // how to implement...
}