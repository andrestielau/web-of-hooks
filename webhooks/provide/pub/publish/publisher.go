package publish

import (
	"encoding/json"
	"woh/package/actor/third/gps/pub"

	"woh/webhooks"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

type Publisher struct {
	Pub *pub.Provider
}

func (p *Publisher) Publish(payload webhooks.Payload) error {
	endpoints, err := p.SearchEndpoints(payload)
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

func (p *Publisher) SearchEndpoints(payload webhooks.Payload) ([]webhooks.Endpoint, error) {
	//TODO implement
	return nil, nil
}
