package raw_block

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/fpanettieri/chain-go"
)

func init() {
	rtr := mux.NewRouter()

	rtr.HandleFunc("/raw/block/",                       baseHandler)
	rtr.HandleFunc("/raw/block/last",			        lastHandler)
	rtr.HandleFunc("/raw/block/random",			       	randomHandler)
	rtr.HandleFunc("/raw/block/{hash:[0-9a-f]{64}}",	blockHandler)

	http.Handle("/raw/block/", rtr)
}

func baseHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, r.URL.Path)
}

func lastHandler(w http.ResponseWriter, r *http.Request) {
	url := chain.ChainUrl(fmt.Sprintf("%s/%s", "blocks", "00000000000000009cc33fe219537756a68ee5433d593034b6dc200b34aa35fa"), CHAIN_KEY)
	chain.ForwardRequest(url, w, r)
}

func randomHandler(w http.ResponseWriter, r *http.Request) {
	url := chain.ChainUrl(fmt.Sprintf("%s/%s", "blocks", "00000000000000009cc33fe219537756a68ee5433d593034b6dc200b34aa35fa"), CHAIN_KEY)
	chain.ForwardRequest(url, w, r)
}

func blockHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	url := chain.ChainUrl(fmt.Sprintf("%s/%s", "blocks", params["hash"]), CHAIN_KEY)
	chain.ForwardRequest(url, w, r)
}
