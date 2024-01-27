package actor

import "context"

type System struct {
	ctx context.Context
	*Base
}

func NewSystem(ctx context.Context) *System {
	return &System{ctx: ctx, Base: New()}
}

func (s *System) Start(actors Actors) (err error) {
	_, err = s.SpawnAll(actors).Start(s.ctx)
	return
}
func (s *System) Run() (err error) {
	_, err = s.Base.Start(s.ctx)
	return
}
func (s *System) Stop() {
	s.Base.Stop(s.ctx)
}
