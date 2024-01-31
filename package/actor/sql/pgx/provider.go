package pgx

import (
	"context"

	"woh/package/actor"

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
	a.Base.Lock()
	defer a.Base.Unlock()
	return a.BaseStart(ctx)
}
func (a *Provider) BaseStart(ctx context.Context) (first bool, err error) {
	if first, err = a.Base.BaseStart(ctx); !first || err != nil {
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
