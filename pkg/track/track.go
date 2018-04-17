package track

import (
	"github.com/cuelabs/sptfy/pkg/album"
	"github.com/cuelabs/sptfy/pkg/artist"
	"io/ioutil"
	"net/http"
	"net/url"
	"errors"
	"fmt"
)

type SptfyTrack struct {
	PlaybackUrl url.URL               `json:"playback_url"`
	Name        *string               `json:"name"`
	Artists     []*artist.SptfyArtist `json:"artists"`
	Album       *album.SptfyAlbum     `json:"album"`
	IsPlayable  bool `json:"is_playable"`
	Id          *string               `json:"id"`
	Uri         *string               `json:"uri"`
	Href        *url.URL               `json:"href"`
}

// Display web API endpoint containing full entry for Spotify track
func (t *SptfyTrack) Details() (*[]byte, error) {
	resp, err := http.Get(t.Href.String())
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	c, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (t *SptfyTrack) Play() error {
	resp, err := http.Get(t.PlaybackUrl.String())
	c, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err == nil {
		// implement
		fmt.Println(c)
		return errors.New("implement")
	}
}
