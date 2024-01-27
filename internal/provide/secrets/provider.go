package secrets

import (
	"woh/internal/domain"
	"woh/package/actor/third/vault"

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
	wire.Bind(new(domain.Secrets), new(*Provider)),
	New,
)
