//go:build wireinject
// +build wireinject

package serve

import (
	"woh/internal/domain/manager"
	"woh/internal/provide/repo"
	"woh/internal/provide/secrets"
	"woh/package/actor"
	"woh/package/actor/third/gps"
	"woh/webhooks/grpc"
	"woh/webhooks/http"
	"woh/webhooks/subs"

	"github.com/google/wire"
)

func Adapters() actor.Actors {
	panic(wire.Build(
		// Providers
		secrets.Set,
		repo.Set,
		// Services
		manager.Set,
		// Adapters
		gps.Set,
		subs.Set,
		grpc.Set,
		http.Set,
		ChooseAdapters,
	))
}
