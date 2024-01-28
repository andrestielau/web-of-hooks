package cron

import (
	"woh/package/actor/cron"
	webhooks "woh/webhooks"

	"github.com/google/wire"
)

type Adapter struct {
	*cron.Adapter
}

func New(w webhooks.Worker) *Adapter {
	a := cron.New(cron.Options{})
	a.Spawn(webhooks.WorkerKey, w)
	return &Adapter{
		Adapter: a,
	}
}

var Set = wire.NewSet(
	New,
)
