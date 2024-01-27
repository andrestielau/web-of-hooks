package actor

import (
	"context"
	"sync"

	"github.com/andrestielau/web-of-hooks/package/utils"
)

type Actors map[string]Actor // Actors by id

type Actor interface {
	Start(context.Context) (bool, error) // Start Actor and Stopped Dependencies
	Stop(context.Context) (bool, error)  // Stop Actor and Orphan Dependencies
	Spawn(string, Actor) Actor           // Register dependency actors
	SpawnAll(Actors) Actor               // Register multiple dependency actors
	IsStarted() bool                     // Checks if actor is running
}

// Base
type Base struct {
	ctr      uint        // count dependants
	children Actors      // manage dependencies
	lock     *sync.Mutex // fearless concurrency
}

var _ Actor = &Base{}

func New() *Base {
	return &Base{
		children: make(Actors),
		lock:     &sync.Mutex{},
	}
}

// Spawn adds an actor as children
func (a *Base) Spawn(k string, v Actor) Actor {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.BaseSpawn(k, v)
}

// To use if Spawn needs to be overriten
func (a *Base) BaseSpawn(k string, v Actor) Actor {
	if _, ok := a.children[k]; !ok {
		a.children[k] = v
	}
	return a
}

// SpawnAll adds actors as children
func (a *Base) SpawnAll(m Actors) Actor {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.BaseSpawnAll(m)
}

// To use if SpawnAll needs to be overriten
func (a *Base) BaseSpawnAll(m Actors) Actor {
	for k, v := range m {
		a.BaseSpawn(k, v)
	}
	return a
}

// Start method
func (a *Base) Start(ctx context.Context) (bool, error) {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.BaseStart(ctx)
}

// To use if Start needs to be overriten, e.g: operations that need to happen only once before starting children.
func (a *Base) BaseStart(ctx context.Context) (bool, error) {
	if a.ctr++; a.ctr > 1 { // skip if not first call
		return false, nil
	}
	return true, utils.ForAll(a.children, func(k string) error {
		_, err := a.children[k].Start(ctx)
		return err
	})
}
func (a *Base) Stop(ctx context.Context) (bool, error) {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.BaseStop(ctx)
}

// To use if Stop needs to be overriten, e.g: operations that need to happen only once before stopping children.
func (a *Base) BaseStop(ctx context.Context) (bool, error) {
	if a.ctr--; a.ctr > 0 { // skip if not last call
		return false, nil
	}
	return true, utils.ForAll(a.children, func(k string) error {
		_, err := a.children[k].Stop(ctx)
		return err
	})
}
func (a *Base) IsStarted() bool { return a.ctr > 0 }
func (a *Base) Unlock()         { a.lock.Unlock() }
func (a *Base) Lock()           { a.lock.Lock() }
