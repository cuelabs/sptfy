package environment

type Envvars struct {
	Version string
}

type Environment struct {
	auth    auth.Authentication
	envvars Envvars
	log     *log.Logger
	client  *SpotifyApiOperations
}
