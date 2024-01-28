package repo

import (
	"woh/package/actor/sql/pgx"
	webhooks "woh/webhooks"
	"woh/webhooks/provide/repo/queries"

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

var _ webhooks.Repository = &Provider{}

var Set = wire.NewSet(
	pgx.ProvideOptions,
	wire.Bind(new(webhooks.Repository), new(*Provider)),
	New,
)
