package track

import (
	"encoding/json"
	"github.com/cuelabs/sptfy/album"
	"github.com/cuelabs/sptfy/artist"
	"io/ioutil"
	"net/http"
	"net/url"
)

type SptfyTrack struct {
	PlaybackUrl url.URL        `json:"playback_url"`
	Name        *string        `json:"name"`
	Artists     []*SptfyArtist `json:"artists"`
	Album       *SptfyAlbum    `json:"album"`
	Id          *string        `json:"id"`
	Href        url.URL        `json:"href"`
}

func (s *SptfyTrack) Details() (*[]byte, error) {
	resp, err := http.Get(s.Href)
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
