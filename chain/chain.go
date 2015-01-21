package chain

import (
	"fmt"
	"net/url"
)

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
