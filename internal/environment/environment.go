package environment

import (
	"github.com/cuelabs/sptfy/internal/auth"
	"log"
	"github.com/cuelabs/sptfy/internal/spotifyclient"
)

type SptfySettings struct {
	Set map[string]interface{}
}

// Return false if already exists
func (s *SptfySettings) SetSetting(k string, v interface{}) bool {
	if s.Set[k] == nil {
        s.Set[k] = v
        return true
	}
	return false
}

func (s *SptfySettings) GetSetting(k string) (interface{}, bool) {
    if s.Set[k] == nil {
    	return nil, false
	}
    return s.Set[k], true
}

type Envvars struct {}

type Environment struct {
	Auth   *auth.Authentication
	Vars   *Envvars
	Log    *log.Logger
	Client spotifyclient.SpotifyApiOperations
	Settings *SptfySettings
}
