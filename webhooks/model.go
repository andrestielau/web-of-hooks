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
	ListApplications(context.Context, ApplicationQuery) ([]Application, error)
	DeleteAttempts(context.Context, []string) error
	GetAttempts(context.Context, []string) ([]Attempt, error)
	ListAttempts(context.Context, AttemptQuery) ([]Attempt, error)
	CreateEndpoints(context.Context, []NewEndpoint) ([]Endpoint, error)
	DeleteEndpoints(context.Context, []string) error
	GetEndpoints(context.Context, []string) ([]Endpoint, error)
	ListEndpoints(context.Context, EndpointQuery) ([]Endpoint, error)
	CreateEventTypes(context.Context, []NewEventType) ([]EventType, error)
	DeleteEventTypes(context.Context, []string) error
	GetEventTypes(context.Context, []string) ([]EventType, error)
	ListEventTypes(context.Context, EventTypeQuery) ([]EventType, error)
	CreateMessages(context.Context, []NewMessage) ([]MessageDetails, error)
	DeleteMessages(context.Context, []string) error
	GetMessages(context.Context, []string) ([]MessageDetails, error)
	ListMessages(context.Context, MessageQuery) ([]Message, error)
	CreateSecrets(context.Context, []NewSecret) ([]Secret, error)
	GetSecrets(context.Context, []string) ([]Secret, error)
	DeleteSecrets(context.Context, []string) error
	ListSecrets(context.Context, SecretQuery) ([]Secret, error)
	SetLastSeen(context.Context, string) error
}

type NewEventType struct {
	Key string
}
type EventTypeQuery struct {
	Limit  int
	Offset int
	After  string
}

type EventType struct {
	ID  int32
	Uid string
	Key string
}

type NewApplication struct {
	Name      string
	TenantID  string
	RateLimit *int32
	Metadata  string
}
type ApplicationQuery struct {
	Limit        int
	Offset       int
	CreatedAfter time.Time
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
type EndpointQuery struct {
	Limit  int
	Offset int
	After  string
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
type MessageQuery struct {
	Limit  int
	Offset int
	After  string
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
type MessageDetails struct {
	Message
	Attempts []Attempt
}
type NewAttempt struct {
	ID        *int32
	Uid       string
	CreatedAt time.Time
}
type AttemptQuery struct {
	Limit  int
	Offset int
	After  string
}
type Attempt struct {
	ID        *int32
	Uid       string
	CreatedAt time.Time
}
type NewSecret struct {
	ApplicationID *int32
	Value         string
}
type SecretQuery struct {
	Limit  int
	Offset int
	After  string
}
type Secret struct {
	ID            *int32
	Uid           string
	ApplicationID *int32
	Value         string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
