package sub

import (
	"context"

	"woh/package/actor"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-googlecloud/pkg/googlecloud"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/wire"
)

const Key = "Subscriber"

type Options struct {
	Config googlecloud.SubscriberConfig
	Logger watermill.LoggerAdapter
}
type Adapter struct {
	*actor.Base
	opts Options
	message.Subscriber
}

func New(opts Options) *Adapter {
	return &Adapter{
		Base: actor.New(),
		opts: opts,
	}
}

func (a *Adapter) Start(ctx context.Context) (first bool, err error) {
	if first, err = a.Base.Start(ctx); !first || err != nil {
		return
	}
	a.Subscriber, err = googlecloud.NewSubscriber(a.opts.Config, a.opts.Logger)
	return
}

func (a *Adapter) Stop(ctx context.Context) (last bool, err error) {
	if last, err = a.Base.Stop(ctx); !last || err != nil {
		return
	}
	if a.Subscriber != nil {
		err = a.Subscriber.Close()
	}
	return
}

var DefaultConfig = googlecloud.SubscriberConfig{
	ProjectID:                "demo", // TODO: change this before deploy-pr
	GenerateSubscriptionName: googlecloud.TopicSubscriptionName,
}

func ProvideConfig() googlecloud.SubscriberConfig {
	o := DefaultConfig
	return o
}

var Set = wire.NewSet(
	wire.Bind(new(message.Subscriber), new(*Adapter)), // TODO: generalize this
	wire.Struct(new(Options), "*"),
	ProvideConfig,
	New,
)
