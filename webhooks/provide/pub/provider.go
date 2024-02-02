package pub

import (
	"woh/package/actor"
	"woh/package/actor/third/gps/pub"

	"woh/webhooks/provide/pub/publish"

	"woh/webhooks"

	"github.com/google/wire"
)

type Options struct {
	Pub  *pub.Provider
	Repo webhooks.Repository
}
type Provider struct {
	*actor.Base
	Pub *pub.Provider
}

func New(opts Options) *Provider {
	provider := &Provider{
		Base: actor.New(),
		Pub:  opts.Pub,
	}

	provider.SpawnAll(actor.Actors{
		pub.Key:          opts.Pub,
		webhooks.RepoKey: opts.Repo,
	})
	return provider
}

var Set = wire.NewSet(
	wire.Struct(new(publish.Publisher), "*"),
	wire.Struct(new(Options), "*"),
	New,
)
