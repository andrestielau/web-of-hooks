package handle

import "net/http"

// CreateSecrets implements webhooksv1.ServerInterface.
func (*Handler) CreateSecrets(w http.ResponseWriter, r *http.Request, applicationId string) {
	panic("unimplemented")
}

// GetSecret implements webhooksv1.ServerInterface.
func (h *Handler) GetSecret(w http.ResponseWriter, r *http.Request, secretId string) {
	panic("unimplemented")
}

// ListApplicationSecrets implements webhooksv1.ServerInterface.
func (h *Handler) ListApplicationSecrets(w http.ResponseWriter, r *http.Request, applicationId string) {
	panic("unimplemented")
}

// RotateSecret implements webhooksv1.ServerInterface.
func (h *Handler) RotateSecret(w http.ResponseWriter, r *http.Request, secretId string) {
	panic("unimplemented")
}
