package handle

import (
	"log"
	"net/http"
	"strings"
	"woh/package/utils"
	webhooksv1 "woh/webhooks/adapt/http/v1"
	"woh/webhooks/provide/repo/queries"
	"woh/webhooks/render/page/applications"
)

// CreateEndpoints implements webhooksv1.ServerInterface.
func (h *Handler) CreateEndpoints(w http.ResponseWriter, r *http.Request, applicationId string) {
	var req webhooksv1.CreateEndpointsPayload
	var err error
	if strings.Contains(r.Header.Get("Content-Type"), "") {
		req, err = utils.FormReq[webhooksv1.CreateEndpointsPayload](w, r)
	} else {
		req, err = utils.JsonReq[webhooksv1.CreateEndpointsPayload](w, r)
	}
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(req)
	h.Repo.CreateEndpoints(r.Context(), []queries.NewEndpoint{})
	if strings.Contains(r.Header.Get("Accept"), "text/html") {
		if err != nil {
			applications.EndpointsError(err).Render(r.Context(), w)
			return
		}
		//page.EndpointItem(endpoint.Uid, endpoint.Url).Render(r.Context(), w)
		return
	}
	utils.JsonRes(w, webhooksv1.ErrorList{}) // TODO: map errors
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
		applications.Endpoints(applications.EndpointsViewModel{ // Todo decouple from DB
			ApplicationId: applicationId,
			Data:          endpoints,
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
