package developer

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/mux"
	"github.com/fabioberger/coinbase-go"

	"appengine"
    "appengine/datastore"
)

func init() {
	rtr := mux.NewRouter()

	rtr.HandleFunc("/developer/", 			baseHandler)
	rtr.HandleFunc("/developer/register", 	registerHandler)
	rtr.HandleFunc("/developer/validate", 	validateHandler)
	rtr.HandleFunc("/developer/auth", 		authHandler)
	rtr.HandleFunc("/developer/games", 		gamesHandler)
	rtr.HandleFunc("/developer/deposit", 	depositHandler)

	http.Handle("/developer/", rtr)
}

func baseHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, r.URL.Path)
}

func emailUsed(email string) bool {
	// TODO: validate email has not been used already
	return false
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	// Encrypt email
	email := r.FormValue("email")
  	if emailUsed(email) {
  		http.Error(w, "Email already used", http.StatusInternalServerError)
		return
  	}

  	// Encrypt pass
  	pass, passErr := bcrypt.GenerateFromPassword([]byte(r.FormValue("pass")), 10);
  	if passErr != nil {
  		http.Error(w, passErr.Error(), http.StatusInternalServerError)
		return
	}

	// Create record
	developer := &Developer{
        Email: email,
        Pass:  pass,
        Auth:  time.Now(),
        Balance: 0,
    }

    // FIXME: using only a base64 looks bad, but we need to get this out fast
    stringId := base64.StdEncoding.EncodeToString([]byte("dev-" + email))

	// Configure coinbase api
	callbackUrl := "http://api.coinding.com/developer/" + stringId + "/deposit"
	coinbs := coinbase.ApiKeyClient(COINBASE_KEY, COINBASE_SECRET)
	addressParams := coinbase.AddressParams{ Label: "Developer address", CallbackUrl: callbackUrl }
	
	// Update developer bitcoin address
	address, addressErr := coinbs.GenerateReceiveAddress(&addressParams)
	if addressErr == nil {
		developer.Address = address
	} else {
		http.Error(w, addressErr.Error(), http.StatusInternalServerError)
        return
	}

	// Store new developer
	key := datastore.NewKey(ctx, "Developer", stringId, 0, nil)
	_, dataErr := datastore.Put(ctx, key, developer)
	if dataErr != nil {
        http.Error(w, dataErr.Error(), http.StatusInternalServerError)
        return
    }

	// Debug data
	fmt.Fprintf(w, "Developer email is %s\n", email)
	fmt.Fprintf(w, "Developer pass is %s\n", pass)
	fmt.Fprintf(w, "Developer key is %s\n", key.StringID())
	fmt.Fprintf(w, "Callback url is %s\n", callbackUrl)

	fmt.Fprintf(w, "Developer address is %s", address)
}

func validateHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(w, r.URL.Path)
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(w, r.URL.Path)
}

func gamesHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(w, r.URL.Path)
}

func depositHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(w, r.URL.Path)
}
