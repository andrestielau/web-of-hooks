//go:build wireinject
// +build wireinject

package serve

import (
	"woh/package/actor"
	"woh/package/actor/third/gps"
	"woh/webhooks/adapt/grpc"
	"woh/webhooks/adapt/http"
	"woh/webhooks/adapt/subs"
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
		gps.Set,
		subs.Set,
		grpc.Set,
		http.Set,
		ChooseAdapters,
	))
}
