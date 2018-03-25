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

func (a *SptfyArtist) Details() (*[]byte, error) {
	resp, err := http.Get(a.Href.String())
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