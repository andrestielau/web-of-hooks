package handle

import (
	"net/http"
	"strings"
	"woh/package/utils"
	webhooksv1 "woh/webhooks/adapt/http/v1"
	"woh/webhooks/render/page"
)

// CreateEndpoints implements webhooksv1.ServerInterface.
func (h *Handler) CreateEndpoints(w http.ResponseWriter, r *http.Request, applicationId string) {
	panic("unimplemented")
}

// DeleteEndpoint implements webhooksv1.ServerInterface.
func (h *Handler) DeleteEndpoint(w http.ResponseWriter, r *http.Request, endpointId string, params webhooksv1.DeleteEndpointParams) {
	panic("unimplemented")
}

// DisableEndpoints implements webhooksv1.ServerInterface.
func (h *Handler) DisableEndpoints(w http.ResponseWriter, r *http.Request, applicationId string, params webhooksv1.DisableEndpointsParams) {
	panic("unimplemented")
}

// GetEndpoint implements webhooksv1.ServerInterface.
func (h *Handler) GetEndpoint(w http.ResponseWriter, r *http.Request, endpointId string) {
	panic("unimplemented")
}

// GetEndpointStats implements webhooksv1.ServerInterface.
func (h *Handler) GetEndpointStats(w http.ResponseWriter, r *http.Request, endpointId string) {
	panic("unimplemented")
}

// ListEndpoints implements webhooksv1.ServerInterface.
func (h *Handler) ListEndpoints(w http.ResponseWriter, r *http.Request, applicationId string, params webhooksv1.ListEndpointsParams) {
	endpoints, err := h.Repo.ListEndpoints(r.Context(), 100, 0)
	if strings.Contains(r.Header.Get("Accept"), "text/html") {
		page.Endpoints(page.EndpointsViewModel{ // Todo decouple from DB
			Data: endpoints,
		}, err).Render(r.Context(), w)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.JsonRes(w, endpoints)
}

// UpdateEndpoint implements webhooksv1.ServerInterface.
func (h *Handler) UpdateEndpoint(w http.ResponseWriter, r *http.Request, endpointId string) {
	panic("unimplemented")
}
