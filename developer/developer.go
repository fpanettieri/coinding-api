package developer

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

	rtr.HandleFunc("/developer/", 				baseHandler)
	rtr.HandleFunc("/developer/register", 		registerHandler)
	rtr.HandleFunc("/developer/balance", 		balanceHandler)
	rtr.HandleFunc("/developer/deposit", 		depositHandler)
	rtr.HandleFunc("/developer/withdraw", 		withdrawHandler)
	rtr.HandleFunc("/developer/coinbase", 		coinbaseHandler)

	http.Handle("/developer/", rtr)
}

func baseHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

    q := datastore.NewQuery("Developer").Project("Name")

	var devs []Developer
	if _, getErr := q.GetAll(c, &devs); getErr != nil {
  		http.Error(w, getErr.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "[")
	for i, d := range devs {
        fmt.Fprintf(w, "{name: %s}", d.Name)
        if(i < len(devs) - 1){ fmt.Fprint(w, ",")}
	}
	fmt.Fprint(w, "]")
}

func nameUsed(w http.ResponseWriter, ctx appengine.Context, name string) bool {
	count, err := datastore.NewQuery("Developer").Filter("Name =", name).KeysOnly().Count(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return count > 0
}

func emailUsed(w http.ResponseWriter, ctx appengine.Context, email string) bool {
	count, err := datastore.NewQuery("Developer").Filter("Email =", email).KeysOnly().Count(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return count > 0
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	name := r.FormValue("name")
	if emailUsed(w, ctx, name) {
  		http.Error(w, "Name already used", http.StatusInternalServerError)
		return
  	}

	// Check email
	email := r.FormValue("email")
  	if emailUsed(w, ctx, email) {
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
		Name: name,
        Email: email,
        Pass:  pass,
        Balance: 0,
    }

    // FIXME: using only a base64 looks bad, but we need to get this out fast
    stringId := base64.StdEncoding.EncodeToString([]byte("dev-" + email))

	// Configure coinbase api
	callbackUrl := "http://api.coinding.com/developer/coinbase?id=" + stringId
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
}

func authDev(w http.ResponseWriter, r *http.Request, ctx appengine.Context, dev *Developer) bool {
	// Prepare params
	email := r.FormValue("email")
  	pass := []byte(r.FormValue("pass"))
	stringId := base64.StdEncoding.EncodeToString([]byte("dev-" + email))
	key := datastore.NewKey(ctx, "Developer", stringId, 0, nil)

	// Retrieve developer
	if getErr := datastore.Get(ctx, key, dev); getErr != nil {
        http.Error(w, getErr.Error(), http.StatusUnauthorized)
        return false
    }

	// Compare pass
	compareErr := bcrypt.CompareHashAndPassword(dev.Pass, pass)
	if compareErr != nil {
  		http.Error(w, "Invalid email or pass", http.StatusUnauthorized)
		return false
	}

	return true
}

func balanceHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	var dev Developer
	if !authDev(w, r, ctx, &dev) { return }
	
	fmt.Fprintf(w, "{email: %s, balance: %f}", dev.Email, dev.Balance)
}

func depositHandler(w http.ResponseWriter, r *http.Request) {
	coinbs := coinbase.ApiKeyClient(COINBASE_KEY, COINBASE_SECRET)
	depositParams := coinbase.TransactionParams{ To: r.FormValue("email"), From: "api@coinding.com", Amount: r.FormValue("amount") }
	_, err := coinbs.RequestMoney(&depositParams)
	
	if err != nil {
  		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func withdrawHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	var dev Developer
	if !authDev(w, r, ctx, &dev) { return }

	amount, parseErr := strconv.ParseFloat(r.FormValue("amount"), 64)
	if parseErr != nil {
        http.Error(w, parseErr.Error(), http.StatusInternalServerError)
        return
    }

	if(dev.Balance < amount){ amount = dev.Balance }

	if(amount <= 0){
		http.Error(w, "Insufficient funds", http.StatusPaymentRequired)
  		return
	}
	
	// FIXME: using only a base64 looks bad, but we need to get this out fast
    stringId := base64.StdEncoding.EncodeToString([]byte("dev-" + dev.Email))
	dev.Balance = dev.Balance - amount

	// Store new developer
	key := datastore.NewKey(ctx, "Developer", stringId, 0, nil)
	_, dataErr := datastore.Put(ctx, key, dev)
	if dataErr != nil {
        http.Error(w, dataErr.Error(), http.StatusInternalServerError)
        return
    }

    // Withdraw money
    amountStr := strconv.FormatFloat(amount, 'f', 8, 64)
	coinbs := coinbase.ApiKeyClient(COINBASE_KEY, COINBASE_SECRET)
	depositParams := coinbase.TransactionParams{ To: r.FormValue("destination"), From: dev.Email, Amount: amountStr, Notes: "Sent via Coinding" }

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
