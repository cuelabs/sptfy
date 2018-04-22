package album

import (
	"net/url"
	"github.com/cuelabs/sptfy/internal/spotifyclient"
	"github.com/cuelabs/sptfy/pkg/artist"
)

type SptfyAlbum struct {
	Name    *string               `json:"name"`
	Artists []*artist.SptfyArtist `json:"artists"`
	Id      *string               `json:"id"`
	Uri     *string               `json:"uri"`
	Href    *url.URL              `json:"href"`
}

type SpotifyAPIAlbumResponse struct {
	Albums struct {
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
			Name string `json:"name"`
			Popularity int `json:"popularity"`
			Type string `json:"type"`
			Uri string `json:"uri"`
		} `json:"items"`
		Limit int `json:"limit"`
		// Next int `json:"next"` // I don't know the type
		Offset int `json:"offset"`
		// Previous int `json:"previous"` // Same here
		Total int `json:"total"`
	} `json:"albums"`
}

type SptfyCLIAlbumDisplayResponse struct {
	ResponseHead string
    MessageFormat string
    Results []*SptfyCLIAlbumDisplayItem
}

type SptfyCLIAlbumDisplayItem struct {
	Tag *spotifyclient.SptfyTag
	AlbumName *string
	ArtistName *string
}
