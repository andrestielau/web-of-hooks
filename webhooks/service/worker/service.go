package worker

import (
	"woh/package/actor"
	webhooks "woh/webhooks"

	"github.com/google/wire"
)

type Service struct {
	*actor.Base
}

func New(
	s webhooks.Secrets,
	r webhooks.Repository,
) *Service {
	base := actor.New()
	base.SpawnAll(actor.Actors{
		webhooks.SecretsKey: s,
		webhooks.RepoKey:    r,
	})
	return &Service{Base: base}
}

var Set = wire.NewSet(
	wire.Bind(new(webhooks.Worker), new(*Service)),
	New,
)
