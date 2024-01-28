package work

import (
	"woh/package/actor"
	"woh/package/actor/third/temporal/client"
	"woh/package/actor/third/temporal/worker"
	"woh/webhooks"

	worker2 "go.temporal.io/sdk/worker"

	"github.com/google/wire"
)

type Options struct {
	Provider *client.Provider
	Handler  *Handler
}
type Adapter struct {
	*worker.Adapter
}

func New(opts Options) *Adapter {
	a := worker.New(worker.Options{
		Provider: opts.Provider,
		Handler: func(w worker2.Worker) {
			// Register workflows and activities
		},
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
