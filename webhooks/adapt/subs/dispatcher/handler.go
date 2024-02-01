package dispatcher

import (
	"bytes"
	"context"
	"log"
	"net/http"
	"time"
	"woh/webhooks"

	"github.com/ThreeDotsLabs/watermill/message"
)

type Handler struct {
	Repo webhooks.Repository
}

func (h *Handler) Handle(msg *message.Message) ([]*message.Message, error) {
	// Get Config
	// Sign Message
	ctx, close := context.WithTimeout(msg.Context(), time.Second)
	defer close()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "", bytes.NewBufferString(""))
	if err != nil {
		return nil, err // TODO: return message with error
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err // TODO: return message with response and error
	}
	log.Println(res)
	// Return result
	return nil, nil
}
