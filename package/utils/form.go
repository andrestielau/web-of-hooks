package utils

import (
	"net/http"

	"github.com/ajg/form"
)

func FormReq[T any](w http.ResponseWriter, r *http.Request) (req T, err error) {
	if err = form.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Form could not be decoded", http.StatusBadRequest)
	}
	return
}
