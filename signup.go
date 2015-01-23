package developer

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"appengine"
    "appengine/datastore"
)

type Contact struct {
    Email		string
}

func init() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/signup/", signupHandler)
	http.Handle("/signup/", rtr)
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	contact := &Contact{ Email: r.FormValue("email") }

    // Store new developer
	key := datastore.NewIncompleteKey(ctx, "Contact", nil)
	_, dataErr := datastore.Put(ctx, key, contact)
	if dataErr != nil {
        http.Error(w, dataErr.Error(), http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "http://coinding.com", 301)
}

func signupAllHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	q := datastore.NewQuery("Contact")

    w.Header().Set("Content-Type", "application/json")

	var cons []Contact
	if _, getErr := q.GetAll(c, &cons); getErr != nil {
  		http.Error(w, getErr.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "[\n")
	for i, d := range cons {
        fmt.Fprintf(w, "  {email: %s}\n", d.Email)
        if(i < len(cons) - 1){ fmt.Fprint(w, ",")}
	}
	fmt.Fprint(w, "]")
}