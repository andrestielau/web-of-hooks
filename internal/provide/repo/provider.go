package repo

import (
	"woh/internal/domain"
	"woh/internal/provide/repo/queries"
	"woh/package/actor/sql/pgx"

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
