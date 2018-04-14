package main

import (
	"net/http"
	"github.com/gorilla/mux"
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