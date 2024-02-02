package webhooks

import (
	"context"

	"woh/package/actor"
)

const (
	API_PORT  = 3000
	GRPC_PORT = 3001
)

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
	CreateApplications(context.Context, []NewApplication) ([]Application, error)
	DeleteApplications(context.Context, []string) error
	GetApplications(context.Context, []string) ([]Application, error)
	GetApplicationsByName(context.Context, []string) ([]Application, error)
	ListApplications(context.Context, ApplicationQuery) ([]Application, error)
	DeleteAttempts(context.Context, []string) error
	GetAttempts(context.Context, []string) ([]Attempt, error)
	ListAttempts(context.Context, AttemptQuery) ([]Attempt, error)
	CreateEndpoints(context.Context, []NewEndpoint) ([]EndpointDetails, error)
	DeleteEndpoints(context.Context, []string) error
	GetEndpoints(context.Context, []string) ([]EndpointDetails, error)
	GetEndpointsByUrl(context.Context, []string) ([]EndpointDetails, error)
	GetEndpointsByTenantAndEventTypes(context.Context, string, []string) ([]EndpointDetails, error)
	ListEndpoints(context.Context, EndpointQuery) ([]Endpoint, error)
	ListApplicationEndpoints(context.Context, string, EndpointQuery) ([]Endpoint, error)
	CreateEventTypes(context.Context, []NewEventType) ([]EventType, error)
	DeleteEventTypes(context.Context, []string) error
	GetEventTypes(context.Context, []string) ([]EventType, error)
	GetEventTypesByKeys(context.Context, []string) ([]EventType, error)
	ListEventTypes(context.Context, EventTypeQuery) ([]EventType, error)
	CreateMessages(context.Context, []NewMessage) ([]MessageDetails, error)
	DeleteMessages(context.Context, []string) error
	GetMessages(context.Context, []string) ([]MessageDetails, error)
	ListMessages(context.Context, MessageQuery) ([]Message, error)
	ListApplicationMessages(context.Context, string, MessageQuery) ([]Message, error)
	CreateSecrets(context.Context, []NewSecret) ([]Secret, error)
	GetSecrets(context.Context, []string) ([]Secret, error)
	DeleteSecrets(context.Context, []string) error
	ListSecrets(context.Context, SecretQuery) ([]Secret, error)
	ListApplicationSecrets(context.Context, string) ([]Secret, error)
	EmitEvent(context.Context, NewEvent) ([]Message, error)
}

const PublisherKey = "webhooks-pub"

type Publisher interface {
	actor.Actor
	Publish(context.Context, Payload) error
}
