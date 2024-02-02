package dispatcher

import (
	"bytes"
	"context"
	"encoding/base64"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"woh/package/auth"
	"woh/webhooks"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
)

type Handler struct {
	Repo webhooks.Repository
}

func (h *Handler) Handle(msg *message.Message) ([]*message.Message, error) {
	eps, err := h.Repo.GetEndpoints(msg.Context(), []string{msg.Metadata["EndpointID"]})
	if err != nil || len(eps) != 1 {
		return nil, err // TODO: return message with error
	}
	ep := eps[0]
	decoded, err := base64.StdEncoding.DecodeString(ep.Secret)
	if err != nil {
		return nil, err // TODO: return message with error
	}
	id := uuid.NewString()
	now := strconv.Itoa(int(time.Now().Unix()))
	toSign := strings.Join([]string{id, now, string(msg.Payload)}, ".")
	sig := base64.StdEncoding.EncodeToString(auth.HMAC(decoded).Sign([]byte(toSign)))

	ctx, close := context.WithTimeout(msg.Context(), time.Second)
	defer close()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, ep.Url, bytes.NewBuffer(msg.Payload))
	if err != nil {
		return nil, err // TODO: return message with error
	}
	// Add Headers
	req.Header.Set("X-MessageId", id)
	req.Header.Set("X-Timestamp", now)
	req.Header.Set("X-Signature", sig)
	log.Println(id, now, sig)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err // TODO: return message with response and error
	}
	log.Println(res)
	// Return result
	return nil, nil
}
