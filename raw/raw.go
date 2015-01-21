package raw

import (
  "github.com/gorilla/mux"
  "fmt"
  "net/http"
)

func init() {
	rtr := mux.NewRouter()

    rtr.HandleFunc("/raw/", 	baseHandler)
    rtr.HandleFunc("/raw/data", dataHandler)

    http.Handle("/raw/", rtr)
}

func baseHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
    fmt.Fprint(w, r.URL.Path)
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
    /* TODO: log received data? */
}
