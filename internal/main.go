package main

import (
	"net/http"

	pb "github.com/cuelabs/sptfy/rpc/sptfyapi"
	"github.com/cuelabs/sptfy/internal/models"
	"log"
	"github.com/gorilla/mux"
)

type Global struct {
	db models.Store
}


func newRouter(g *Global) *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.Handle(pb.TrackSearchServicePathPrefix, searchTrackHandler(g))
	//r.Handle(pb.ArtistSearchServicePathPrefix, searchArtistHandler(g))
	//r.Handle(pb.AlbumSearchServicePathPrefix, searchAlbumHandler(g))
	return r
}


func main() {
	db, err := models.NewDB("postgres://postgres:sptfy-dev-password/sptfy")
	if err != nil {
		log.Print("Could not create a new database in main(). Exiting")
		log.Panic(err)
	}

	globals := &Global{db} // implement the methods in respective data models

	r := newRouter(globals)
	http.ListenAndServe("10102", r)
}

func searchTrackHandler() http.Handler {
	return pb.NewTrackSearchServiceServer()
}