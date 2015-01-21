package bitcoin

import (
	"github.com/gorilla/mux"
	"fmt"
	"net/http"
)

func init() {
	rtr := mux.NewRouter()

	rtr.HandleFunc("/bitcoin/",         baseHandler)
	rtr.HandleFunc("/bitcoin/send",     sendHandler)
	rtr.HandleFunc("/bitcoin/deposit",  depositHandler)
	rtr.HandleFunc("/bitcoin/withdraw", withdrawHandler)

	http.Handle("/bitcoin/", rtr)
}

func baseHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, r.URL.Path)
}

func sendHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(w, r.URL.Path)
}

func depositHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(w, r.URL.Path)
}

func withdrawHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(w, r.URL.Path)
}
