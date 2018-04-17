package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func newRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	//r.Handle(pb.ArtistSearchServicePathPrefix, searchArtistHandler(g))
	//r.Handle(pb.AlbumSearchServicePathPrefix, searchAlbumHandler(g))
	return r
}

func main() {

	r := newRouter()
	http.ListenAndServe("10102", r)
}
