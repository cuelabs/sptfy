package user

import (
	"github.com/cuelabs/sptfy/pkg/playlist"
	"net/url"
)

type SptfyUser struct {
	DisplayName *string                   `json:"display_name"`
	Playlists   []*playlist.SptfyPlaylist `json:"playlists"`
	Id          *string                   `json:"id"`
	Uri         *string                   `json:""`
	Href        url.URL                   `json:"url"`
}

//func (u *SptfyUser) Playlists(n int) ([]*playlist.SptfyPlaylist, error) {}
