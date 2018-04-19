package environment

import (
	"github.com/cuelabs/sptfy/internal/auth"
	"github.com/cuelabs/sptfy/internal/spotifyclient"
	"log"
)

type Envvars struct {
	Version string
	ClientType string
}

type Environment struct {
	Auth   *auth.Authentication
	Vars   *Envvars
	Log    *log.Logger
	Client *spotifyclient.SpotifyApiOperations
}
