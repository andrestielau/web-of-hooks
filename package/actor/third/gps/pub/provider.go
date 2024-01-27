package pub

import (
	"context"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-googlecloud/pkg/googlecloud"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/andrestielau/web-of-hooks/package/actor"
)

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
