package environment

import (
	"github.com/cuelabs/sptfy/internal/auth"
	"github.com/cuelabs/sptfy/pkg/album"
	"github.com/cuelabs/sptfy/pkg/artist"
	"github.com/cuelabs/sptfy/pkg/track"
	"github.com/cuelabs/sptfy/pkg/user"
	"log"
)

type SpotifyApiOperations interface {
	RetrieveInfo(e *Environment) (*user.SptfyUser, error)

	RetrieveAuth(e *Environment) error

	PlaybackNext(e *Environment) (*track.SptfyTrack, error)
	PlaybackPlay(e *Environment) (*track.SptfyTrack, error)
	PlaybackPause(e *Environment) (*track.SptfyTrack, error)

	SearchAlbum(query string, e *Environment) ([]*album.SptfyAlbum, error)
	SearchArtist(query string, e *Environment) ([]*artist.SptfyArtist, error)
	SearchTrack(query string, e *Environment) ([]*track.SptfyTrack, error)
}

type Envvars struct {
	Version    string
	ClientType string
}

type Environment struct {
	Auth   *auth.Authentication
	Vars   *Envvars
	Log    *log.Logger
	Client SpotifyApiOperations
}
