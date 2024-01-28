package subs

import (
	"woh/package/actor/third/gps/sub"
	webhooks "woh/webhooks"

	"github.com/google/wire"
)

type Adapter struct {
	*sub.Adapter
}

func New(m webhooks.Manager, o sub.Options) *Adapter {
	a := sub.New(o)
	a.Spawn(webhooks.ManagerKey, m)
	return &Adapter{
		Adapter: a,
	}
}

var Set = wire.NewSet(
	sub.ProvideOptions,
	sub.Set, // TODO: this might fail with more subscribers
	New,
)
