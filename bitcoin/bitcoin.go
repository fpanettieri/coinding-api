package bitcoin

import (
  "github.com/gorilla/mux"
  "fmt"
  "net/http"
)

func init() {
	rtr := mux.NewRouter()

    rtr.HandleFunc("/bitcoin/",       baseHandler)
    rtr.HandleFunc("/bitcoin/notify", notifyHandler)

    http.Handle("/bitcoin/", rtr)
}

func baseHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
    fmt.Fprint(w, r.URL.Path)
}

func notifyHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
    /* TODO: log received data? */
}
