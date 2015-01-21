package developer

import (
	"github.com/gorilla/mux"
	"fmt"
	"net/http"
)

func init() {
	rtr := mux.NewRouter()

	rtr.HandleFunc("/developer/", 			baseHandler)
	rtr.HandleFunc("/developer/register", 	registerHandler)
	rtr.HandleFunc("/developer/validate", 	validateHandler)
	rtr.HandleFunc("/developer/login", 		loginHandler)
	rtr.HandleFunc("/developer/games", 		gamesHandler)

	http.Handle("/developer/", rtr)
}

func baseHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, r.URL.Path)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(w, r.URL.Path)
}

func validateHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(w, r.URL.Path)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(w, r.URL.Path)
}

func gamesHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(w, r.URL.Path)
}
