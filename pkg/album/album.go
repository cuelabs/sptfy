package album

import (
	"net/url"
	"net/http"
    "github.com/cuelabs/sptfy/pkg/artist"
	"io/ioutil"
)

type SptfyAlbum struct {
    Name *string `json:"name"`
    Artists []*artist.SptfyArtist `json:"artists"`
	Id   *string `json:"id"`
	Uri  *string `json:"uri"`
	Href *url.URL `json:"href"`
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