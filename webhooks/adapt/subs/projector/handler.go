package projector

import (
	"woh/webhooks"

	"github.com/ThreeDotsLabs/watermill/message"
)

type Handler struct {
	Repo webhooks.Repository
}

func (h *Handler) Handle(msg *message.Message) ([]*message.Message, error) {
	return nil, nil
}
