package router

import (
	"context"

	"woh/package/actor"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/wire"
)

const Key = "Router"

type HandlerOptions struct {
	In, Out string
	Pub     message.Publisher
	Sub     message.Subscriber
	Func    message.HandlerFunc
}
type Options struct {
	Router *message.Router
}
type Adapter struct {
	*message.Router
	*actor.Base
}

func New(opts Options) *Adapter {
	return &Adapter{
		Base:   actor.New(),
		Router: opts.Router,
	}
}
func (a *Adapter) Handle(opts map[string]HandlerOptions) {
	for key, opt := range opts {
		a.Router.AddHandler(key, opt.In, opt.Sub, opt.Out, opt.Pub, opt.Func)
	}
}
func (h *Adapter) Start(ctx context.Context) (first bool, err error) {
	if first, err = h.Base.Start(ctx); !first || err != nil {
		return
	}
	err = h.Router.Run(ctx)
	return
}
func (h *Adapter) Stop(ctx context.Context) (last bool, err error) {
	if last, err = h.Base.Stop(ctx); !last || err != nil {
		return
	}
	err = h.Router.Close()
	return
}

var Set = wire.NewSet(
	wire.Struct(new(Options), "*"),
	message.NewRouter,
	ProvideConfig,
	ProvideLogger,
	New,
)

func ProvideConfig() message.RouterConfig {
	return message.RouterConfig{}
}
func ProvideLogger() watermill.LoggerAdapter {
	return watermill.NewStdLogger(true, true)
}
