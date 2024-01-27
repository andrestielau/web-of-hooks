package grpc

import (
	"woh/internal/domain"
	"woh/package/actor/net/grpc/server"

	"github.com/google/wire"
)

type Adapter struct {
	*server.Adapter
}

func New(m domain.Manager) *Adapter {
	a := server.New(server.Options{})
	a.Spawn(domain.ManagerKey, m)
	return &Adapter{
		Adapter: a,
	}
}

var Set = wire.NewSet(
	New,
)
