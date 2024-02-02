package handle

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"woh/package/utils/media"
	"woh/webhooks"
	webhooksv1 "woh/webhooks/adapt/http/v1"
	"woh/webhooks/render/page/applications"

	"github.com/samber/lo"
)

// CreateApplications implements webhooksv1.ServerInterface.
func (h *Handler) CreateApplications(w http.ResponseWriter, r *http.Request) {
	var req webhooksv1.CreateApplicationsPayload
	var ret webhooksv1.CreatedApplications
	// TODO:
	if r.Body == nil {
		return
	}
	var err error
	defer r.Body.Close()
	if strings.Contains(r.Header.Get("Content-Type"), "form") {
		req = []webhooksv1.NewApplication{{
			Name:     r.FormValue("name"),
			TenantId: r.FormValue("tenantId"),
		}}
		log.Println(req)
	} else {
		err = json.NewDecoder(r.Body).Decode(&ret)
	}
	// ENDTODO
	if err != nil { // err := media.Req(r, &req);
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := h.Repo.CreateApplications(r.Context(), h.Convert.NewApplications(req))
	if err != nil {
		errs, stop := webhooks.HttpErrors(w, err)
		if stop {
			return
		}
		ret.Errors = errs
	} else {
		ret.Data = h.Convert.Applications(res)
	}
	if media.ShouldRender(r) {
		applications.ApplicationItems(res).Render(r.Context(), w)
	} else if err := media.Res(w, media.Accept(r), ret); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// DeleteApplication implements webhooksv1.ServerInterface.
func (h *Handler) DeleteApplication(w http.ResponseWriter, r *http.Request, applicationId string, params webhooksv1.DeleteApplicationParams) {
	webhooks.HttpError(w, h.Repo.DeleteApplications(r.Context(), []string{applicationId}))
}

// DeleteApplications implements webhooksv1.ServerInterface.
func (h *Handler) DeleteApplications(w http.ResponseWriter, r *http.Request) {
	var req webhooksv1.DeleteApplicationsPayload
	if err := media.Req(r, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	webhooks.HttpError(w, h.Repo.DeleteApplications(r.Context(), lo.Map[struct {
		Id *string "json:\"id,omitempty\""
	}](req, func(ids struct {
		Id *string "json:\"id,omitempty\""
	}, _ int) string {
		return *ids.Id
	})))
}

// GetApplication implements webhooksv1.ServerInterface.
func (h *Handler) GetApplication(w http.ResponseWriter, r *http.Request, applicationId string) {
	var ret webhooksv1.Application
	res, err := h.Repo.GetApplications(r.Context(), []string{applicationId})
	if webhooks.HttpError(w, err) {
		return
	} else if len(res) == 0 {
		http.Error(w, fmt.Sprintf("Application with uid %s not found", applicationId), http.StatusNotFound)
		return
	} else if len(res) == 1 {
		ret = h.Convert.ApplicationDetail(res[0])
	}
	if media.ShouldRender(r) {
		res2, _ := h.Repo.ListEventTypes(r.Context(), webhooks.EventTypeQuery{})
		applications.Application(applications.ApplicationViewModel{
			ApplicationDetails: res[0],
			EventTypes: lo.SliceToMap(res2, func(e webhooks.EventType) (string, string) {
				return e.Uid, e.Key
			}),
		}).Render(r.Context(), w)
	} else if err := media.Res(w, media.Accept(r), ret); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// ListApplications implements webhooksv1.ServerInterface.
func (h *Handler) ListApplications(w http.ResponseWriter, r *http.Request, params webhooksv1.ListApplicationsParams) {
	if params.Limit == nil {
		params.Limit = lo.ToPtr(LIMIT_DEFAULT)
	}
	if res, err := h.Repo.ListApplications(r.Context(), h.Convert.ApplicationQuery(params)); err != nil {
		webhooks.HttpError(w, err)
	} else if media.ShouldRender(r) {
		applications.Applications(applications.ApplicationsViewModel{
			Data: res,
		}, nil).Render(r.Context(), w)
	} else {
		media.Res(w, media.Accept(r), res)
	}
}

// GetApplicationStats implements webhooksv1.ServerInterface.
func (h *Handler) GetApplicationStats(w http.ResponseWriter, r *http.Request, applicationId string) {
	var ret webhooksv1.Application
	if res, err := h.Repo.GetApplications(r.Context(), []string{applicationId}); webhooks.HttpError(w, err) {
		return
	} else if len(res) == 1 {
		ret = h.Convert.ApplicationDetail(res[0])
	}
	if media.ShouldRender(r) {
		// TODO: partial
	} else if err := media.Res(w, media.Accept(r), ret); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// UpdateApplication implements webhooksv1.ServerInterface.
func (h *Handler) UpdateApplication(w http.ResponseWriter, r *http.Request, applicationId string, params webhooksv1.UpdateApplicationParams) {
	panic("unimplemented")
}
