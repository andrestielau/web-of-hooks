package grpc

import (
	"woh/package/actor/net/grpc/server"
	webhooks "woh/webhooks"

	"github.com/google/wire"
)

type Adapter struct {
	*server.Adapter
}

func New(m webhooks.Manager) *Adapter {
	a := server.New(server.Options{})
	a.Spawn(webhooks.ManagerKey, m)
	return &Adapter{
		Adapter: a,
	}
}

var Set = wire.NewSet(
	New,
)
