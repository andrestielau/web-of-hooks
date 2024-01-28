package http

import (
	"woh/package/actor"
	"woh/package/actor/net/http/server"
	"woh/webhooks"
	webhooksv1 "woh/webhooks/adapt/http/v1"

	"github.com/alexedwards/scs/v2"
	"github.com/google/wire"
)

type Options struct {
	Handler *Handler
}
type Adapter struct {
	*server.Adapter
}

func New(opts Options) *Adapter {
	sessionManager = scs.New()
	a := server.New(server.Options{
		Handler: sessionManager.LoadAndSave(webhooksv1.Handler(opts.Handler)),
		Addr:    ":3000",
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
	wire.Struct(new(Handler), "*"),
	wire.Struct(new(Options), "*"),
	New,
)
