package cron

import (
	"context"
	"log"
	"sync"
	"time"

	"woh/package/actor"
)

type HandlerOptions struct {
	Func   func(context.Context)
	Period time.Duration
}
type Adapter struct {
	*actor.Base
	wg   sync.WaitGroup
	t    map[string]*time.Ticker
	opts map[string]HandlerOptions
}

func New() *Adapter {
	return &Adapter{
		Base: actor.New(),
		t:    make(map[string]*time.Ticker),
		opts: make(map[string]HandlerOptions),
	}
}
func (a *Adapter) Start(ctx context.Context) (first bool, err error) {
	if first, err = a.Base.Start(ctx); !first || err != nil {
		return first, err
	}
	a.wg.Add(len(a.opts))
	for k, v := range a.opts {
		t := time.NewTicker(v.Period)
		go a.Run(ctx, t, v.Func)
		a.t[k] = t
	}
	return true, nil
}
func (a *Adapter) Stop(ctx context.Context) (last bool, err error) {
	if last, err = a.Base.Start(ctx); !last || err != nil {
		return last, err
	}
	for k, v := range a.t {
		delete(a.t, k)
		v.Stop()
	}
	return true, nil
}
func (a *Adapter) Add(k string, v HandlerOptions) {
	a.Lock()
	defer a.Unlock()
	a.opts[k] = v
}
func (a *Adapter) AddAll(opts map[string]HandlerOptions) {
	a.Lock()
	defer a.Unlock()
	for k, opt := range opts {
		a.opts[k] = opt
	}
}
func (a *Adapter) Run(ctx context.Context, t *time.Ticker, fn func(context.Context)) {
	defer a.wg.Done()
	for range t.C {
		if a.run(ctx, fn) {
			break
		}
	}
}

func (a *Adapter) run(ctx context.Context, fn func(context.Context)) (stop bool) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
			stop = true
			return
		}
	}()
	fn(ctx) // TODO: should the should stop come from the handler?
	return stop
}
