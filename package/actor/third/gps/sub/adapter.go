package sub

import (
	"context"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-googlecloud/pkg/googlecloud"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/andrestielau/web-of-hooks/package/actor"
	"github.com/google/wire"
)

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
	ProjectID:                "demo",
	GenerateSubscriptionName: googlecloud.TopicSubscriptionName,
}

func ProvideOptions() googlecloud.SubscriberConfig {
	o := DefaultConfig
	return o
}

var Set = wire.NewSet(
	wire.Struct(new(Options), "*"),
	New,
)
