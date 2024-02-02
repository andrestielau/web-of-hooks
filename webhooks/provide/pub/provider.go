package pub

import (
	"woh/package/actor"
	"woh/package/actor/third/gps/pub"

	"woh/webhooks/provide/pub/publish"
	"github.com/google/wire"
)

type Options struct {
	Pub *pub.Provider
}
type Provider struct {
	*actor.Base
	Pub *pub.Provider
}

func New(o Options) *Provider {
	provider := &Provider{
		Base: actor.New(),
		Pub:  o.Pub,
	}

	provider.SpawnAll(actor.Actors{
		pub.Key: o.Pub,
	})
	return provider
}

var Set = wire.NewSet(
	wire.Struct(new(publish.Publisher), "*"),
	wire.Struct(new(Options), "*"),
	New,
)
