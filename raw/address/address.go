package raw_address

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

	rtr.HandleFunc("/raw/address/",                                                     baseHandler)
	rtr.HandleFunc("/raw/address/random",			                                    randomHandler)
	rtr.HandleFunc("/raw/address/{hash:[13][a-km-zA-HJ-NP-Z0-9]{26,33}}",               addressHandler)
	rtr.HandleFunc("/raw/address/{hash:[13][a-km-zA-HJ-NP-Z0-9]{26,33}}/transactions",  transactionsHandler)
	rtr.HandleFunc("/raw/address/{hash:[13][a-km-zA-HJ-NP-Z0-9]{26,33}}/unspents",      unspentsHandler)

	http.Handle("/raw/address/", rtr)
}

func baseHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, r.URL.Path)
}

func randomHandler(w http.ResponseWriter, r *http.Request) {
	url := chainUrl(fmt.Sprintf("%s/%s", "addresses", "17x23dNjXJLzGMev6R63uyRhMWP1VHawKc"))
	forwardRequest(url, w, r)
}

func addressHandler(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	url := chainUrl(fmt.Sprintf("%s/%s", "addresses", params["hash"]))
	forwardRequest(url, w, r)
}

func transactionsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	url := chainUrl(fmt.Sprintf("%s/%s/transactions", "addresses", params["hash"]))
	forwardRequest(url, w, r)
}

func unspentsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	url := chainUrl(fmt.Sprintf("%s/%s/unspents", "addresses", params["hash"]))
	forwardRequest(url, w, r)
}
