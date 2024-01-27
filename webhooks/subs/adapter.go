package subs

import (
	"github.com/andrestielau/web-of-hooks/internal/domain"
	"github.com/andrestielau/web-of-hooks/package/actor/third/gps/sub"
	"github.com/google/wire"
)

type Adapter struct {
	*sub.Adapter
}

func New(m domain.Manager, o sub.Options) *Adapter {
	a := sub.New(o)
	a.Spawn(domain.ManagerKey, m)
	return &Adapter{
		Adapter: a,
	}
}

var Set = wire.NewSet(
	sub.ProvideOptions,
	sub.Set, // TODO: this might fail with more subscribers
	New,
)
