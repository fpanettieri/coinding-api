package raw_transaction

import (
	"github.com/gorilla/mux"
	"fmt"
	"net/http"
)

func init() {
	rtr := mux.NewRouter()

	rtr.HandleFunc("/raw/transaction/",                                   baseHandler)
	rtr.HandleFunc("/raw/transaction/random",			                        randomHandler)
	rtr.HandleFunc("/raw/transaction/{hash:[a-fA-F0-9]{64}}",             transactionHandler)
	rtr.HandleFunc("/raw/transaction/{hash:[a-fA-F0-9]{64}}/confidence",  confidenceHandler)
	rtr.HandleFunc("/raw/transaction/{hash:[a-fA-F0-9]{64}}/hex",         hexHandler)

	http.Handle("/raw/transaction/", rtr)
}

func baseHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, r.URL.Path)
}

func randomHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(w, r.URL.Path)
}

func transactionHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(w, r.URL.Path)
}

func confidenceHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(w, r.URL.Path)
}

func hexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(w, r.URL.Path)
}
