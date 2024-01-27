package vault

import (
	"context"

	"woh/package/actor"

	vault "github.com/hashicorp/vault/api"
)

type Options struct {
	*vault.Config
}
type Provider struct {
	*actor.Base
	opts   Options
	Client *vault.Client
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
	a.Client, err = vault.NewClient(a.opts.Config)
	return
}

func ProvideOptions(key string) Options {
	return Options{
		Config: &vault.Config{
			Address: "http://localhost:8201",
		},
	}
}
