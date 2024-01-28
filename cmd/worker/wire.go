//go:build wireinject
// +build wireinject

package worker

import (
	"woh/package/actor"
	"woh/webhooks/adapt/cron"
	"woh/webhooks/provide/repo"
	"woh/webhooks/provide/secrets"
	"woh/webhooks/service/worker"

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
