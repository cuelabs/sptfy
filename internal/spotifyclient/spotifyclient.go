package spotifyclient

import (
	"math/rand"
	"fmt"
	"github.com/cuelabs/sptfy/pkg/track"
	"github.com/cuelabs/sptfy/pkg/album"
	"github.com/cuelabs/sptfy/pkg/artist"
	"github.com/cuelabs/sptfy/internal/environment"
	"github.com/cuelabs/sptfy/pkg/user"
)

type SpotifyApiOperations interface {
	RetrieveInfo(e *environment.Environment) (*user.SptfyUser, error)

	RetrieveAuth(e *environment.Environment) error

	PlaybackNext(e *environment.Environment) (*track.SptfyTrack, error)
	PlaybackPlay(e *environment.Environment) (*track.SptfyTrack, error)
	PlaybackPause(e *environment.Environment) (*track.SptfyTrack, error)

	SearchAlbum(query string, e *environment.Environment) ([]*album.SptfyAlbum, error)
	SearchArtist(query string, e *environment.Environment) ([]*artist.SptfyArtist, error)
	SearchTrack(query string, e *environment.Environment) ([]*track.SptfyTrack, error)
}

