package player

import (
	"github.com/gorilla/mux"
	"fmt"
	"net/http"
)

func init() {
	rtr := mux.NewRouter()

	rtr.HandleFunc("/player/", 		 		            baseHandler)
	rtr.HandleFunc("/player/register",                  registerHandler)
	rtr.HandleFunc("/player/login",                     loginHandler)
	rtr.HandleFunc("/player/guest",                     guestHandler)
	rtr.HandleFunc("/player/{id:[A-Za-z0-9_]+}",        getHandler)
	rtr.HandleFunc("/player/{id:[A-Za-z0-9_]+}/bits",	bitsHandler)

	http.Handle("/player/", rtr)
}

func baseHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, r.URL.Path)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(w, r.URL.Path)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(w, r.URL.Path)
}

func guestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(w, r.URL.Path)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(w, r.URL.Path)
}

func bitsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(w, r.URL.Path)
}