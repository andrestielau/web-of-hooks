package webhooks

import (
	"time"
)

type NewEvent struct {
	TenantId      string
	EventTypeKeys []string
	ReferenceID   string
}

type NewEventType struct {
	Key string
}
type EventTypeQuery struct {
	Limit        int
	Offset       int
	CreatedAfter time.Time
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
	TenantID     string
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
	FilterTypeIds []string
	Channels      []string
	SecretId      string
}
type EndpointQuery struct {
	Limit        int
	Offset       int
	CreatedAfter time.Time
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
type EndpointDetails struct {
	Endpoint
	FilterTypeIds []string
	Secret        string
}
type NewMessage struct {
	ApplicationID string
	EventTypeID   string
	EventID       string
	Payload       string
}
type MessageQuery struct {
	Limit        int
	Offset       int
	CreatedAfter time.Time
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
	Limit        int
	Offset       int
	CreatedAfter time.Time
}
type Attempt struct {
	ID             *int32
	Uid            string
	EndpointId     int32
	MessageId      int32
	Status         int32
	Retry          int32
	ResponseStatus int32
	Response       string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
type NewSecret struct {
	ApplicationID string
	Value         string
}
type SecretQuery struct {
	Limit        int
	Offset       int
	CreatedAfter time.Time
}
type Secret struct {
	ID            *int32
	Uid           string
	ApplicationID *int32
	Value         string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Payload struct {
	EventTypeKey string
	ReferenceID  string
}
