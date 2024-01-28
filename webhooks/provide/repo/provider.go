package repo

import (
	"context"
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

func (a *Provider) Start(ctx context.Context) (first bool, err error) {
	if first, err = a.Provider.Start(ctx); !first || err != nil {
		return
	}
	a.Querier = queries.NewQuerier(a.Conn)
	return
}

var _ webhooks.Repository = &Provider{}

var Set = wire.NewSet(
	pgx.ProvideOptions,
	wire.Bind(new(webhooks.Repository), new(*Provider)),
	New,
)
