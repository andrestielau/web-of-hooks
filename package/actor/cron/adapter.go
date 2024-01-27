package cron

import (
	"context"
	"log"
	"time"

	"woh/package/actor"
)

type Options struct {
	Handler func(context.Context)
	Period  time.Duration
}
type Adapter struct {
	*actor.Base
	opts   Options
	t      *time.Ticker
	closer chan struct{}
}

func New(opts Options) *Adapter {
	return &Adapter{
		Base: actor.New(),
		opts: opts,
	}
}

func (a *Adapter) Start(ctx context.Context) (first bool, err error) {
	if first, err = a.Base.Start(ctx); !first || err != nil {
		return first, err
	}
	a.t = time.NewTicker(a.opts.Period)
	go func() {
		defer close(a.closer)
		for range a.t.C {
			if a.run(ctx) {
				break
			}
		}
	}()
	return true, nil
}
func (a *Adapter) run(ctx context.Context) (stop bool) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
			return
		}
		stop = true
	}()
	a.opts.Handler(ctx)
	return stop
}
func (a *Adapter) Stop(ctx context.Context) (last bool, err error) {
	if last, err = a.Base.Start(ctx); !last || err != nil {
		return last, err
	}
	a.t.Stop()
	<-a.closer
	return true, nil
}
