package app

import (
	//"github.com/cuelabs/sptfy/pkg/artist"
	//"github.com/cuelabs/sptfy/pkg/album"
	//"github.com/cuelabs/sptfy/pkg/track"
	//"github.com/cuelabs/cue/pkg/cue"
	"github.com/cuelabs/cue/environment"
	"net/http"
	"encoding/json"
)

type CueApplicationEndpoints interface {
	/*
		SearchSpotifyArtists(q string) ([]*artist.SptfyArtist, error)
		SearchSpotifyAlbums(q string) ([]*album.SptfyAlbum, error)
		SearchSpotifyTracks(q string) ([]*track.SptfyTrack, error)
	*/

	CreateEvent(env *environment.Environment) http.Handler
	AddUserToEvent(env *environment.Environment) http.Handler

	//AddTrackToCue(track *track.SptfyTrack, cue *cue.Cue) error

	//Play
	//Next
	//Pause

	//EndEvent(evid string) error
}

type CueApplication struct{}

func NewApp() *CueApplication {
	// some checks...
	return &CueApplication{}
}

type cueResp struct {
	code int    `json:"code"`
	msg  string `json:"message,omitempty"`
}

type cueApplicationResponse struct {
	status string `json:"status""`
	resp *cueResp `json:"resp"`
}

func (app *CueApplication) CreateEvent(env *environment.Environment) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			error := &cueResp{125, "HTTP method not allowed"}
			json.NewEncoder(w).Encode(cueApplicationResponse{"failure", error})
			return
		}
		// Parse the form
		r.ParseForm()
		name := r.Form.Get("name")
		uid := env.User.UserId
		if err := env.Db.CreateEvent(name, uid); err != nil {
			error := &cueResp{155, err.Error()}
			json.NewEncoder(w).Encode(cueApplicationResponse{"failure", error})
			return
		}
		// Send success
		succ := &cueResp{0, "Created an event"}
		json.NewEncoder(w).Encode(cueApplicationResponse{"success", succ})
		return
	})
}

func (app *CueApplication) AddUserToEvent(env *environment.Environment) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
