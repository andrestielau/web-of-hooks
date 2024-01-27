package utils

import (
	"errors"
	"os"
	"os/signal"
	"sync"

	"github.com/samber/lo"
)

func ForAll[T any](in map[string]T, fn func(string) error) error {
	if len(in) == 0 {
		return nil
	}
	wg := sync.WaitGroup{}
	wg.Add(len(in))
	errs := make([]error, len(in))
	for i, c := range lo.Keys(in) {
		go func(i int, k string) {
			defer wg.Done()
			errs[i] = fn(k)
		}(i, c)
	}
	wg.Wait()
	return errors.Join(errs...)
}

func WaitSig(sigs ...os.Signal) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, sigs...)
	<-ch
}

func Apply[T any](t T, fns []func(T)) T {
	for i := range fns {
		fns[i](t)
	}
	return t
}
