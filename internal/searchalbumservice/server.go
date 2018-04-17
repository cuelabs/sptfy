package main

import (
	pb "github.com/cuelabs/sptfy/rpc/sptfyapi"
	"context"
	"net/url"
	"fmt"
)

type AlbumSearchServer struct{}

func (s *AlbumSearchServer) SearchAlbum(ctx context.Context, req *pb.AlbumSearchRequest) (*pb.AlbumSearchResponse, error) {
	q := req.Query
	// make request to spotify API with the currently logged in user
	p := fmt.Sprintf("/v1/search?q=name:%v&type=album", q)
	u := url.URL{
		Scheme: "https",
		Host: "api.spotify.com",
		Path: p,
	}
	return nil, nil
}