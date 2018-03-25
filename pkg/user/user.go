package user

import (
	"net/url"
	"github.com/cuelabs/sptfy/pkg/playlist"
)

type SptfyUser struct {
	DisplayName *string `json:"display_name"`
	Playlists []*playlist.SptfyPlaylist `json:"playlists"`
	Id *string `json:"id"`
	Uri *string `json:""`
	Href url.URL `json:"url"`
}

// Return a Spotify user's playlists
func (u *SptfyUser) GetPlaylists(n int) (*playlist.SptfyPlaylist, error) {
	//
	}
