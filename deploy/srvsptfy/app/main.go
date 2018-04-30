package main

import (
	"net/http"
        "github.com/gorilla/mux"
        "io"
)

func newRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
        r.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
io.WriteString(w, "Works.")
     })

        return r
}


func main() {
   r := newRouter()
	http.ListenAndServe("10101", r)
}
