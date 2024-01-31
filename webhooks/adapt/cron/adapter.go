package cron

import (
	"context"
	"log"
	"time"
	"woh/package/actor"
	"woh/package/actor/cron"
	webhooks "woh/webhooks"

	"github.com/google/uuid"
	"github.com/google/wire"
)

type Options struct {
	Handler *Handler
}
type Adapter struct {
	*cron.Adapter
	Handler *Handler
}

func New(opts Options) *Adapter {
	a := cron.New(cron.Options{
		Handler: opts.Handler.Work,
		Period:  time.Second, // TODO: make provider
	})
	a.SpawnAll(actor.Actors{
		webhooks.RepoKey: opts.Handler.Repo,
	})
	return &Adapter{
		Adapter: a,
		Handler: opts.Handler,
	}
}

func (a *Adapter) Start(ctx context.Context) (first bool, err error) {
	if first, err = a.Base.Start(ctx); !first || err != nil {
		return
	}
	if err = a.Handler.Repo.SetLastSeen(ctx, string(a.Handler.Id)); err != nil {
		return
	}
	time.Sleep(10 * time.Second) // Let other workers drain their apps messages
	log.Println("Donde está lá biblioteca?")
	a.Run(ctx)
	return
}

type WorkerId string

func ProvideWorkerId() WorkerId {
	return WorkerId(uuid.NewString())
}

var Set = wire.NewSet(
	ProvideWorkerId,
	wire.Struct(new(Handler), "*"),
	wire.Struct(new(Options), "*"),
	New,
)
