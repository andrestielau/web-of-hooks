// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package worker

import (
	"woh/package/actor"
	"woh/package/actor/sql/pgx"
	"woh/webhooks/adapt/cron"
	"woh/webhooks/provide/repo"
	"woh/webhooks/provide/secrets"
	"woh/webhooks/service/worker"
)

// Injectors from wire.go:

func Adapters() actor.Actors {
	provider := secrets.New()
	options := pgx.ProvideOptions()
	repoProvider := repo.New(options)
	service := worker.New(provider, repoProvider)
	adapter := cron.New(service)
	actors := ChooseAdapters(adapter)
	return actors
}
