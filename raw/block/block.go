package raw_block

import (
	"github.com/gorilla/mux"
	"fmt"
	"net/http"
)

func init() {
	rtr := mux.NewRouter()

	rtr.HandleFunc("/raw/block/",                       baseHandler)
	rtr.HandleFunc("/raw/block/last",			        lastHandler)
	rtr.HandleFunc("/raw/block/random",			       	randomHandler)
	rtr.HandleFunc("/raw/block/{hash:[0-9a-f]{64}}",	blockHandler)

	http.Handle("/raw/block/", rtr)
}

func baseHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, r.URL.Path)
}

func lastHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(w, r.URL.Path)
}

func randomHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(w, r.URL.Path)
}

func blockHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(w, r.URL.Path)
}
