package main

import (
	"github.com/gorilla/mux"
	"net/http"
	pb "github.com/cuelabs/sptfy/rpc/sptfyapi"
)

func newRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	//r.Handle(pb.ArtistSearchServicePathPrefix, searchArtistHandler(g))
	//r.Handle(pb.AlbumSearchServicePathPrefix, searchAlbumHandler(g))
	return r
}

func main() {

        http.Handle(pb.ArtistSearchServicePathPrefix,

	r := newRouter()
	http.ListenAndServe("10102", nil)
}
