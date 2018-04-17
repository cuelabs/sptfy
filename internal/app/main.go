package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/cuelabs/sptfy/rpc/sptfyapi"
	"github.com/cuelabs/sptfy/internal/searchartistservice"
)

func newRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	//r.Handle(pb.ArtistSearchServicePathPrefix, searchArtistHandler(g))
	//r.Handle(pb.AlbumSearchServicePathPrefix, searchAlbumHandler(g))
	return r
}

func main() {

	artistSearchServer := searchartistserver.
	artistSearchHandler := http.Handle(pb.ArtistSearchServicePathPrefix,
	http.ListenAndServe("10102", nil)
}
