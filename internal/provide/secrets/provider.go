package secrets

import (
	"github.com/andrestielau/web-of-hooks/internal/domain"
	"github.com/andrestielau/web-of-hooks/package/actor/third/vault"
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
