package user

import (
	"github.com/cuelabs/sptfy/pkg/playlist"
	"net/url"
)

type SptfyUser struct {
	DisplayName *string                   `json:"display_name"`
	Playlists   []*playlist.SptfyPlaylist `json:"playlists"`
	Id          *string                   `json:"id"`
	Uri         *string                   `json:"uri"`
	Href        url.URL                   `json:"href"`
}

type SpotifyUserInfo struct {
	DisplayName string
	LoginEmail string
	NowPlaying string
}


//func (u *SptfyUser) Playlists(n int) ([]*playlist.SptfyPlaylist, error) {}
