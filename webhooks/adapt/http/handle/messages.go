package handle

import (
	"net/http"
	"woh/package/utils/media"
	webhooksv1 "woh/webhooks/adapt/http/v1"
	"woh/webhooks/render/page/messages"

	"github.com/samber/lo"
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
	if params.Limit == nil {
		params.Limit = lo.ToPtr(LIMIT_DEFAULT)
	}
	if res, err := h.Repo.ListApplicationMessages(r.Context(), applicationId, h.Convert.MessageQuery(params)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if media.ShouldRender(r) {
		messages.Messages(messages.MessagesViewModel{ // Todo decouple from DB
			// ApplicationId: applicationId,
			// Data:          nil,
		}).Render(r.Context(), w)
	} else {
		media.Res(w, media.Accept(r), res)
	}
}
