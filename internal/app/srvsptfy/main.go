package main

import (
	"github.com/gorilla/mux"
	"net/http"
	pb "github.com/cuelabs/sptfy/rpc/sptfyapi"
	"github.com/cuelabs/sptfy/internal/searchartistservice"
)

func newRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
        is := infoservice.InfoService
	r.Handle(pb.InfoServicePathPrefix,

        //r.Handle(pb.ArtistSearchServicePathPrefix, searchArtistHandler(g))
	//r.Handle(pb.AlbumSearchServicePathPrefix, searchAlbumHandler(g))
	return r
}

func main() {

	artistSearchServer := searchartistservice.ArtistSearchServer{}
	artistSearchHandler := http.Handle(pb.ArtistSearchServicePathPrefix,
	http.ListenAndServe("10102", nil)
}
