package client

import (
	"context"
	"net/http"

	"woh/package/actor"
)

type Options struct {
}
type Provider struct {
	*actor.Base
	opts Options
	*http.Client
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
	a.Client = &http.Client{}
	return
}
