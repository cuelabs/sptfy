package spotifyclient

type SpotifyHttpClient struct {}

type SptfyRpcClient struct {
	SptfyHost *url.URL
}

type SpotifyApiOperations interface {
	RetrieveInfo(e *environment.Environment) (*user.SptfyUser, error)

	RetrieveAuth() (*auth.Authentication, error)

	PlaybackNext(e *environment.Environment) (*track.SptfyTrack, error)
	PlaybackPlay(a *auth.Authentication) (*track.SptfyTrack, error)
	PlaybackPause(a *auth.Authentication) (*track.SptfyTrack, error)

	// Do I need to pass in auth.Authentication for this?
	SearchAlbum(query string) (*album.SptfyAlbum, error)
	SearchArtist(query string) (*artist.SptfyArtist, error)
	SearchTrack(query string) (*track.SptfyTrack, error)
}

func (s *SpotifyHttpClient) RetrieveInfo(e *Environment) (*user.SptfyUser, error) {
	// Create a client
	token, err := e.auth.Token()
	 if err != nil {
	 	e.log.Println("Unable to retrieve token in RetrieveInfo()")
	 	return nil, err
	 }
	 conf, err := e.auth.Config()
	 if err != nil {
	 	e.log.Println("Unable to create retrieve OAuth2 config in RetrieveInfo()")
	 	return nil, err
	 }
	 client := conf.Client(context.Background(), token)
	u := &url.URL{
		Scheme: "https",
		Host: "api.spotify.com",
		Path: "/v1/me",
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		e.log.Println("Unable to create a new request in RetrieveInfo()")
		return nil, err
	}
	 resp, err := client.Do(req)
	 if err != nil {
	 	e.log.Println("Request failed in RetrieveInfo()")
	 }
	 // marshall data response to
}

func (s *SpotifyHttpClient) RetrieveAuth() (*auth.Authentication, error) {
	return nil, errors.New("Not implemented")
}

// These handlers control playback with a Spotify
func (s *SpotifyHttpClient) PlaybackNext(a *auth.Authentication) (*track.SptfyTrack, error) {
	return nil, errors.New("Not implemented")
}

func (s *SpotifyHttpClient) PlaybackPlay() (*track.SptfyTrack, error) {
    return nil, errors.New("Not implemented")
}

func (s *SpotifyHttpClient) PlaybackPause(e *environment.Environment) (*track.SptfyTrack, error) {
    return nil, errors.New("Not implemented")
}

// These handlers search the Spotify API with(out) authentication
func (s *SpotifyHttpClient) SearchAlbum(query string) (*album.SptfyAlbum, error) {
	// search path
	sp := fmt.Sprintf("/v1/search?q=%v&type=album", query)
	// search url
	su := &url.URL{
		Scheme: "https",
		Host: "api.spotify.com",
		Path: sp}
		resp, err := http.Get(su.String())
		if err != nil {
			return nil, err
			}
			defer resp.Body.Close()
			// no auth?
			//
    b, err := ioutil.ReadAll(resp.Body)
    fmt.Println(b)
    fmt.Println("Got to the enc of SearchAlbum()")
    return nil, nil
}

func (s *SpotifyHttpClient) SearchArtist(query string) (*artist.SptfyArtist, error) {
	// sanitize query string; replace spaces with '%20'
	sp := fmt.Sprintf("/v1/search?q=%v")
}
