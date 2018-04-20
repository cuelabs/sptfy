package infoserver

import "github.com/cuelabs/sptfy/pkg/track"

type UserInfoResponse struct {
	Name string
	Email string
	CurrentTrack *track.SptfyTrack
}