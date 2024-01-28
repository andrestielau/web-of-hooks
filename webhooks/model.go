package webhooks

import (
	"context"
	"woh/package/actor"
	"woh/webhooks/provide/repo/queries"
)

const ManagerKey = "webhooks-manager"

type Manager interface {
	actor.Actor
	Repo() Repository
	Secrets() Secrets
	CreateEndpoints(context.Context) error
}

const WorkerKey = "webhooks-worker"

type Worker interface {
	actor.Actor
}

const SecretsKey = "webhooks-secrets"

type Secrets interface {
	actor.Actor
}

const RepoKey = "webhooks-repo"

type Repository interface {
	actor.Actor
	queries.Querier
}
