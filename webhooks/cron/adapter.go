package cron

import (
	"github.com/andrestielau/web-of-hooks/internal/domain"
	"github.com/andrestielau/web-of-hooks/package/actor/cron"
	"github.com/google/wire"
)

type Adapter struct {
	*cron.Adapter
}

func New(w domain.Worker) *Adapter {
	a := cron.New(cron.Options{})
	a.Spawn(domain.WorkerKey, w)
	return &Adapter{
		Adapter: a,
	}
}

var Set = wire.NewSet(
	New,
)
