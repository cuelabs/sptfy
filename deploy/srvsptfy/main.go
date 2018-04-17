package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"context"
)

type Global struct {
	db Store
}

type TrackSearchServer struct{}
type ArtistSearchServer struct{}
type AlbumSearchServer struct{}

func (s *TrackSearchServer) SearchTrack(ctx context.Context, req pb.TrackSearchRequest) (*pb.TrackSearchResponse) {
	return &pb.TrackSearchResponse{nil}
	}

func newRouter(g *Global) *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.Handle(pb.TrackSearchServicePathPrefix, searchTrackHandler(g))
	r.Handle(pb.ArtistSearchServicePathPrefix, searchArtistHandler(g))
	r.Handle(pb.AlbumSearchServicePathPrefix, searchAlbumHandler(g))
	return r
}


func main() {
   db, err := NewDB("postgres://postgres:sptfy-dev-password/sptfy")
   if err != nil {
   	log.Print("Could not create a new database in main(). Exiting")
   	log.Panic(err)
   }

   globals := &Global{db} // implement the methods in respective data models

   r := newRouter(globals)
	http.ListenAndServe("10102", r)
}



func searchArtistHandler(g *Global) http.Handler {
	svc :=  pb.NewAlbumSearchServiceServer()

	res := pb.NewArtistSearchServiceServer(svc, nil)
}

func searchAlbumHandler(g *Global) http.Handler {
	return pb.NewAlbumSearchServiceServer(&AlbumSearchServer{}, nil)
}
