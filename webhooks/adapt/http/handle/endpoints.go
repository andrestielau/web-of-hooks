package handle

import (
	"net/http"
	"woh/package/utils/media"
	webhooksv1 "woh/webhooks/adapt/http/v1"
	"woh/webhooks/render/page/applications"

	"github.com/andrestielau/web-of-hooks/webhooks"
)

// CreateEndpoints implements webhooksv1.ServerInterface.
func (h *Handler) CreateEndpoints(w http.ResponseWriter, r *http.Request, applicationId string) {
	var req webhooksv1.CreateEndpointsPayload
	var ret webhooksv1.CreatedEndpoints
	if err := media.Req(r, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if res, err := h.Repo.CreateEndpoints(r.Context(), h.Convert.NewEndpoints(req)); err != nil {
		errs, stop := webhooks.HttpErrors(w, err)
		if stop {
			return
		}
		ret.Errors = errs
	} else {
		ret.Data = h.Convert.EndpointDetails(res)
	}
	if media.ShouldRender(r) {
		// TODO: partial
	} else if err := media.Res(w, media.Accept(r), ret); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// DeleteEndpoint implements webhooksv1.ServerInterface.
func (h *Handler) DeleteEndpoint(w http.ResponseWriter, r *http.Request, endpointId string, params webhooksv1.DeleteEndpointParams) {
	h.Repo.DeleteEndpoints(r.Context(), []string{endpointId})
}

// DisableEndpoints implements webhooksv1.ServerInterface.
func (h *Handler) DisableEndpoints(w http.ResponseWriter, r *http.Request, applicationId string, params webhooksv1.DisableEndpointsParams) {
	// TODO: update endpoints
}

// GetEndpoint implements webhooksv1.ServerInterface.
func (h *Handler) GetEndpoint(w http.ResponseWriter, r *http.Request, endpointId string) {
	var ret webhooksv1.Endpoint
	if res, err := h.Repo.GetEndpoints(r.Context(), []string{endpointId}); webhooks.HttpError(w, err) {
		return
	} else if len(res) == 1 {
		ret = h.Convert.EndpointDetail(res[0])
	}
	if media.ShouldRender(r) {
		// TODO: partial
	} else if err := media.Res(w, media.Accept(r), ret); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetEndpointStats implements webhooksv1.ServerInterface.
func (h *Handler) GetEndpointStats(w http.ResponseWriter, r *http.Request, endpointId string) {
	panic("unimplemented")
}

// ListEndpoints implements webhooksv1.ServerInterface.
func (h *Handler) ListEndpoints(w http.ResponseWriter, r *http.Request, applicationId string, params webhooksv1.ListEndpointsParams) {
	if res, err := h.Repo.ListEndpoints(r.Context(), h.Convert.EndpointQuery(params)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if media.ShouldRender(r) {
		applications.Endpoints(applications.EndpointsViewModel{ // Todo decouple from DB
			ApplicationId: applicationId,
			Data:          nil,
		}, nil).Render(r.Context(), w)
	} else {
		media.Res(w, media.Accept(r), res)
	}
}

// UpdateEndpoint implements webhooksv1.ServerInterface.
func (h *Handler) UpdateEndpoint(w http.ResponseWriter, r *http.Request, endpointId string) {
	panic("unimplemented")
}
