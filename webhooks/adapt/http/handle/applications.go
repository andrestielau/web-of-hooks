package handle

import (
	"net/http"
	webhooksv1 "woh/webhooks/adapt/http/v1"
)

// CreateApplications implements webhooksv1.ServerInterface.
func (h *Handler) CreateApplications(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// DeleteApplication implements webhooksv1.ServerInterface.
func (h *Handler) DeleteApplication(w http.ResponseWriter, r *http.Request, applicationId string, params webhooksv1.DeleteApplicationParams) {
	panic("unimplemented")
}

// DeleteApplications implements webhooksv1.ServerInterface.
func (h *Handler) DeleteApplications(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// GetApplication implements webhooksv1.ServerInterface.
func (h *Handler) GetApplication(w http.ResponseWriter, r *http.Request, applicationId string) {
	panic("unimplemented")
}

// GetApplications implements webhooksv1.ServerInterface.
func (h *Handler) GetApplications(w http.ResponseWriter, r *http.Request, params webhooksv1.GetApplicationsParams) {
	panic("unimplemented")
}

// GetStats implements webhooksv1.ServerInterface.
func (h *Handler) GetStats(w http.ResponseWriter, r *http.Request, applicationId string) {
	panic("unimplemented")
}

// UpdateApplication implements webhooksv1.ServerInterface.
func (h *Handler) UpdateApplication(w http.ResponseWriter, r *http.Request, applicationId string, params webhooksv1.UpdateApplicationParams) {
	panic("unimplemented")
}
