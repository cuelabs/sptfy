package environment

import (
	"github.com/cuelabs/cue/models"
)

// major DRY violation

type Environment struct {
	Db   models.Store
	Vars Envvars
	User models.User
}

type Envvars struct {
	Version string
	Dbconn  string
}
