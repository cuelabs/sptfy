package spotifyclient

import (
	"net/url"
	"github.com/cuelabs/sptfy/internal/auth"
	"github.com/cuelabs/sptfy/internal/environment"
	"github.com/cuelabs/sptfy/pkg/track"
	"github.com/cuelabs/sptfy/pkg/album"
	"github.com/cuelabs/sptfy/pkg/artist"
	"github.com/cuelabs/sptfy/pkg/user"
	"net/http"
	"fmt"
	"context"
	"io/ioutil"
	"errors"
	"math/rand"
)

const (
	SPTFY_CACHE_PATH string = "~/.sptfy/token.json"
)

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

