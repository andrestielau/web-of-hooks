package repo

import (
	"github.com/andrestielau/web-of-hooks/internal/domain"
	"github.com/andrestielau/web-of-hooks/internal/provide/repo/queries"
	"github.com/andrestielau/web-of-hooks/package/actor/sql/pgx"
	"github.com/google/wire"
)

type Provider struct {
	*pgx.Provider
	queries.Querier
}

func New(o pgx.Options) *Provider {
	return &Provider{
		Provider: pgx.New(o),
	}
}

var _ domain.Repository = &Provider{}

var Set = wire.NewSet(
	pgx.ProvideOptions,
	wire.Bind(new(domain.Repository), new(*Provider)),
	New,
)
