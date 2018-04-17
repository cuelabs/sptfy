package main

import (
	pb "github.com/cuelabs/sptfy/rpc/sptfyapi"
	"context"

)

type TrackSearchServer struct{}

func (t *TrackSearchServer) SearchTrack(ctx context.Context, req *pb.TrackSearchRequest) (*pb.TrackSearchResponse, error) {
	s0 := &pb.TrackResult{Suri: "a suri", Name: "track name"}
	s1 := &pb.TrackResult{Suri: "another suri", Name: "my fave", Artist: "my fave"}
	return &pb.TrackSearchResponse{[]*pb.TrackResult{s0, s1}}, nil
}

