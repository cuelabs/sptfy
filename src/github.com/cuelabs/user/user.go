package user

import "net/url"

// Basest type of song identifier: the track uri on Spotify
type SpotifyUserUri struct {
	Uri string `json:"uri"`
}

// Ensure correct number of digits and 'spotify' prefix
func (s *SpotifyUserUri) validate(r string) (bool, error) {
	// regexp to ensure a spotify Uri

	// IMPLEMENTATION

	return false, nil
}

type User struct {

}

// Spotify track: id, name, playback url
type SpotifyUser struct {
	Name url.URL `json:"playback_uri"`
	UserUri   *SpotifyUserUri
	UserUrl url.URL `json:"user-url"`
}