package manager

import (
	"context"
	"woh/package/actor"
	webhooks "woh/webhooks"

	"github.com/google/wire"
)

type Service struct {
	*actor.Base
	secrets webhooks.Secrets
	repo    webhooks.Repository
}

var _ webhooks.Manager = &Service{}

func New(
	s webhooks.Secrets,
	r webhooks.Repository,
) *Service {
	base := actor.New()
	base.SpawnAll(actor.Actors{
		webhooks.SecretsKey: s,
		webhooks.RepoKey:    r,
	})
	return &Service{
		Base:    base,
		secrets: s,
		repo:    r,
	}
}

var Set = wire.NewSet(
	wire.Bind(new(webhooks.Manager), new(*Service)),
	New,
)

func (s *Service) CreateEndpoints(context.Context) error {
	// Creat
	panic("unimplemented")
}

func (s *Service) Repo() webhooks.Repository { return s.repo }
func (s *Service) Secrets() webhooks.Secrets { return s.secrets }
