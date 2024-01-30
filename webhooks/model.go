package webhooks

import (
	"context"
	"time"
	"woh/package/actor"
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
	CreateApplications(context.Context, []NewApplication) ([]Application, error)
	DeleteApplications(context.Context, []string) error
	GetApplications(context.Context, []string) ([]Application, error)
	ListApplications(context.Context, int, int) ([]Application, error)
	DeleteAttempts(context.Context, []string) error
	GetAttempts(context.Context, []string) ([]Attempt, error)
	ListAttempts(context.Context,  int,  int) ([]Attempt, error)
	CreateEndpoints(context.Context, []NewEndpoint) ([]Endpoint, error)
	DeleteEndpoints(context.Context, []string) error
	GetEndpoints(context.Context, []string) ([]Endpoint, error)
	ListEndpoints(context.Context,  int,  int) ([]Endpoint, error)
	CreateEventTypes(context.Context, []NewEventType) ([]EventType, error)
	DeleteEventTypes(context.Context, []string) error
	GetEventTypes(context.Context, []string) ([]EventType, error)
	ListEventTypes(context.Context,  int,  int) ([]EventType, error)
	CreateMessages(context.Context, []NewMessage) ([]Message, error)
	DeleteMessages(context.Context, []string) error
	GetMessages(context.Context, []string) ([]Message, error)
	ListMessages(context.Context,  int,  int) ([]Message, error)
	CreateSecrets(context.Context, []NewSecret) ([]Secret, error)
	GetSecrets(context.Context, []string) ([]Secret, error)
	DeleteSecrets(context.Context, []string) error
	ListSecrets(context.Context, int, int) ([]Secret, error)
}

type NewEventType struct {
	Key string
}

type EventType struct {
	Key string
}

type NewApplication struct {
	ID        string
	Name      string
	TenantID  string
	RateLimit *int32
	Metadata  string
}
type Application struct {
	ID        *int32
	Uid       string
	Name      string
	TenantID  string
	RateLimit *int32
	Metadata  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
type NewEndpoint struct {
	Url           string
	Name          string
	ApplicationID string
	RateLimit     *int32
	Metadata      string
	Description   string
	FilterTypes   []string
	Channels      []string
}
type Endpoint struct {
	ID            *int32
	Uid           string
	Url           string
	Name          string
	Metadata      string
	Disabled      *bool
	RateLimit     *int32
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Description   string
	ApplicationID *int32
}
type NewMessage struct {
	ApplicationID string
	EventTypeID   string
	EventID       string
	Payload       string
}
type Message struct {
	ID            *int32
	Uid           string
	ApplicationID *int32
	EventTypeID   *int32
	EventID       string
	Payload       string
	CreatedAt     time.Time
}
type NewAttempt struct {
	ID        *int32
	Uid       string
	CreatedAt time.Time
}
type Attempt struct {
	ID        *int32
	Uid       string
	CreatedAt time.Time
}
type NewSecret struct {
	TenantID string
	ID       string
}
type Secret struct {
	ID  *int32
	Uid string
}
