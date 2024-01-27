package domain

import (
	"woh/internal/provide/repo/queries"
	"woh/package/actor"
)

const ManagerKey = "webhooks-manager"

type Manager interface {
	actor.Actor
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
