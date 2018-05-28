package track

import (
	"github.com/cuelabs/sptfy/pkg/album"
	"github.com/cuelabs/sptfy/pkg/artist"
	"net/url"
)

type SptfyTrack struct {
	PlaybackUrl url.URL               `json:"playback_url"`
	Name        *string               `json:"name"`
	Artists     []*artist.SptfyArtist `json:"artists"`
	Album       *album.SptfyAlbum     `json:"album"`
	IsPlayable  bool                  `json:"is_playable"`
	Id          *string               `json:"id"`
	Uri         *string               `json:"uri"`
	Href        *url.URL              `json:"href"`
}

type SpotifyAPITrackSearchResponse struct {
	Tracks struct {
		Href  string `json:"href"`
		Items []struct {
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Genres []struct {} `json:"genres"` // need a working example of this
			Href string `json:"href"`
			Id string `json:"id"`
			Images []struct {
				Height int `json:"height"`
				Url string `json:"url"`
				Width string `json:"width"`
			} `json:"images"`
		} `json:"items"`
		Limit int `json:"limit"`
		// Next int `json:"next"` // I don't know the type
		Offset int `json:"offset"`
		// Previous int `json:"previous"` // Same here
		Total int `json:"total"`
	} `json:"tracks"`
}

type SptfyCLITrackDisplayResponse struct {
	ResponseHead string
	MessageFormat string
	Results []*SptfyCLITrackDisplayItem
}

type SptfyCLITrackDisplayItem struct {
	TrackName *string
	Artists []*string
	AlbumName *string
	//Duration *time.Duration add in v0.3+
}
