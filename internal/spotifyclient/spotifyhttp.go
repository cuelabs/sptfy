package spotifyclient

import (
	"github.com/cuelabs/sptfy/internal/environment"
	"github.com/cuelabs/sptfy/pkg/album"
	"github.com/cuelabs/sptfy/internal/auth"
	"fmt"
	"net/url"
	"net/http"
	"io/ioutil"
	"github.com/cuelabs/sptfy/pkg/artist"
	"github.com/cuelabs/sptfy/pkg/track"
	"github.com/cuelabs/sptfy/pkg/user"
	"context"
	"errors"
	"encoding/json"
)

type SpotifyHttpClient struct {}

func (s *SpotifyHttpClient) RetrieveInfo(e *environment.Environment) (*user.SptfyUser, error) {
	// Create a client
	token, err := e.Auth.GetToken()
	if err != nil {
		e.Log.Println("Unable to retrieve token in RetrieveInfo()")
		return nil, err
	}
	conf, err := e.Auth.Config()
	if err != nil {
		e.Log.Println("Unable to create retrieve OAuth2 config in RetrieveInfo()")
		return nil, err
	}
	client := conf.Client(context.Background(), token)
	urlChk := &url.URL{
		Scheme: "https",
		Host: "api.spotify.com",
		Path: "/v1/me",
	}
	req, err := http.NewRequest("GET", urlChk.String(), nil)
	if err != nil {
		e.Log.Println("Unable to create a new request in RetrieveInfo()")
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		e.Log.Println("Request failed in RetrieveInfo()")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		e.Log.Println("Unable to read response body in RetrieveInfo()")
		return nil, err
	}
	// Marsall Spotify API response usable format
	var u user.SpotifyAPIUserResponse
	if err := json.Unmarshal(body, &u); err != nil {
		e.Log.Println("Unable to unmarashal response body to user.SpotifyAPIUserResponse type in RetrieveInfo()")
       return nil, err
	}
    href, err := url.Parse(u.Href)
    if err != nil {
    	e.Log.Println("Unable to parse 'href' in response to a valid URL in RetrieveInfo()")
	return nil, err
    	}
	return &user.SptfyUser{
		DisplayName: &u.DisplayName,
		Email: &u.Email,
		Id: &u.Id,
		Uri: &u.Uri,
		Href: href,
	}, nil
}

func (s *SpotifyHttpClient) RetrieveAuth(e *environment.Environment) error {
	if e.Auth == nil {
		e.Log.Println("Auth not found. Beginning authentication.")
		a := &auth.Authentication{"~/.sptfy/token.json", nil}
		fmt.Print("Please authenticate")
		a.Authenticate()
		e.Log.Println("Authentication complete.")
		fmt.Println("Thank you for authenticating.")
		e.Auth = a
	}
	// Ensure access token exists and is active, nil}
	fmt.Print("Please authenticate")
	/*
u, err := e.Client
	if err != nil {
         e.Log.Println(("Unable to retrieve a user in RetrieveAuth()"))
         return err
	}
	*/
	// Display some authorization content
	e.Log.Println("RetrieveAuth() successful. Displaying auth info for user ", "fake ID") //u.Id
	fmt.Println("Display Name: ", "Placehold") // u.DisplayName
	fmt.Println("Email Address: ", "Placeholder 2") // u.Email
	fmt.Print("\nSuccessfully authoirzed!\n\n")
	return nil
}

// These handlers control playback with a Spotify
func (s *SpotifyHttpClient) PlaybackNext(e *environment.Environment) (*track.SptfyTrack, error) {
	return nil, errors.New("Not implemented")
}

func (s *SpotifyHttpClient) PlaybackPlay(e * environment.Environment) (*track.SptfyTrack, error) {
	return nil, errors.New("Not implemented")
}

func (s *SpotifyHttpClient) PlaybackPause(e *environment.Environment) (*track.SptfyTrack, error) {
	return nil, errors.New("Not implemented")
}

// Return a list of sptfy albums
func (s *SpotifyHttpClient) SearchAlbum(query string, e *environment.Environment) ([]*album.SptfyAlbum, error) {
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

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		e.Log.Println("Unable to read response body in SearchAlbum()")
		return nil, err
	}
	var res album.SpotifyAPIAlbumResponse
	if err := json.Unmarshal(b, &res); err != nil{
		e.Log.Println("Unable to unmarshal album search result in SearchAlbum()")
		return nil, err
	}
	var albums []*album.SptfyAlbum
	for _, alb := range(res.Albums.Items) {
		u, err := url.Parse(alb.Href)
		if err != nil {
			e.Log.Println("Unable to parse href URL in SearchAlbum()")
			e.Log.Print(err.Error())
			return nil, err
		}
		n := &album.SptfyAlbum{
			&alb.Name,
			&alb.Id,
			&alb.Uri,
			u,
		}
		albums = append(albums, n)
	}
	return albums, nil
}

func (s *SpotifyHttpClient) SearchArtist(query string, e *environment.Environment) ([]*artist.SptfyArtist, error) {
	// sanitize query string; replace spaces with '%20'
	sp := fmt.Sprintf("/v1/search?q=%v&type=artist", query)
	// search url
	u := &url.URL{
		Scheme: "https",
		Host: "api.spotify.com",
		Path: sp,
	}
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	fmt.Println(b)
	fmt.Println("Got to end of SearchArtistr()")
	return nil, errors.New("Not implemented")
}

func (s *SpotifyHttpClient) SearchTrack(query string, e *environment.Environment) ([]*track.SptfyTrack, error) {
	return nil, errors.New("Not implemented.")
}
