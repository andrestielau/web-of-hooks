package pgx

import (
	"context"

	"github.com/andrestielau/web-of-hooks/package/actor"
	"github.com/jackc/pgx/v4"
)

type Options struct {
	URL string
}
type Provider struct {
	*actor.Base
	opts Options
	*pgx.Conn
}

func New(opts Options) *Provider {
	return &Provider{
		Base: actor.New(),
		opts: opts,
	}
}

func (a *Provider) Start(ctx context.Context) (first bool, err error) {
	if first, err = a.Base.Start(ctx); !first || err != nil {
		return
	}
	a.Conn, err = pgx.Connect(ctx, a.opts.URL)
	return
}
func ProvideOptions() Options {
	return Options{
		URL: "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable",
	}
}
