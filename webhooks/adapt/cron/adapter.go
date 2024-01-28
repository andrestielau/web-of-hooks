package cron

import (
	"time"
	"woh/package/actor"
	"woh/package/actor/cron"
	webhooks "woh/webhooks"

	"github.com/google/wire"
)

type Options struct {
	Handler *Handler
}
type Adapter struct {
	*actor.Base
}

func New(opts Options) *Adapter {
	a := actor.New()
	deps := actor.Actors{
		webhooks.SecretsKey: opts.Handler.Secrets,
		webhooks.RepoKey:    opts.Handler.Repo,
	}
	a.SpawnAll(actor.Actors{
		"Send": cron.New(cron.Options{
			Handler: opts.Handler.Send,
			Period:  time.Minute,
		}).SpawnAll(deps),
		"Sign": cron.New(cron.Options{
			Handler: opts.Handler.Sign,
			Period:  time.Minute,
		}).SpawnAll(deps),
	})
	return &Adapter{Base: a}
}

var Set = wire.NewSet(
	wire.Struct(new(Handler), "*"),
	wire.Struct(new(Options), "*"),
	New,
)
