package user

import (
	"net/url"

)

type SptfyUser struct {
	DisplayName *string                   `json:"display_name"`
	Id          *string                `json:"id"`
	Uri         *string                   `json:"uri"`
	Href        url.URL                   `json:"href"`
}

type SpotifyUserInfo struct {
	DisplayName string
	LoginEmail string
	NowPlaying string
}


//func (u *SptfyUser) Playlists(n int) ([]*playlist.SptfyPlaylist, error) {}
