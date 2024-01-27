package utils

import (
	"encoding/json"
	"net/http"
)

// JSON Encoding and Decoding

func JsonReq[T any](w http.ResponseWriter, r *http.Request) (req T, err error) {
	if r.Body == nil {
		return
	}
	defer r.Body.Close()
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	return
}

func JsonRes[T any](w http.ResponseWriter, res T) error {
	return json.NewEncoder(w).Encode(res)
}
