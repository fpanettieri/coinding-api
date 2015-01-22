package tip

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/mux"

	"appengine"
    "appengine/datastore"
)

func init() {
	rtr := mux.NewRouter()

	rtr.HandleFunc("/tip/", 		baseHandler)
	rtr.HandleFunc("/tip/player", 	tipPlayerHandler)
	rtr.HandleFunc("/tip/dev", 		tipDevHandler)

	http.Handle("/tip/", rtr)
}

func baseHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, r.URL.Path)
}

func authPlayer(w http.ResponseWriter, r *http.Request, ctx appengine.Context) bool {
	var player Player

	// Prepare params
	name := r.FormValue("name")
  	pass := []byte(r.FormValue("pass"))

	stringId := base64.StdEncoding.EncodeToString([]byte("player-" + name))
	key := datastore.NewKey(ctx, "Player", stringId, 0, nil)

	// Retrieve developer
	if getErr := datastore.Get(ctx, key, player); getErr != nil {
        http.Error(w, getErr.Error(), http.StatusUnauthorized)
        return false
    }

	// Compare pass
	compareErr := bcrypt.CompareHashAndPassword(player.Pass, pass)
	if compareErr != nil {
  		http.Error(w, "Invalid name or pass", http.StatusUnauthorized)
		return false
	}

	return true
}

func tipPlayerHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	if !authPlayer(w, r, ctx) { return }

	// Validate p1 can send the tip
	var p1 Player
	p1Id := base64.StdEncoding.EncodeToString([]byte("player-" + r.FormValue("name")))
	p1Key := datastore.NewKey(ctx, "Player", p1Id, 0, nil)

	amount, parseErr := strconv.ParseFloat(r.FormValue("amount"), 64)
	if parseErr != nil {
        http.Error(w, parseErr.Error(), http.StatusInternalServerError)
        return
    }

    if(amount <= 0){
		http.Error(w, "Invalid amount", http.StatusPaymentRequired)
  		return
	}

	if(p1.Balance < amount){ amount = p1.Balance }

	if(amount <= 0){
		http.Error(w, "Insufficient funds", http.StatusPaymentRequired)
  		return
	}

	// Prepare p2
	var p2 Player
	p2Id := base64.StdEncoding.EncodeToString([]byte("player-" + r.FormValue("to")))
	p2Key := datastore.NewKey(ctx, "Player", p2Id, 0, nil)

	if getErr := datastore.Get(ctx, p2Key, &p2); getErr != nil {
        http.Error(w, getErr.Error(), http.StatusUnauthorized)
        return
    }

    // Make the transaction
    p1.Balance = p1.Balance - amount
    p2.Balance = p2.Balance + amount

    datastore.Put(ctx, p1Key, p1)
    datastore.Put(ctx, p2Key, p2)
}

func tipDevHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	if !authPlayer(w, r, ctx) { return }

	// Validate p1 can send the tip
	var player Player
	playerId := base64.StdEncoding.EncodeToString([]byte("player-" + r.FormValue("name")))
	playerKey := datastore.NewKey(ctx, "Player", playerId, 0, nil)

	amount, parseErr := strconv.ParseFloat(r.FormValue("amount"), 64)
	if parseErr != nil {
        http.Error(w, parseErr.Error(), http.StatusInternalServerError)
        return
    }

    if(amount <= 0){
		http.Error(w, "Invalid amount", http.StatusPaymentRequired)
  		return
	}

	if(player.Balance < amount){ amount = player.Balance }

	if(amount <= 0){
		http.Error(w, "Insufficient funds", http.StatusPaymentRequired)
  		return
	}

	// Prepare p2
	var dev Developer
	devId := base64.StdEncoding.EncodeToString([]byte("dev-" + r.FormValue("to")))
	devKey := datastore.NewKey(ctx, "Developer", devId, 0, nil)

	if getErr := datastore.Get(ctx, devKey, &dev); getErr != nil {
        http.Error(w, getErr.Error(), http.StatusUnauthorized)
        return
    }

    // Make the transaction
    player.Balance = player.Balance - amount
    dev.Balance = dev.Balance + amount

    datastore.Put(ctx, playerKey, player)
    datastore.Put(ctx, devKey, dev)
}
