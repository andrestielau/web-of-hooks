package publish

import (
	"context"
	"encoding/json"
	"woh/package/actor/third/gps/pub"

	"woh/webhooks"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/samber/lo"
)

type Publisher struct {
	Pub  *pub.Provider
	Repo webhooks.Repository
}

func (p *Publisher) Publish(ctx context.Context, payload webhooks.Payload) error {
	endpoints, err := p.SearchEndpoints(ctx, payload)
	if err != nil {
		return err
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	msg := message.NewMessage(
		watermill.NewUUID(), jsonPayload,
	)
	for _, endpoint := range endpoints {
		msg.Metadata = map[string]string{
			"EndpointID": endpoint.Uid,
		}
		err = p.Pub.Publish("dispatcher", msg)
		if err != nil {
			return err
		}
	}
	return err
}

func (p *Publisher) SearchEndpoints(ctx context.Context, payload webhooks.Payload) ([]webhooks.Endpoint, error) {
	endpoints, err := p.Repo.GetEndpointsByTenantAndEventTypes(ctx, payload.TenantID, []string{payload.EventTypeKey})
	if err != nil {
		return nil, err
	}
	return lo.Map(endpoints, func(e webhooks.EndpointDetails, _ int) webhooks.Endpoint {
		return e.Endpoint
	}), nil
	// return []webhooks.Endpoint{{Uid: "018d6907-fb08-7227-537f-f32dc3364eec"}}, nil
}
