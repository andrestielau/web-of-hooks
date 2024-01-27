package worker

import (
	"woh/internal/domain"
	"woh/package/actor"

	"github.com/google/wire"
)

type Service struct {
	*actor.Base
}

func New(
	s domain.Secrets,
	r domain.Repository,
) *Service {
	base := actor.New()
	base.SpawnAll(actor.Actors{
		domain.SecretsKey: s,
		domain.RepoKey:    r,
	})
	return &Service{Base: base}
}

var Set = wire.NewSet(
	wire.Bind(new(domain.Worker), new(*Service)),
	New,
)
