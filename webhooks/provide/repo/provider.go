package repo

import (
	"context"
	"woh/package/actor/sql/pgx"
	webhooks "woh/webhooks"
	"woh/webhooks/provide/repo/convert"
	"woh/webhooks/provide/repo/queries"

	"github.com/google/wire"
)

type Provider struct {
	*pgx.Provider
	*Repository
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
	a.Repository = &Repository{
		Querier: queries.NewQuerier(a.Conn),
		Convert: &convert.ConverterImpl{},
	}
	return
}

var _ webhooks.Repository = &Provider{}

var Set = wire.NewSet(
	pgx.ProvideOptions,
	wire.Bind(new(webhooks.Repository), new(*Provider)),
	New,
)
