package http

import (
	"net/http"
	"strings"

	"woh/package/utils"
	"woh/webhooks"
	"woh/webhooks/adapt/html/page"
	webhooksv1 "woh/webhooks/adapt/http/v1"

	"woh/webhooks/adapt/html/style/theme"

	"github.com/alexedwards/scs/v2"
)

type Handler struct {
	webhooks.Manager
}

var sessionManager *scs.SessionManager
var _ webhooksv1.ServerInterface = &Handler{}

// CreateAttempt implements webhooksv1.ServerInterface.
func (h *Handler) CreateAttempt(w http.ResponseWriter, r *http.Request, attemptId string, params webhooksv1.CreateAttemptParams) {
	panic("unimplemented")
}

// CreateAttempts implements webhooksv1.ServerInterface.
func (h *Handler) CreateAttempts(w http.ResponseWriter, r *http.Request, params webhooksv1.CreateAttemptsParams) {
	panic("unimplemented")
}

// CreateEndpointAttempts implements webhooksv1.ServerInterface.
func (h *Handler) CreateEndpointAttempts(w http.ResponseWriter, r *http.Request, endpointId string, params webhooksv1.CreateEndpointAttemptsParams) {
	panic("unimplemented")
}

// CreateEndpoints implements webhooksv1.ServerInterface.
func (h *Handler) CreateEndpoints(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// CreateMessage implements webhooksv1.ServerInterface.
func (h *Handler) CreateMessage(w http.ResponseWriter, r *http.Request, messageId string, params webhooksv1.CreateMessageParams) {
	panic("unimplemented")
}

// CreateMessages implements webhooksv1.ServerInterface.
func (h *Handler) CreateMessages(w http.ResponseWriter, r *http.Request, params webhooksv1.CreateMessagesParams) {
	panic("unimplemented")
}

// CreateMessagesAttempts implements webhooksv1.ServerInterface.
func (h *Handler) CreateMessagesAttempts(w http.ResponseWriter, r *http.Request, messageId string, params webhooksv1.CreateMessagesAttemptsParams) {
	panic("unimplemented")
}

// DeleteAttempt implements webhooksv1.ServerInterface.
func (h *Handler) DeleteAttempt(w http.ResponseWriter, r *http.Request, attemptId string, params webhooksv1.DeleteAttemptParams) {
	panic("unimplemented")
}

// DeleteAttempts implements webhooksv1.ServerInterface.
func (h *Handler) DeleteAttempts(w http.ResponseWriter, r *http.Request, params webhooksv1.DeleteAttemptsParams) {
	panic("unimplemented")
}

// DeleteEndpoint implements webhooksv1.ServerInterface.
func (h *Handler) DeleteEndpoint(w http.ResponseWriter, r *http.Request, endpointId string, params webhooksv1.DeleteEndpointParams) {
	panic("unimplemented")
}

// DeleteEndpointAttempts implements webhooksv1.ServerInterface.
func (h *Handler) DeleteEndpointAttempts(w http.ResponseWriter, r *http.Request, endpointId string, params webhooksv1.DeleteEndpointAttemptsParams) {
	panic("unimplemented")
}

// DeleteMessage implements webhooksv1.ServerInterface.
func (h *Handler) DeleteMessage(w http.ResponseWriter, r *http.Request, messageId string, params webhooksv1.DeleteMessageParams) {
	panic("unimplemented")
}

// DeleteMessageAttempts implements webhooksv1.ServerInterface.
func (h *Handler) DeleteMessageAttempts(w http.ResponseWriter, r *http.Request, messageId string, params webhooksv1.DeleteMessageAttemptsParams) {
	panic("unimplemented")
}

// DeleteMessages implements webhooksv1.ServerInterface.
func (h *Handler) DeleteMessages(w http.ResponseWriter, r *http.Request, params webhooksv1.DeleteMessagesParams) {
	panic("unimplemented")
}

// DisableEndpoints implements webhooksv1.ServerInterface.
func (h *Handler) DisableEndpoints(w http.ResponseWriter, r *http.Request, params webhooksv1.DisableEndpointsParams) {
	panic("unimplemented")
}

// GetAttempt implements webhooksv1.ServerInterface.
func (h *Handler) GetAttempt(w http.ResponseWriter, r *http.Request, attemptId string) {
	panic("unimplemented")
}

// GetEndpoint implements webhooksv1.ServerInterface.
func (h *Handler) GetEndpoint(w http.ResponseWriter, r *http.Request, endpointId string) {
	panic("unimplemented")
}

// GetEndpointSecret implements webhooksv1.ServerInterface.
func (h *Handler) GetEndpointSecret(w http.ResponseWriter, r *http.Request, endpointId string) {
	panic("unimplemented")
}

// GetEndpointStats implements webhooksv1.ServerInterface.
func (h *Handler) GetEndpointStats(w http.ResponseWriter, r *http.Request, endpointId string) {
	panic("unimplemented")
}

const status = "ok"

// GetHealth implements webhooksv1.ServerInterface.
func (h *Handler) GetHealth(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.Header.Get("Accept"), "text/html") {
		currentCount := sessionManager.GetInt(r.Context(), "count")
		sessionManager.Put(r.Context(), "count", currentCount+1)
		t := "dark"
		if currentCount%2 == 1 {
			t = "light"
		}
		page.Health(status).Render(theme.Set(r.Context(), t), w)
		return
	}
	utils.JsonRes(w, map[string]any{"status": status})
}

// GetMessage implements webhooksv1.ServerInterface.
func (h *Handler) GetMessage(w http.ResponseWriter, r *http.Request, messageId string) {
	panic("unimplemented")
}

// GetStats implements webhooksv1.ServerInterface.
func (h *Handler) GetStats(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// Jwks implements webhooksv1.ServerInterface.
func (h *Handler) Jwks(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// ListAttempts implements webhooksv1.ServerInterface.
func (h *Handler) ListAttempts(w http.ResponseWriter, r *http.Request, params webhooksv1.ListAttemptsParams) {
	panic("unimplemented")
}

// ListEndpointAttempr implements webhooksv1.ServerInterface.
func (h *Handler) ListEndpointAttempr(w http.ResponseWriter, r *http.Request, endpointId string, params webhooksv1.ListEndpointAttemprParams) {
	panic("unimplemented")
}

// ListEndpoints implements webhooksv1.ServerInterface.
func (h *Handler) ListEndpoints(w http.ResponseWriter, r *http.Request, params webhooksv1.ListEndpointsParams) {
	panic("unimplemented")
}

// ListEventTypes implements webhooksv1.ServerInterface.
func (h *Handler) ListEventTypes(w http.ResponseWriter, r *http.Request, params webhooksv1.ListEventTypesParams) {
	//h.Manager.Repo().ListEventTypes(r.Context())
	if strings.Contains(r.Header.Get("Accept"), "text/html") {
		page.EventTypes(page.EventTypesView()).Render(r.Context(), w)
		return
	}
	utils.JsonRes(w, []string{})
}

// ListMessageAttempts implements webhooksv1.ServerInterface.
func (h *Handler) ListMessageAttempts(w http.ResponseWriter, r *http.Request, messageId string, params webhooksv1.ListMessageAttemptsParams) {
	panic("unimplemented")
}

// ListMessages implements webhooksv1.ServerInterface.
func (h *Handler) ListMessages(w http.ResponseWriter, r *http.Request, params webhooksv1.ListMessagesParams) {
	panic("unimplemented")
}

// RotateEndpointSecret implements webhooksv1.ServerInterface.
func (h *Handler) RotateEndpointSecret(w http.ResponseWriter, r *http.Request, endpointId string) {
	panic("unimplemented")
}

// UpdateEndpoint implements webhooksv1.ServerInterface.
func (h *Handler) UpdateEndpoint(w http.ResponseWriter, r *http.Request, endpointId string) {
	panic("unimplemented")
}
