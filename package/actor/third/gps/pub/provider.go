package pub

import (
	"context"

	"woh/package/actor"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-googlecloud/pkg/googlecloud"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/wire"
)

const Key = "Publisher"

type Options struct {
	Config googlecloud.PublisherConfig
	Logger watermill.LoggerAdapter
}
type Provider struct {
	*actor.Base
	opts Options
	message.Publisher
}

func New(opts Options) *Provider {
	return &Provider{
		Base: actor.New(),
		opts: opts,
	}
}

func (a *Provider) Start(ctx context.Context) (first bool, err error) {
	if first, err = a.Base.Start(ctx); !first || err != nil {
		return
	}
	a.Publisher, err = googlecloud.NewPublisher(a.opts.Config, a.opts.Logger)
	return
}

func (a *Provider) Stop(ctx context.Context) (last bool, err error) {
	if last, err = a.Base.Stop(ctx); !last || err != nil {
		return
	}
	if a.Publisher != nil {
		err = a.Publisher.Close()
	}
	return
}

var DefaultConfig = googlecloud.PublisherConfig{
	ProjectID: "demo", // TODO: change this before deploy-pr
}

func ProvideConfig() googlecloud.PublisherConfig {
	o := DefaultConfig
	return o
}

var Set = wire.NewSet(
	wire.Bind(new(message.Publisher), new(*Provider)), // TODO: generalize this
	wire.Struct(new(Options), "*"),
	ProvideConfig,
	New,
)
