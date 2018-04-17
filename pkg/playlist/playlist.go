package playlist

import (
	"net/url"
	"github.com/cuelabs/sptfy/pkg/track"
	"github.com/cuelabs/sptfy/pkg/user"

)

type SptfyPlaylist struct {
	Name   *string             `json:"name"`
	Owner  *user.SptfyUser     `json:"owner"`
	Tracks []*track.SptfyTrack `json:"tracks"`
	Id     *string             `json:"id"`
	Uri    *string             `json:"uri"`
	Href   *url.URL            `json:"href"`
}
