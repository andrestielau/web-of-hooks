package cron

import (
	"time"
	webhooks "woh/webhooks"

	"woh/package/actor/cron"
	"woh/webhooks/adapt/cron/monitor"

	"github.com/google/wire"
)

type Options struct {
	Monitor *monitor.Handler
}
type Adapter struct {
	*cron.Adapter
}

func New(opts Options) *Adapter {
	a := cron.New()
	a.Add("monitor", cron.HandlerOptions{
		Period: time.Second,
		Func:   opts.Monitor.Run,
	})
	a.Spawn(webhooks.RepoKey, opts.Monitor.Repo)
	return &Adapter{a}
}

var Set = wire.NewSet(
	wire.Struct(new(monitor.Handler), "*"),
	wire.Struct(new(Options), "*"),
	New,
)
