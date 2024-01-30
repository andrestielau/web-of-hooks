package handle

import (
	"net/http"
	"woh/package/utils/media"
	webhooksv1 "woh/webhooks/adapt/http/v1"

	"github.com/andrestielau/web-of-hooks/webhooks/render/page/messages"
)

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

// DeleteEndpointAttempts implements webhooksv1.ServerInterface.
func (h *Handler) DeleteEndpointAttempts(w http.ResponseWriter, r *http.Request, endpointId string, params webhooksv1.DeleteEndpointAttemptsParams) {
	panic("unimplemented")
}

// DeleteMessageAttempts implements webhooksv1.ServerInterface.
func (h *Handler) DeleteMessageAttempts(w http.ResponseWriter, r *http.Request, messageId string, params webhooksv1.DeleteMessageAttemptsParams) {
	panic("unimplemented")
}

// GetAttempt implements webhooksv1.ServerInterface.
func (h *Handler) GetAttempt(w http.ResponseWriter, r *http.Request, attemptId string) {
	panic("unimplemented")
}

// ListAttempts implements webhooksv1.ServerInterface.
func (h *Handler) ListAttempts(w http.ResponseWriter, r *http.Request, params webhooksv1.ListAttemptsParams) {
	if res, err := h.Repo.ListMessages(r.Context(), 100, 0); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if media.ShouldRender(r) {
		messages.Messages(messages.MessagesViewModel{}).Render(r.Context(), w)
	} else {
		media.Res(w, media.Accept(r), res)
	}
}

// ListEndpointAttempts implements webhooksv1.ServerInterface.
func (h *Handler) ListEndpointAttempts(w http.ResponseWriter, r *http.Request, endpointId string, params webhooksv1.ListEndpointAttemptsParams) {
	panic("unimplemented")
}

// ListMessageAttempts implements webhooksv1.ServerInterface.
func (h *Handler) ListMessageAttempts(w http.ResponseWriter, r *http.Request, messageId string, params webhooksv1.ListMessageAttemptsParams) {
	panic("unimplemented")
}
