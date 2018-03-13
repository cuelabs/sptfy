package track

import (
	"net/url"
)

// Basest type of song identifier: the track uri on Spotify
type SpotifyUri struct {
	Uri string `json:"uri"`
}

// Ensure correct number of digits and 'spotify' prefix
func (s *SpotifyUri) validate(r string) (bool, error) {
	// regexp to ensure a spotify Uri

	// IMPLEMENTATION

	return false, nil
}

// A Spotify track resource identifier
type SpotifyTrackUri struct {
	TrackUri *SpotifyUri
}

// Confirm prefix 'spotify:track' and correct number of characters
func (s *SpotifyTrackUri) validate(r string) {

	// IMPLEMENT

}

// Spotify track: id, name, playback url
type SpotifyTrack struct {
	PlaybackUrl url.URL `json:"playback_uri"`
	TrackUri    *SpotifyTrackUri
	TrackName   string `json:"track_name"`
}