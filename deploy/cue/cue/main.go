package main

import (
	"encoding/json"
	"github.com/GeertJohan/go.rice"
	"github.com/llater/cue/app"
	"github.com/llater/cue/environment"
	"github.com/llater/cue/models"
	"github.com/cuelabs/sptfy/pkg/user"
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var env environment.Environment

func init() {
	log.SetOutput(os.Stdout)
	// Process environment variables
	log.Print("Processing environment variables")
	if err := envconfig.Process("cue", &env.Vars); err != nil {
		log.Print("Could not process environment variables. Exiting")
		log.Panic(err)
	}
	if env.Vars.Dbconn == "" || env.Vars.Version == "" {
		log.Print("Required environment variables not set. Exiting")
		os.Exit(1)
	}
	log.Print("Processed environment variables")
	log.Printf("Cue will server version %v on database at %v", env.Vars.Version, env.Vars.Dbconn)
}

func main() {
	// Add the database
	db, err := models.NewDB(env.Vars.Dbconn)
	env.Db = db
	if err != nil {
		log.Println("Could not create a new database in main(). Exiting")
		log.Panic(err)
	}
	if err := db.Ping(); err != nil {
		log.Println("Could not ping a new database in main(). Exiting")
		log.Panic(err)
	}

	// Get user from Spotify directly
	var userResp user.SpotifyAPIUserResponse
	resp, err := http.Get("https://api.spotify.com/v1/me")
	if err != nil {
		log.Print("Failed to retrieve a response from the Spotify API in main()")
		log.Print(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print("Failed to read the Spotify API response body in main()")
		log.Print(err)
	}
	if err := json.Unmarshal(body, &userResp); err != nil {
		log.Print("Failed to unmarshal Spotify API response to JSON in main()")
		log.Print(err)
	}

	// Look up the user, or create one
	id := userResp.Id
	user, err := env.Db.ReadUser(id)
	if err != nil {
		// If the user does not exist, create a new one
		if strings.Contains(err.Error(), "FAILEDUSERQUERYERROR") {
			log.Printf("User with id %v does not exist. Creating a new user.")
			if err := env.Db.CreateUser(id, userResp.DisplayName); err != nil {
				log.Print("Unable to create a user in main()")
				log.Print(err)
			}
		} else {
			log.Print("Unable to read a user in main()")
			log.Print(err)
		}
	}
	env.User = *user

        var cueApp app.CueApplicationEndpoints

	cueApp = app.NewApp()
	if err != nil {
		log.Print("Unable to create application interface. Exiting")
		log.Panic(err)
	}

	r := mux.NewRouter().StrictSlash(true)
	r.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("static").HTTPBox()))
	r.Handle("/create-event", app.CreateEvent(&env))
	r.Handle("/add-user-to-event", app.AddUserToEvent(&env))
	log.Fatal(http.ListenAndServe(":10000", r))

}
