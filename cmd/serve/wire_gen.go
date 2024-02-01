// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package serve

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/alexedwards/scs/v2"
	"woh/package/actor"
	"woh/package/actor/router"
	"woh/package/actor/sql/pgx"
	"woh/package/actor/third/gps/pub"
	"woh/package/actor/third/gps/sub"
	"woh/webhooks/adapt/cron"
	"woh/webhooks/adapt/cron/monitor"
	"woh/webhooks/adapt/grpc"
	convert2 "woh/webhooks/adapt/grpc/convert"
	handle2 "woh/webhooks/adapt/grpc/handle"
	"woh/webhooks/adapt/http"
	"woh/webhooks/adapt/http/convert"
	"woh/webhooks/adapt/http/handle"
	"woh/webhooks/adapt/subs"
	"woh/webhooks/adapt/subs/dispatcher"
	"woh/webhooks/adapt/subs/projector"
	"woh/webhooks/adapt/subs/subscriber"
	"woh/webhooks/provide/repo"
	"woh/webhooks/provide/secrets"
)

// Injectors from wire.go:

func Adapters() (actor.Actors, error) {
	sessionManager := scs.New()
	options := pgx.ProvideOptions()
	provider := repo.New(options)
	secretsProvider := secrets.New()
	converterImpl := _wireConverterImplValue
	handler := &handle.Handler{
		Session: sessionManager,
		Repo:    provider,
		Secrets: secretsProvider,
		Convert: converterImpl,
	}
	httpOptions := http.Options{
		Handler: handler,
	}
	adapter := http.New(httpOptions)
	convertConverterImpl := _wireConvertConverterImplValue
	handleHandler := &handle2.Handler{
		Repo:    provider,
		Secrets: secretsProvider,
		Convert: convertConverterImpl,
	}
	grpcOptions := grpc.Options{
		Handler: handleHandler,
	}
	grpcAdapter := grpc.New(grpcOptions)
	publisherConfig := pub.ProvideConfig()
	loggerAdapter := router.ProvideLogger()
	pubOptions := pub.Options{
		Config: publisherConfig,
		Logger: loggerAdapter,
	}
	pubProvider := pub.New(pubOptions)
	monitorHandler := &monitor.Handler{
		Publisher: pubProvider,
		Repo:      provider,
	}
	cronOptions := cron.Options{
		Monitor: monitorHandler,
	}
	cronAdapter := cron.New(cronOptions)
	subscriberHandler := subscriber.Handler{
		Repo: provider,
	}
	dispatcherHandler := dispatcher.Handler{
		Repo: provider,
	}
	projectorHandler := projector.Handler{
		Repo: provider,
	}
	subscriberConfig := sub.ProvideConfig()
	subOptions := sub.Options{
		Config: subscriberConfig,
		Logger: loggerAdapter,
	}
	subAdapter := sub.New(subOptions)
	routerConfig := router.ProvideConfig()
	messageRouter, err := message.NewRouter(routerConfig, loggerAdapter)
	if err != nil {
		return nil, err
	}
	routerOptions := router.Options{
		Router: messageRouter,
	}
	routerAdapter := router.New(routerOptions)
	subsOptions := subs.Options{
		Subscriber: subscriberHandler,
		Dispatcher: dispatcherHandler,
		Projector:  projectorHandler,
		Repo:       provider,
		Pub:        pubProvider,
		Sub:        subAdapter,
		Adapter:    routerAdapter,
	}
	subsAdapter := subs.New(subsOptions)
	actors := ChooseAdapters(adapter, grpcAdapter, cronAdapter, subsAdapter)
	return actors, nil
}

var (
	_wireConverterImplValue        = &convert.ConverterImpl{}
	_wireConvertConverterImplValue = &convert2.ConverterImpl{}
)
