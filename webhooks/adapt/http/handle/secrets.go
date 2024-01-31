package handle

import (
	"fmt"
	"net/http"
	"woh/package/utils/media"

	"woh/webhooks"
	webhooksv1 "woh/webhooks/adapt/http/v1"
	"woh/webhooks/render/page/applications"

	"github.com/samber/lo"
)

// CreateSecret implements webhooksv1.ServerInterface.
func (h *Handler) CreateSecret(w http.ResponseWriter, r *http.Request) {
	var req webhooksv1.CreateSecretsPayload
	var ret webhooksv1.CreatedSecrets
	if err := media.Req(r, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if res, err := h.Repo.CreateSecrets(r.Context(), h.Convert.NewSecrets(req)); err != nil {
		errs, stop := webhooks.Errors(w, err)
		if stop {
			return
		}
		ret.Errors = errs
	} else {
		ret.Data = h.Convert.Secrets(res)
	}
	if media.ShouldRender(r) {
		// TODO: partial
	} else if err := media.Res(w, media.Accept(r), ret); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// DeleteSecret implements webhooksv1.ServerInterface.
func (h *Handler) DeleteSecret(w http.ResponseWriter, r *http.Request, applicationId string, params webhooksv1.DeleteSecretParams) {
	webhooks.Error(w, h.Repo.DeleteSecrets(r.Context(), []string{applicationId}))
}

// DeleteSecrets implements webhooksv1.ServerInterface.
func (h *Handler) DeleteSecrets(w http.ResponseWriter, r *http.Request) {
	var req webhooksv1.DeleteSecretsPayload
	if err := media.Req(r, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	webhooks.Error(w, h.Repo.DeleteSecrets(r.Context(), lo.Map[struct {
		Id *string "json:\"id,omitempty\""
	}](req, func(ids struct {
		Id *string "json:\"id,omitempty\""
	}, _ int) string {
		return *ids.Id
	})))
}

// GetSecret implements webhooksv1.ServerInterface.
func (h *Handler) GetSecret(w http.ResponseWriter, r *http.Request, secretId string) {
	var ret webhooksv1.Secret
	if res, err := h.Repo.GetSecrets(r.Context(), []string{secretId}); webhooks.Error(w, err) {
		return
	} else if len(res) == 0 {
		http.Error(w, fmt.Sprintf("Secret with uid %s not found", secretId), http.StatusNotFound)
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

// ListSecrets implements webhooksv1.ServerInterface.
func (h *Handler) ListSecrets(w http.ResponseWriter, r *http.Request, params webhooksv1.ListSecretsParams) {
	if params.Limit == nil {
		params.Limit = lo.ToPtr(20)
	}
	if res, err := h.Repo.ListSecrets(r.Context(), h.Convert.SecretQuery(params)); err != nil {
		webhooks.Error(w, err)
	} else if media.ShouldRender(r) {
		applications.Secrets(applications.SecretViewModel{
			Data: res,
		}, nil).Render(r.Context(), w)
	} else {
		media.Res(w, media.Accept(r), res)
	}
}

// RotateSecret implements webhooksv1.ServerInterface.
func (h *Handler) RotateSecret(w http.ResponseWriter, r *http.Request, secretId string) {
	panic("unimplemented")
}


// CreateSecret implements webhooksv1.ServerInterface.
func (h *Handler) CreateApplicationSecrets(w http.ResponseWriter, r *http.Request, applicationId string) {
	panic("unimplemented")
}

// ListApplicationSecrets implements webhooksv1.ServerInterface.
func (h *Handler) ListApplicationSecrets(w http.ResponseWriter, r *http.Request, applicationId string) {
	panic("unimplemented")
}
