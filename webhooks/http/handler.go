package http

import (
	"net/http"
	"strings"

	"github.com/andrestielau/web-of-hooks/internal/domain"
	"github.com/andrestielau/web-of-hooks/package/utils"
	webhooksv1 "github.com/andrestielau/web-of-hooks/webhooks/http/v1"
	"github.com/gomarkdown/markdown"
)

type Handler struct {
	domain.Manager
}

var _ webhooksv1.ServerInterface = &Handler{}

// CreateAttempt implements webhooksv1.ServerInterface.
func (*Handler) CreateAttempt(w http.ResponseWriter, r *http.Request, attemptId string, params webhooksv1.CreateAttemptParams) {
	panic("unimplemented")
}

// CreateAttempts implements webhooksv1.ServerInterface.
func (*Handler) CreateAttempts(w http.ResponseWriter, r *http.Request, params webhooksv1.CreateAttemptsParams) {
	panic("unimplemented")
}

// CreateEndpointAttempts implements webhooksv1.ServerInterface.
func (*Handler) CreateEndpointAttempts(w http.ResponseWriter, r *http.Request, endpointId string, params webhooksv1.CreateEndpointAttemptsParams) {
	panic("unimplemented")
}

// CreateEndpoints implements webhooksv1.ServerInterface.
func (*Handler) CreateEndpoints(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// CreateMessage implements webhooksv1.ServerInterface.
func (*Handler) CreateMessage(w http.ResponseWriter, r *http.Request, messageId string, params webhooksv1.CreateMessageParams) {
	panic("unimplemented")
}

// CreateMessages implements webhooksv1.ServerInterface.
func (*Handler) CreateMessages(w http.ResponseWriter, r *http.Request, params webhooksv1.CreateMessagesParams) {
	panic("unimplemented")
}

// CreateMessagesAttempts implements webhooksv1.ServerInterface.
func (*Handler) CreateMessagesAttempts(w http.ResponseWriter, r *http.Request, messageId string, params webhooksv1.CreateMessagesAttemptsParams) {
	panic("unimplemented")
}

// DeleteAttempt implements webhooksv1.ServerInterface.
func (*Handler) DeleteAttempt(w http.ResponseWriter, r *http.Request, attemptId string, params webhooksv1.DeleteAttemptParams) {
	panic("unimplemented")
}

// DeleteAttempts implements webhooksv1.ServerInterface.
func (*Handler) DeleteAttempts(w http.ResponseWriter, r *http.Request, params webhooksv1.DeleteAttemptsParams) {
	panic("unimplemented")
}

// DeleteEndpoint implements webhooksv1.ServerInterface.
func (*Handler) DeleteEndpoint(w http.ResponseWriter, r *http.Request, endpointId string, params webhooksv1.DeleteEndpointParams) {
	panic("unimplemented")
}

// DeleteEndpointAttempts implements webhooksv1.ServerInterface.
func (*Handler) DeleteEndpointAttempts(w http.ResponseWriter, r *http.Request, endpointId string, params webhooksv1.DeleteEndpointAttemptsParams) {
	panic("unimplemented")
}

// DeleteMessage implements webhooksv1.ServerInterface.
func (*Handler) DeleteMessage(w http.ResponseWriter, r *http.Request, messageId string, params webhooksv1.DeleteMessageParams) {
	panic("unimplemented")
}

// DeleteMessageAttempts implements webhooksv1.ServerInterface.
func (*Handler) DeleteMessageAttempts(w http.ResponseWriter, r *http.Request, messageId string, params webhooksv1.DeleteMessageAttemptsParams) {
	panic("unimplemented")
}

// DeleteMessages implements webhooksv1.ServerInterface.
func (*Handler) DeleteMessages(w http.ResponseWriter, r *http.Request, params webhooksv1.DeleteMessagesParams) {
	panic("unimplemented")
}

// DisableEndpoints implements webhooksv1.ServerInterface.
func (*Handler) DisableEndpoints(w http.ResponseWriter, r *http.Request, params webhooksv1.DisableEndpointsParams) {
	panic("unimplemented")
}

// GetAttempt implements webhooksv1.ServerInterface.
func (*Handler) GetAttempt(w http.ResponseWriter, r *http.Request, attemptId string) {
	panic("unimplemented")
}

// GetEndpoint implements webhooksv1.ServerInterface.
func (*Handler) GetEndpoint(w http.ResponseWriter, r *http.Request, endpointId string) {
	panic("unimplemented")
}

// GetEndpointSecret implements webhooksv1.ServerInterface.
func (*Handler) GetEndpointSecret(w http.ResponseWriter, r *http.Request, endpointId string) {
	panic("unimplemented")
}

// GetEndpointStats implements webhooksv1.ServerInterface.
func (*Handler) GetEndpointStats(w http.ResponseWriter, r *http.Request, endpointId string) {
	panic("unimplemented")
}

const status = "ok"

// GetHealth implements webhooksv1.ServerInterface.
func (*Handler) GetHealth(w http.ResponseWriter, r *http.Request) {
	// If negotiator returns html render health page
	if strings.Contains(r.Header.Get("Accept"), "application/html") {
		w.WriteHeader(200)
		w.Header().Add("Content-Type", "application/html")
		w.Write(markdown.ToHTML([]byte(`# Health
	## Status: `+status+`
	`), nil, nil))
		return
	}
	utils.JsonRes(w, map[string]any{"status": status})
}

// GetMessage implements webhooksv1.ServerInterface.
func (*Handler) GetMessage(w http.ResponseWriter, r *http.Request, messageId string) {
	panic("unimplemented")
}

// GetStats implements webhooksv1.ServerInterface.
func (*Handler) GetStats(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// Jwks implements webhooksv1.ServerInterface.
func (*Handler) Jwks(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// ListAttempts implements webhooksv1.ServerInterface.
func (*Handler) ListAttempts(w http.ResponseWriter, r *http.Request, params webhooksv1.ListAttemptsParams) {
	panic("unimplemented")
}

// ListEndpointAttempr implements webhooksv1.ServerInterface.
func (*Handler) ListEndpointAttempr(w http.ResponseWriter, r *http.Request, endpointId string, params webhooksv1.ListEndpointAttemprParams) {
	panic("unimplemented")
}

// ListEndpoints implements webhooksv1.ServerInterface.
func (*Handler) ListEndpoints(w http.ResponseWriter, r *http.Request, params webhooksv1.ListEndpointsParams) {
	panic("unimplemented")
}

// ListEventTypes implements webhooksv1.ServerInterface.
func (*Handler) ListEventTypes(w http.ResponseWriter, r *http.Request, params webhooksv1.ListEventTypesParams) {
	panic("unimplemented")
}

// ListMessageAttempts implements webhooksv1.ServerInterface.
func (*Handler) ListMessageAttempts(w http.ResponseWriter, r *http.Request, messageId string, params webhooksv1.ListMessageAttemptsParams) {
	panic("unimplemented")
}

// ListMessages implements webhooksv1.ServerInterface.
func (*Handler) ListMessages(w http.ResponseWriter, r *http.Request, params webhooksv1.ListMessagesParams) {
	panic("unimplemented")
}

// RotateEndpointSecret implements webhooksv1.ServerInterface.
func (*Handler) RotateEndpointSecret(w http.ResponseWriter, r *http.Request, endpointId string) {
	panic("unimplemented")
}

// UpdateEndpoint implements webhooksv1.ServerInterface.
func (*Handler) UpdateEndpoint(w http.ResponseWriter, r *http.Request, endpointId string) {
	panic("unimplemented")
}
