package playlist

import "net/url"

type Playlist struct {
	Name string `json:"name"`
	Url url.URL `json:"url"`
}