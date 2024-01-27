//go:build wireinject
// +build wireinject

package worker

import (
	"github.com/andrestielau/web-of-hooks/internal/domain/worker"
	"github.com/andrestielau/web-of-hooks/internal/provide/repo"
	"github.com/andrestielau/web-of-hooks/internal/provide/secrets"
	"github.com/andrestielau/web-of-hooks/package/actor"
	"github.com/andrestielau/web-of-hooks/webhooks/cron"
	"github.com/google/wire"
)

func Adapters() actor.Actors {
	panic(wire.Build(
		// Providers
		secrets.Set,
		repo.Set,
		// Services
		worker.Set,
		// Adapters
		cron.Set,
		ChooseAdapters,
	))
}
