package http

import (
	"github.com/andrestielau/web-of-hooks/internal/domain"
	"github.com/andrestielau/web-of-hooks/package/actor/net/http/server"
	webhooksv1 "github.com/andrestielau/web-of-hooks/webhooks/http/v1"
	"github.com/google/wire"
)

type Adapter struct {
	*server.Adapter
}

func New(m domain.Manager) *Adapter {
	a := server.New(server.Options{
		Handler: webhooksv1.Handler(&Handler{}),
		Addr:    ":3000",
	})
	a.Spawn(domain.ManagerKey, m)
	return &Adapter{
		Adapter: a,
	}
}

var Set = wire.NewSet(
	New,
)
