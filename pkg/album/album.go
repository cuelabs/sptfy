package album

import (
	"github.com/cuelabs/sptfy/pkg/artist"
	"io/ioutil"
	"net/http"
	"net/url"
)

type SptfyAlbum struct {
	Name    *string               `json:"name"`
	Artists []*artist.SptfyArtist `json:"artists"`
	Id      *string               `json:"id"`
	Uri     *string               `json:"uri"`
	Href    *url.URL              `json:"href"`
}

type SpotifyApiAlbumResponse struct {
	Artists struct {
		Href `json:"href"`
		items []struct {

		}
	} `json:"artists"`

}

func (a *SptfyAlbum) Details() (*[]byte, error) {
	resp, err := http.Get(a.Href.String())
	if err != nil {
		return nil, err
	}
	c, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
