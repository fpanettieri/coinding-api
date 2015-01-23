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

	rtr.HandleFunc("/game/",	baseHandler)
	rtr.HandleFunc("/game/new",	newHandler)

	http.Handle("/game/", rtr)
}

func baseHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

    q := datastore.NewQuery("Game")

	var games []Game
	if _, getErr := q.GetAll(c, &games); getErr != nil {
  		http.Error(w, getErr.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	
	fmt.Fprint(w, "[")
	for i, g := range games {
        fmt.Fprintf(w, "{name: %s, dev: %s, url: %s}", g.Name, g.Developer, g.Url)
        if(i < len(games) - 1){ fmt.Fprint(w, ",")}
	}
	fmt.Fprint(w, "]")
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

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
