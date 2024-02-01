package subs

import (
	"woh/package/actor"
	"woh/package/actor/router"
	"woh/package/actor/third/gps/pub"
	"woh/package/actor/third/gps/sub"
	webhooks "woh/webhooks"
	"woh/webhooks/adapt/subs/dispatcher"
	"woh/webhooks/adapt/subs/projector"

	"github.com/google/wire"
)

type Options struct {
	Dispatcher dispatcher.Handler
	Projector  projector.Handler
	Repo       webhooks.Repository // This should be in childrent but that's too much work for now
	Pub        *pub.Provider
	Sub        *sub.Adapter
	*router.Adapter
}
type Adapter struct {
	*router.Adapter
}

func New(opts Options) *Adapter {
	opts.Adapter.Handle(map[string]router.HandlerOptions{
		"dispatcher": {In: "dispatcher", Sub: opts.Sub, Out: "dispatcher", Pub: opts.Pub, Func: opts.Dispatcher.Handle},
		"projector":  {In: "projector", Sub: opts.Sub, Out: "projector", Pub: opts.Pub, Func: opts.Projector.Handle},
	})
	opts.Adapter.SpawnAll(actor.Actors{
		webhooks.RepoKey: opts.Repo,
		pub.Key:          opts.Pub,
		sub.Key:          opts.Sub,
	})
	return &Adapter{Adapter: opts.Adapter}
}

var Set = wire.NewSet(
	wire.Struct(new(dispatcher.Handler), "*"),
	wire.Struct(new(projector.Handler), "*"),
	wire.Struct(new(Options), "*"),
	New,
)
