package handle

import (
	"net/http"
	"woh/package/utils/media"

	webhooksv1 "woh/webhooks/adapt/http/v1"
	"github.com/andrestielau/web-of-hooks/webhooks/adapt/http/convert"
)

// CreateSecrets implements webhooksv1.ServerInterface.
func (*Handler) CreateSecrets(w http.ResponseWriter, r *http.Request, applicationId string) {
	panic("unimplemented")
}

// GetSecret implements webhooksv1.ServerInterface.
func (h *Handler) GetSecret(w http.ResponseWriter, r *http.Request, secretId string) {
	var ret webhooksv1.Secret
	if res, err := h.Repo.GetSecrets(r.Context(), []string{secretId}); convert.Error(w, err) {
		return
	} else if len(res) == 1 {
		ret = h.Convert.Secret(res[0])
	}
	if media.ShouldRender(r) {
		// TODO: partial
	} else if err := media.Res(w, media.Accept(r), ret); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// ListApplicationSecrets implements webhooksv1.ServerInterface.
func (h *Handler) ListApplicationSecrets(w http.ResponseWriter, r *http.Request, applicationId string) {
	panic("unimplemented")
}

// RotateSecret implements webhooksv1.ServerInterface.
func (h *Handler) RotateSecret(w http.ResponseWriter, r *http.Request, secretId string) {
	panic("unimplemented")
}


// CreateSecret implements webhooksv1.ServerInterface.
func (*Handler) CreateSecret(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// ListSecrets implements webhooksv1.ServerInterface.
func (*Handler) ListSecrets(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}