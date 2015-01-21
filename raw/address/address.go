package raw_address

import (
	"github.com/gorilla/mux"
	"fmt"
	"net/http"
)

func init() {
	rtr := mux.NewRouter()

	rtr.HandleFunc("/raw/address/",                                                     baseHandler)
	rtr.HandleFunc("/raw/address/random",			                                    randomHandler)
	rtr.HandleFunc("/raw/address/{hash:[13][a-km-zA-HJ-NP-Z0-9]{26,33}}",               addressHandler)
	rtr.HandleFunc("/raw/address/{hash:[13][a-km-zA-HJ-NP-Z0-9]{26,33}}/transactions",  transactionsHandler)
	rtr.HandleFunc("/raw/address/{hash:[13][a-km-zA-HJ-NP-Z0-9]{26,33}}/unspents",      unspentsHandler)

	http.Handle("/raw/address/", rtr)
}

func baseHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, r.URL.Path)
}

func randomHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(w, r.URL.Path)
}

func addressHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(w, r.URL.Path)
}

func transactionsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(w, r.URL.Path)
}

func unspentsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(w, r.URL.Path)
}
