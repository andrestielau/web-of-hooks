package handle

import (
	"net/http"
	"woh/package/utils/media"
	webhooksv1 "woh/webhooks/adapt/http/v1"
	"woh/webhooks/render/page"
)

// ListChannels implements webhooksv1.ServerInterface.
func (h *Handler) ListChannels(w http.ResponseWriter, r *http.Request, params webhooksv1.ListChannelsParams) {
	panic("unimplemented")
}

// ListEventTypes implements webhooksv1.ServerInterface.
func (h *Handler) ListEventTypes(w http.ResponseWriter, r *http.Request, params webhooksv1.ListEventTypesParams) {
	//	h.Repo.CreateEventTypes(r.Context(), []queries.NewEventType{{"test-" + uuid.NewString()}}) uncomment to test types
	eventTypes, err := h.Repo.ListEventTypes(r.Context(), 100, 0)
	if media.ShouldRender(r) {
		page.EventTypes(page.EventTypesViewModel{ // Todo decouple from DB
			Data: eventTypes,
		}, err).Render(r.Context(), w)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	media.Res(w, media.Accept(r), eventTypes)
}
