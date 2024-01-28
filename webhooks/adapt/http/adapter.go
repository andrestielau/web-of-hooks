package http

import (
	"woh/package/actor"
	"woh/package/actor/net/http/server"
	"woh/webhooks"
	"woh/webhooks/adapt/http/handle"
	webhooksv1 "woh/webhooks/adapt/http/v1"

	"github.com/alexedwards/scs/v2"
	"github.com/google/wire"
)

type Options struct {
	Handler *handle.Handler
}
type Adapter struct {
	*server.Adapter
}

func New(opts Options) *Adapter {
	a := server.New(server.Options{
		Handler: opts.Handler.Session.LoadAndSave( // wrap in session middleware
			webhooksv1.Handler(opts.Handler),
		),
		Addr: ":3000",
	})
	a.SpawnAll(actor.Actors{
		webhooks.SecretsKey: opts.Handler.Secrets,
		webhooks.RepoKey:    opts.Handler.Repo,
	})
	return &Adapter{
		Adapter: a,
	}
}

var Set = wire.NewSet(
	wire.Struct(new(handle.Handler), "*"),
	wire.Struct(new(Options), "*"),
	scs.New,
	New,
)
