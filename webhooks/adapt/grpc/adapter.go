package grpc

import (
	"woh/package/actor"
	"woh/package/actor/net/grpc/server"
	"woh/webhooks"
	webhooksv1 "woh/webhooks/adapt/grpc/v1"

	"github.com/google/wire"
	"google.golang.org/grpc"
)

type Options struct {
	Handler *Handler
}
type Adapter struct {
	*server.Adapter
}

func New(opts Options) *Adapter {
	a := server.New(server.Options{
		Handler: func(s *grpc.Server) {
			webhooksv1.RegisterWebHookServiceServer(s, opts.Handler)
		},
	})
	a.SpawnAll(actor.Actors{
		webhooks.SecretsKey: opts.Handler.Secrets,
		webhooks.RepoKey:    opts.Handler.Repo,
	})
	return &Adapter{
		Adapter: a,
	}
}

var Set = wire.NewSet(
	wire.Struct(new(Handler), "*"),
	wire.Struct(new(Options), "*"),
	New,
)
