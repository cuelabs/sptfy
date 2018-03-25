package track

import (
	"github.com/cuelabs/sptfy/pkg/album"
	"github.com/cuelabs/sptfy/pkg/artist"
	"io/ioutil"
	"net/http"
	"net/url"
)

type SptfyTrack struct {
	PlaybackUrl url.URL               `json:"playback_url"`
	Name        *string               `json:"name"`
	Artists     []*artist.SptfyArtist `json:"artists"`
	Album       *album.SptfyAlbum     `json:"album"`
	Id          *string               `json:"id"`
	Uri         *string               `json:"uri"`
	Href        url.URL               `json:"href"`
}

func (t *SptfyTrack) Details() (*[]byte, error) {
	resp, err := http.Get(t.Href.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	c, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
