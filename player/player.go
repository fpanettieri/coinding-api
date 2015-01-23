package player

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/mux"
	"github.com/fabioberger/coinbase-go"

	"appengine"
    "appengine/datastore"
)

func init() {
	rtr := mux.NewRouter()

	rtr.HandleFunc("/player/", 				baseHandler)
	rtr.HandleFunc("/player/register", 		registerHandler)
	rtr.HandleFunc("/player/balance", 		balanceHandler)
	rtr.HandleFunc("/player/deposit", 		depositHandler)
	rtr.HandleFunc("/player/withdraw", 		withdrawHandler)
	rtr.HandleFunc("/player/coinbase", 		coinbaseHandler)

	http.Handle("/player/", rtr)
}

func baseHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

    q := datastore.NewQuery("Player").Project("Name")

	var players []Player
	if _, getErr := q.GetAll(c, &players); getErr != nil {
  		http.Error(w, getErr.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprint(w, "[")
	for i, p := range players {
        fmt.Fprintf(w, "{name: %s}", p.Name)
        if(i < len(players) - 1){ fmt.Fprint(w, ",")}
	}
	fmt.Fprint(w, "]")
}

func nameUsed(w http.ResponseWriter, ctx appengine.Context, name string) bool {
	count, err := datastore.NewQuery("Player").Filter("Name =", name).KeysOnly().Count(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return count > 0
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	// Check name
	name := r.FormValue("name")
  	if nameUsed(w, ctx, name) {
  		http.Error(w, "Name already used", http.StatusInternalServerError)
		return
  	}

  	// Encrypt pass
  	pass, passErr := bcrypt.GenerateFromPassword([]byte(r.FormValue("pass")), 10);
  	if passErr != nil {
  		http.Error(w, passErr.Error(), http.StatusInternalServerError)
		return
	}

	// Create record
	player := &Player{
        Name: name,
        Pass: pass,
        Balance: PLAYER_FUNDS,
    }

    // FIXME: using only a base64 looks bad, but we need to get this out fast
    stringId := base64.StdEncoding.EncodeToString([]byte("player-" + name))

	// Configure coinbase api
	callbackUrl := "http://api.coinding.com/player/coinbase?id=" + stringId
	coinbs := coinbase.ApiKeyClient(COINBASE_KEY, COINBASE_SECRET)
	addressParams := coinbase.AddressParams{ Label: "Player address", CallbackUrl: callbackUrl }
	
	// Update developer bitcoin address
	address, addressErr := coinbs.GenerateReceiveAddress(&addressParams)
	if addressErr == nil {
		player.Address = address
	} else {
		http.Error(w, addressErr.Error(), http.StatusInternalServerError)
        return
	}

	// Store new developer
	key := datastore.NewKey(ctx, "Player", stringId, 0, nil)
	_, dataErr := datastore.Put(ctx, key, player)
	if dataErr != nil {
        http.Error(w, dataErr.Error(), http.StatusInternalServerError)
        return
    }
}

func authPlayer(w http.ResponseWriter, r *http.Request, ctx appengine.Context, player *Player) bool {
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

func balanceHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	var player Player
	if !authPlayer(w, r, ctx, &player) { return }
	
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{name: %s, balance: %f}", player.Name, player.Balance)
}

func depositHandler(w http.ResponseWriter, r *http.Request) {
	coinbs := coinbase.ApiKeyClient(COINBASE_KEY, COINBASE_SECRET)
	depositParams := coinbase.TransactionParams{ To: r.FormValue("name"), From: "api@coinding.com", Amount: r.FormValue("amount"), Notes: "Sent via Coinding" }
	_, err := coinbs.RequestMoney(&depositParams)
	
	if err != nil {
  		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func withdrawHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	var player Player
	if !authPlayer(w, r, ctx, &player) { return }

	amount, parseErr := strconv.ParseFloat(r.FormValue("amount"), 64)
	if parseErr != nil {
        http.Error(w, parseErr.Error(), http.StatusInternalServerError)
        return
    }

	if(player.Balance < amount){ amount = player.Balance }

	if(amount <= 0){
		http.Error(w, "Insufficient funds", http.StatusPaymentRequired)
  		return
	}
	
	// FIXME: using only a base64 looks bad, but we need to get this out fast
    stringId := base64.StdEncoding.EncodeToString([]byte("player-" + player.Name))
	player.Balance = player.Balance - amount

	// Store new developer
	key := datastore.NewKey(ctx, "Player", stringId, 0, nil)
	_, dataErr := datastore.Put(ctx, key, player)
	if dataErr != nil {
        http.Error(w, dataErr.Error(), http.StatusInternalServerError)
        return
    }

    // Withdraw money
    amountStr := strconv.FormatFloat(amount, 'f', 8, 64)
	coinbs := coinbase.ApiKeyClient(COINBASE_KEY, COINBASE_SECRET)
	depositParams := coinbase.TransactionParams{ To: r.FormValue("destination"), From: player.Name, Amount: amountStr, Notes: "Sent via Coinding" }

	_, err := coinbs.SendMoney(&depositParams)
	
	if err != nil {
  		http.Error(w, err.Error(), http.StatusInternalServerError)
  		return
	}
}

func coinbaseHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(w, r.URL.Path)
}
