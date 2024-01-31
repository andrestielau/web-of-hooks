package grpc

import (
	"woh/package/actor"
	"woh/package/actor/net/grpc/server"
	"woh/webhooks"
	"woh/webhooks/adapt/grpc/handle"
	webhooksv1 "woh/webhooks/adapt/grpc/v1"

	"woh/webhooks/adapt/grpc/convert"
	"github.com/google/wire"
	"google.golang.org/grpc"
)

type Options struct {
	Handler *handle.Handler
}
type Adapter struct {
	*server.Adapter
}

func New(opts Options) *Adapter {
	a := server.New(server.Options{
		Handler: func(s *grpc.Server) {
			webhooksv1.RegisterWebHookServiceServer(s, opts.Handler)
		},
		Addr: ":3001",
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
	wire.Bind(new(convert.Converter), new(*convert.ConverterImpl)),
	wire.Struct(new(handle.Handler), "*"),
	wire.Value(&convert.ConverterImpl{}),
	wire.Struct(new(Options), "*"),
	New,
)
