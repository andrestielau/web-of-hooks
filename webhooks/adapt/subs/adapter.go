package subs

import (
	"woh/package/actor/third/gps/sub"
	webhooks "woh/webhooks"

	"woh/package/actor"

	"github.com/google/wire"
)

type Options struct {
	Handler *Handler
	sub.Options
}
type Adapter struct {
	*sub.Adapter
}

func New(opts Options) *Adapter {
	a := sub.New(opts.Options)
	a.SpawnAll(actor.Actors{
		webhooks.SecretsKey: opts.Handler.Secrets,
		webhooks.RepoKey:    opts.Handler.Repo,
	})
	return &Adapter{
		Adapter: a,
	}
}

var Set = wire.NewSet(
	wire.Struct(new(Handler), "*"),
	wire.Struct(new(Options), "*"),
	sub.ProvideOptions,
	sub.Set, // TODO: this might fail with more subscribers
	New,
)
