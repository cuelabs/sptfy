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

type SptfyTag struct {
	Tag []byte
}
func (t *SptfyTag) New() error {
	tag := make([]byte, 6)
	if _, err := rand.Read(tag); err != nil {
		return fmt.Errorf("Unable to create random bytes in SptyTag.New(): %v", err)
	}
	t.Tag = tag
	return nil
}

