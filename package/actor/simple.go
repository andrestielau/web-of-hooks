package actor

import (
	"context"
	"log"
)

type Simple[T any] struct {
	fn     func(context.Context, T)
	closer chan struct{}
	ch     chan T
	*Base
}

func NewSimple[T any](ch chan T, fn func(context.Context, T)) *Simple[T] {
	return &Simple[T]{ch: ch, fn: fn, Base: New()}
}

func (a *Simple[T]) Start(ctx context.Context) (first bool, err error) {
	if first, err = a.Base.Start(ctx); !first || err != nil {
		return first, err
	}
	a.Run(ctx)
	return true, nil
}

func (a *Simple[T]) Run(ctx context.Context) {
	a.closer = make(chan struct{})
	go func() {
		defer close(a.closer)
		for msg := range a.ch {
			if a.run(ctx, msg) {
				break
			}
		}
	}()
}

func (a *Simple[T]) run(ctx context.Context, msg T) (stop bool) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
			stop = true
			return
		}
	}()
	a.fn(ctx, msg)
	return
}
func (a *Simple[T]) Stop(ctx context.Context) (last bool, err error) {
	if last, err = a.Base.Start(ctx); !last || err != nil {
		return last, err
	}
	close(a.ch)
	<-a.closer
	return true, nil
}
