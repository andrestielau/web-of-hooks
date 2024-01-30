package media

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ajg/form"
)

type NegotiatorParser struct {
}

func Accept(r *http.Request) string { return r.Header.Get("Accept") }

func Req(r *http.Request, ret any) (err error) {
	if r.Body == nil {
		return
	}
	defer r.Body.Close()
	if strings.Contains(r.Header.Get("Content-Type"), "form-data") {
		err = form.NewDecoder(r.Body).Decode(&ret)
	} else {
		err = json.NewDecoder(r.Body).Decode(&ret)
	}
	return
}
func ShouldRender(r *http.Request) bool { return strings.Contains(r.Header.Get("Accept"), "text/html") }
func Res(w http.ResponseWriter, ct string, t any) (err error) {
	switch {
	default:
		err = json.NewEncoder(w).Encode(t)
	}
	return
}
