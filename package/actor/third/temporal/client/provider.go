package client

import (
	"context"
	"woh/package/actor"

	"go.temporal.io/sdk/client"
)

const Key = "Temporal"

type Options struct {
	client.Options
}
type Provider struct {
	*actor.Base
	opts Options
	client.Client
}

func New(opts Options) *Provider {
	return &Provider{
		Base: actor.New(),
		opts: opts,
	}
}

func (p *Provider) Start(ctx context.Context) (first bool, err error) {
	if first, err = p.Base.Start(ctx); !first || err != nil {
		return
	}
	p.Client, err = client.Dial(p.opts.Options)
	return
}
func (p *Provider) Stop(ctx context.Context) (last bool, err error) {
	if last, err = p.Base.Stop(ctx); !last || err != nil {
		return
	}
	p.Client.Close()
	return
}
