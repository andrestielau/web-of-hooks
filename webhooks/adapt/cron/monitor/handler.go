package monitor

import (
	"context"

	"woh/webhooks"

	"github.com/ThreeDotsLabs/watermill/message"
)

type Handler struct {
	Publisher message.Publisher
	Repo      webhooks.Repository
}

func (h *Handler) Run(context.Context) {

}
