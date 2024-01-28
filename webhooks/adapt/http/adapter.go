package http

import (
	"woh/package/actor/net/http/server"
	webhooks "woh/webhooks"
	webhooksv1 "woh/webhooks/adapt/http/v1"

	"github.com/alexedwards/scs/v2"
	"github.com/google/wire"
)

type Adapter struct {
	*server.Adapter
}

func New(m webhooks.Manager) *Adapter {
	sessionManager = scs.New()
	a := server.New(server.Options{
		Handler: sessionManager.LoadAndSave(webhooksv1.Handler(&Handler{
			Manager: m,
		})),
		Addr: ":3000",
	})
	a.Spawn(webhooks.ManagerKey, m)
	return &Adapter{
		Adapter: a,
	}
}

var Set = wire.NewSet(
	New,
)
