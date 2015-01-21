package raw_block

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
	url := chainUrl(fmt.Sprintf("%s/%s", "blocks", "00000000000000009cc33fe219537756a68ee5433d593034b6dc200b34aa35fa"))
	forwardRequest(url, w, r)
}

func randomHandler(w http.ResponseWriter, r *http.Request) {
	url := chainUrl(fmt.Sprintf("%s/%s", "blocks", "00000000000000009cc33fe219537756a68ee5433d593034b6dc200b34aa35fa"))
	forwardRequest(url, w, r)
}

func blockHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	url := chainUrl(fmt.Sprintf("%s/%s", "blocks", params["hash"]))
	forwardRequest(url, w, r)
}
