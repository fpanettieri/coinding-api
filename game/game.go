package game

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"appengine"
    "appengine/datastore"
)

func init() {
	rtr := mux.NewRouter()

	rtr.HandleFunc("/game/", 		 						baseHandler)
	rtr.HandleFunc("/game/all",								allHandler)
	rtr.HandleFunc("/game/new",								newHandler)

	http.Handle("/game/", rtr)
}

func baseHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, r.URL.Path)
}

func allHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(w, r.URL.Path)
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	// Check email
	name := r.FormValue("name")
	dev := r.FormValue("dev")
	url := r.FormValue("url")

	// Create record
	game := &Game{
        Name: name,
        Developer: dev,
        Url: url,
    }

    // Store new developer
	key := datastore.NewKey(ctx, "Game", "", 0, nil)
	_, dataErr := datastore.Put(ctx, key, game)
	if dataErr != nil {
        http.Error(w, dataErr.Error(), http.StatusInternalServerError)
        return
    }
}
