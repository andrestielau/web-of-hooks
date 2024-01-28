//go:build wireinject
// +build wireinject

package worker

import (
	"woh/package/actor"
	"woh/webhooks/adapt/cron"
	"woh/webhooks/provide/repo"
	"woh/webhooks/provide/secrets"

	"github.com/google/wire"
)

func Adapters() actor.Actors {
	panic(wire.Build(
		// Providers
		secrets.Set,
		repo.Set,
		// Adapters
		cron.Set,
		ChooseAdapters,
	))
}
