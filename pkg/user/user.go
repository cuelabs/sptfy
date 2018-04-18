package user

import (
	"net/url"
)

type SptfyUser struct {
	DisplayName *string `json:"display_name"`
	Id          *string `json:"id"`
	Uri         *string `json:"uri"`
	Href        url.URL `json:"href"`
}


// TODO look at https://api.cue.zone/v1/me
type SpotifyUserInfoResponse struct {
	DisplayName string `json:"display_name"`
	LoginEmail  string `json:""`
	NowPlaying  string
}

//func (u *SptfyUser) Playlists(n int) ([]*playlist.SptfyPlaylist, error) {}
