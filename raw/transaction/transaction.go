package raw_transaction

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/fpanettieri/chain-go"
)

func init() {
	rtr := mux.NewRouter()

	rtr.HandleFunc("/raw/transaction/",                                   baseHandler)
	rtr.HandleFunc("/raw/transaction/random",			                        randomHandler)
	rtr.HandleFunc("/raw/transaction/{hash:[a-fA-F0-9]{64}}",             transactionHandler)
	rtr.HandleFunc("/raw/transaction/{hash:[a-fA-F0-9]{64}}/confidence",  confidenceHandler)
	rtr.HandleFunc("/raw/transaction/{hash:[a-fA-F0-9]{64}}/hex",         hexHandler)

	http.Handle("/raw/transaction/", rtr)
}

func baseHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, r.URL.Path)
}

func randomHandler(w http.ResponseWriter, r *http.Request) {
	url := chain.ChainUrl(fmt.Sprintf("%s/%s", "transactions", "f5e26c8b82401c585235c572ba8265f16f7d9304ed8e31c198eab571754f5331"), CHAIN_KEY)
	chain.ForwardRequest(url, w, r)
}

func transactionHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	url := chain.ChainUrl(fmt.Sprintf("%s/%s", "transactions", params["hash"]), CHAIN_KEY)
	chain.ForwardRequest(url, w, r)
}

func confidenceHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	url := chain.ChainUrl(fmt.Sprintf("%s/%s/confidence", "transactions", params["hash"]), CHAIN_KEY)
	chain.ForwardRequest(url, w, r)
}

func hexHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	url := chain.ChainUrl(fmt.Sprintf("%s/%s/hex", "transactions", params["hash"]), CHAIN_KEY)
	chain.ForwardRequest(url, w, r)
}
