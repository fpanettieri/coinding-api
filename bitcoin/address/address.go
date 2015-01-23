package bitcoin_address

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/fpanettieri/chain-go"
)

func init() {
	rtr := mux.NewRouter()

	rtr.HandleFunc("/bitcoin/address/",                                                     baseHandler)
	rtr.HandleFunc("/bitcoin/address/random",			                                    randomHandler)
	rtr.HandleFunc("/bitcoin/address/{hash:[13][a-km-zA-HJ-NP-Z0-9]{26,33}}",               addressHandler)
	rtr.HandleFunc("/bitcoin/address/{hash:[13][a-km-zA-HJ-NP-Z0-9]{26,33}}/transactions",  transactionsHandler)
	rtr.HandleFunc("/bitcoin/address/{hash:[13][a-km-zA-HJ-NP-Z0-9]{26,33}}/unspents",      unspentsHandler)

	http.Handle("/bitcoin/address/", rtr)
}

func baseHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, r.URL.Path)
}

func randomHandler(w http.ResponseWriter, r *http.Request) {
	url := chain.ChainUrl(fmt.Sprintf("%s/%s", "addresses", "17x23dNjXJLzGMev6R63uyRhMWP1VHawKc"), CHAIN_KEY)
	w.Header().Set("Content-Type", "application/json")
	chain.ForwardRequest(url, w, r)
}

func addressHandler(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	url := chain.ChainUrl(fmt.Sprintf("%s/%s", "addresses", params["hash"]), CHAIN_KEY)
	w.Header().Set("Content-Type", "application/json")
	chain.ForwardRequest(url, w, r)
}

func transactionsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	url := chain.ChainUrl(fmt.Sprintf("%s/%s/transactions", "addresses", params["hash"]), CHAIN_KEY)
	w.Header().Set("Content-Type", "application/json")
	chain.ForwardRequest(url, w, r)
}

func unspentsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	url := chain.ChainUrl(fmt.Sprintf("%s/%s/unspents", "addresses", params["hash"]), CHAIN_KEY)
	w.Header().Set("Content-Type", "application/json")
	chain.ForwardRequest(url, w, r)
}
