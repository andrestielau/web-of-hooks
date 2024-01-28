package worker

import (
	"context"
	"woh/package/actor"
	webhooksv1 "woh/webhooks/adapt/work/v1"

	"woh/package/actor/third/temporal/client"

	"go.temporal.io/sdk/worker"
)

type Options struct {
	Provider *client.Provider
	worker.Options
	Handler func(worker.Worker)
}
type Adapter struct {
	*actor.Base
	opts Options
	worker.Worker
}

func New(opts Options) *Adapter {
	a := actor.New()
	a.Spawn(client.Key, opts.Provider)
	return &Adapter{
		Base: a,
		opts: opts,
	}
}

func (a *Adapter) Start(ctx context.Context) (first bool, err error) {
	if first, err = a.Base.Start(ctx); !first || err != nil {
		return
	}
	a.Worker = worker.New(a.opts.Provider.Client, webhooksv1.Queue, a.opts.Options)
	if a.opts.Handler == nil {
		return
	}
	a.opts.Handler(a.Worker)
	return
}
func (a *Adapter) Stop(ctx context.Context) (last bool, err error) {
	if last, err = a.Base.Stop(ctx); !last || err != nil {
		return
	}
	a.Worker.Stop()
	return
}
