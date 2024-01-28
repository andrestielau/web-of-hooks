package handle

import (
	"net/http"
	webhooksv1 "woh/webhooks/adapt/http/v1"
)

// CreateMessage implements webhooksv1.ServerInterface.
func (h *Handler) CreateMessage(w http.ResponseWriter, r *http.Request, messageId string, params webhooksv1.CreateMessageParams) {
	panic("unimplemented")
}

// CreateMessages implements webhooksv1.ServerInterface.
func (h *Handler) CreateMessages(w http.ResponseWriter, r *http.Request, applicationId string, params webhooksv1.CreateMessagesParams) {
	panic("unimplemented")
}

// DeleteMessage implements webhooksv1.ServerInterface.
func (h *Handler) DeleteMessage(w http.ResponseWriter, r *http.Request, messageId string, params webhooksv1.DeleteMessageParams) {
	panic("unimplemented")
}

// DeleteMessages implements webhooksv1.ServerInterface.
func (h *Handler) DeleteMessages(w http.ResponseWriter, r *http.Request, applicationId string, params webhooksv1.DeleteMessagesParams) {
	panic("unimplemented")
}

// GetMessage implements webhooksv1.ServerInterface.
func (h *Handler) GetMessage(w http.ResponseWriter, r *http.Request, messageId string) {
	panic("unimplemented")
}

// ListMessages implements webhooksv1.ServerInterface.
func (h *Handler) ListMessages(w http.ResponseWriter, r *http.Request, applicationId string, params webhooksv1.ListMessagesParams) {
	panic("unimplemented")
}
