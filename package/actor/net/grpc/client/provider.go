package client

import (
	"context"

	"woh/package/actor"

	"google.golang.org/grpc"
)

type Options struct {
	Addr string
}
type Provider struct {
	*actor.Base
	opts Options
	*grpc.ClientConn
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
	a.ClientConn, err = grpc.Dial(a.opts.Addr)
	return
}

func (a *Provider) Stop(ctx context.Context) (last bool, err error) {
	if last, err = a.Base.Stop(ctx); !last || err != nil {
		return
	}
	err = a.ClientConn.Close()
	return
}
