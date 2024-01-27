//go:build wireinject
// +build wireinject

package serve

import (
	"github.com/andrestielau/web-of-hooks/internal/domain/manager"
	"github.com/andrestielau/web-of-hooks/internal/provide/repo"
	"github.com/andrestielau/web-of-hooks/internal/provide/secrets"
	"github.com/andrestielau/web-of-hooks/package/actor"
	"github.com/andrestielau/web-of-hooks/package/actor/third/gps"
	"github.com/andrestielau/web-of-hooks/webhooks/grpc"
	"github.com/andrestielau/web-of-hooks/webhooks/http"
	"github.com/andrestielau/web-of-hooks/webhooks/subs"
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
