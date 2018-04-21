package spotifyclient

import (
	"math/rand"
	"fmt"
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

