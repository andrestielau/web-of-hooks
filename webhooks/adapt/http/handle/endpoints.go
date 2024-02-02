package handle

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"woh/package/utils/media"
	"woh/webhooks"
	webhooksv1 "woh/webhooks/adapt/http/v1"
	"woh/webhooks/render/page/applications"

	"github.com/samber/lo"
)

// CreateEndpoints implements webhooksv1.ServerInterface.
func (h *Handler) CreateEndpoints(w http.ResponseWriter, r *http.Request, applicationId string) {
	var req webhooksv1.CreateEndpointsPayload
	var ret webhooksv1.CreatedEndpoints
	// TODO:
	if r.Body == nil {
		return
	}
	var err error
	defer r.Body.Close()
	if strings.Contains(r.Header.Get("Content-Type"), "form") {
		req = []webhooksv1.NewEndpoint{{
			Name:          lo.ToPtr(r.FormValue("name")),
			Url:           r.FormValue("url"),
			FilterTypeIds: lo.ToPtr([]string{}),
		}}
		log.Println(req)
	} else {
		err = json.NewDecoder(r.Body).Decode(&req)
	}
	// ENDTODO
	if err != nil { // err := media.Req(r, &req);
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i := range req {
		req[i].ApplicationId = &applicationId
	}
	res, err := h.Repo.CreateEndpoints(r.Context(), h.Convert.NewEndpoints(req))
	if err != nil {
		errs, stop := webhooks.HttpErrors(w, err)
		if stop {
			return
		}
		ret.Errors = errs
	} else {
		ret.Data = h.Convert.EndpointDetails(res)
	}
	if media.ShouldRender(r) {
		applications.EndpointItems(applicationId, lo.Map(res, func(e webhooks.EndpointDetails, _ int) webhooks.Endpoint {
			return e.Endpoint
		}))
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
	if params.Limit == nil {
		params.Limit = lo.ToPtr(LIMIT_DEFAULT)
	}
	if res, err := h.Repo.ListApplicationEndpoints(r.Context(), applicationId, h.Convert.EndpointQuery(params)); err != nil {
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
