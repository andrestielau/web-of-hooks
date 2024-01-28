package secrets

import (
	"woh/package/actor/third/vault"
	webhooks "woh/webhooks"

	"github.com/google/wire"
)

type Provider struct {
	*vault.Provider
}

func New() *Provider {
	return &Provider{
		Provider: vault.New(vault.Options{}),
	}
}

var Set = wire.NewSet(
	vault.ProvideOptions,
	wire.Bind(new(webhooks.Secrets), new(*Provider)),
	New,
)
