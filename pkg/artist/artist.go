package artist

import (
	"net/url"
	"net/http"
	"io/ioutil"
)

type SptfyArtist struct {
	Name *string  `json:"name"`
	Id   *string  `json:"id"`
	Uri  *string  `json:"uri"`
	Href *url.URL `json:"href"`
}

//
type SpotifyAPIArtistResponse struct {

}
