package album

import (
	"net/url"
    "github.com/cuelabs/sptfy/artist"
	)

type SptfyAlbum struct {
    Name *string `json:"name"`
    Artists []*artist.SptfyArtist `json:"artists"`
	Id   *string `json:"id"`
	Uri  *string `json:"uri"`
}

func (a *SptfyAlbum) Details() (*[]byte, error) {
	resp, err := http.Get()
}
