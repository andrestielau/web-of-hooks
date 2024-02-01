package server

import (
	"context"
	"log"
	"net/http"

	"woh/package/actor"
	"woh/package/utils"
)

type Options struct {
	Handler http.Handler
	Addr    string
}
type Adapter struct {
	*actor.Base
	opts Options
	*http.Server
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
	a.Server = &http.Server{Addr: a.opts.Addr, Handler: a.opts.Handler}
	go func() {
		defer a.closer.Close()
		log.Println("http serving at " + a.Addr)
		if err := a.Server.ListenAndServe(); err != nil {
			log.Println("http server", err)
		}
	}()
	return
}

func (a *Adapter) Stop(ctx context.Context) (last bool, err error) {
	if last, err = a.Base.Stop(ctx); !last || err != nil {
		return
	} else if err = a.Server.Close(); err != nil {
		return
	}
	a.closer.Wait()
	return
}
