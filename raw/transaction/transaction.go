package raw_transaction

import (
	"fmt"
	"net/http"
	"net/url"
	"io/ioutil"
	"log"

	"github.com/gorilla/mux"

	"appengine"
    "appengine/urlfetch"
)

///-- CHAIN.COM Library

const API_KEY = "7e72affc260a0e1d1f13a5a01f3d64e0"
const BASE_URL = "https://api.chain.com/v2/bitcoin/"

func chainUrl(path string) string {
	params := url.Values{}
	return chainUrlParams(path, params)
}

func chainUrlParams(path string, params url.Values) string {
	params.Add("api-key-id", API_KEY)
	return fmt.Sprintf("%s%s?%s", BASE_URL, path, params.Encode())
}

func forwardRequest(url string, w http.ResponseWriter, r *http.Request){
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)

	resp, err := client.Get(url)
    if err != nil {
    	log.Fatal(err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    defer resp.Body.Close()

    bodyBytes, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(err)
        http.Error(w, readErr.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(bodyBytes))
}

///!- CHAIN.COM Library

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
	url := chainUrl(fmt.Sprintf("%s/%s", "transactions", "f5e26c8b82401c585235c572ba8265f16f7d9304ed8e31c198eab571754f5331"))
	forwardRequest(url, w, r)
}

func transactionHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	url := chainUrl(fmt.Sprintf("%s/%s", "transactions", params["hash"]))
	forwardRequest(url, w, r)
}

func confidenceHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	url := chainUrl(fmt.Sprintf("%s/%s/confidence", "transactions", params["hash"]))
	forwardRequest(url, w, r)
}

func hexHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	url := chainUrl(fmt.Sprintf("%s/%s/hex", "transactions", params["hash"]))
	forwardRequest(url, w, r)
}
