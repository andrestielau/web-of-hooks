//go:build wireinject
// +build wireinject

package worker

import (
	"woh/internal/domain/worker"
	"woh/internal/provide/repo"
	"woh/internal/provide/secrets"
	"woh/package/actor"
	"woh/webhooks/cron"

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
