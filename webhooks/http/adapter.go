package http

import (
	"woh/internal/domain"
	"woh/package/actor/net/http/server"
	webhooksv1 "woh/webhooks/http/v1"

	"github.com/alexedwards/scs/v2"
	"github.com/google/wire"
)

type Adapter struct {
	*server.Adapter
}

func New(m domain.Manager) *Adapter {
	sessionManager = scs.New()
	a := server.New(server.Options{
		Handler: sessionManager.LoadAndSave(webhooksv1.Handler(&Handler{})),
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
