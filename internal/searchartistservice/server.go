package main


import (
	"context"
	pb "github.com/cuelabs/sptfy/rpc/sptfyapi"
	"fmt"
	"net/url"
	"net/http"
	"log"
	"github.com/pkg/errors"
)


type ArtistResponse struct {
    Href string `json:"href"`
    Id string `json:"id"`
    Name string `json:"name"`
    Uri string `json:"uri"`
}

type ArtistSearchServer struct{}

func (a *ArtistSearchServer) SearchArtist(ctx context.Context, req *pb.ArtistSearchRequest) (*pb.ArtistSearchResponse, error) {
	q := req.Query
	// convert spaces to encoding wiith %20 ?
	p := fmt.Sprintf("/v1/search?q:name=%v&type=artist", q)
	u := url.URL{
		Scheme: "https",
		Host: "api.spotify.com",
		Path: p,
	}

	type SearchArtistResponse struct {
		Artists []*ArtistResponse `json:"items"`
		Limit int `json:"limit"`
		Offset int `json:"offset"`
		Total int `json:"total"`
	}



	resp, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
    	return nil, err
	}
	log.Print(resp)

	return &pb.ArtistSearchResponse{[]*pb.ArtistResult{}}, errors.New("Not implemented")
}
