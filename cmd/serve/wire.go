//go:build wireinject
// +build wireinject

package serve

import (
	"woh/package/actor"
	"woh/webhooks/adapt/cron"
	"woh/webhooks/adapt/grpc"
	"woh/webhooks/adapt/http"
	"woh/webhooks/adapt/subs"
	pubs "woh/webhooks/provide/pub"
	"woh/webhooks/provide/repo"
	"woh/webhooks/provide/secrets"

	"woh/package/actor/router"
	"woh/package/actor/third/gps/pub"
	"woh/package/actor/third/gps/sub"

	"github.com/google/wire"
)

func Adapters() (actor.Actors, error) {
	panic(wire.Build(
		// Providers
		secrets.Set,
		repo.Set,
		pubs.Set,
		pub.Set,
		sub.Set,
		router.Set,
		// Adapters
		subs.Set,
		grpc.Set,
		http.Set,
		cron.Set,
		ChooseAdapters,
	))
}
