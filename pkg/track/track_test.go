package track

import (
      "testing"
     "net/url"
)

func TestSptfyTrack_Details( t *testing.T) {
    trk := SptfyTrack{PlaybackUrl: url.URL{
        Scheme: "https",
        Host: "api.spotify.com",
        Path: "/v1/tracks/",
    }}
    u0, err := url.Parse("")

}

func TestSptfyTrack_Play(t *testing.T) {

}
