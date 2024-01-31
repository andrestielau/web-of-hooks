package server

import (
	"context"
	"log"
	"net"

	"woh/package/actor"
	"woh/package/utils"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Options struct {
	Handler func(*grpc.Server)
	Addr    string
}
type Adapter struct {
	*actor.Base
	opts Options
	*grpc.Server
	closer utils.Closer
}

func New(opts Options) *Adapter {
	return &Adapter{
		Base:   actor.New(),
		opts:   opts,
		closer: utils.NewCloser(),
	}
}

func (a *Adapter) Start(ctx context.Context) (first bool, err error) {
	if first, err = a.Base.Start(ctx); !first || err != nil {
		return
	}
	l, err := net.Listen("tcp", a.opts.Addr)
	if err != nil {
		return
	}
	a.Server = grpc.NewServer()
	// Register reflection service on gRPC server.
	reflection.Register(a.Server)
	a.opts.Handler(a.Server)
	go func() {
		defer a.closer.Close()
		log.Println("grpc listening at http://localhost" + a.opts.Addr)
		log.Fatal(a.Server.Serve(l))
	}()
	return
}

func (a *Adapter) Stop(ctx context.Context) (last bool, err error) {
	if last, err = a.Base.Stop(ctx); !last || err != nil {
		return
	}
	a.Server.GracefulStop()
	a.closer.Wait()
	return
}
