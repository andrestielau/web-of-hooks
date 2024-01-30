package convert

import (
	"time"
	"woh/webhooks"
	"woh/webhooks/provide/repo/queries"
)

// goverter:converter
// goverter:extend ConvertTime
// goverter:output:file ./generated.go
// goverter:output:package github.com/andrestielau/web-of-hooks/webhooks/provide/repo/convert
type Converter interface {
	EventTypes([]queries.Application) []webhooks.EventType
	// goverter:ignore Key
	EventType(queries.Application) webhooks.EventType
	NewEventTypes([]webhooks.NewEventType) []queries.NewEventType
	NewEventType(webhooks.NewEventType) queries.NewEventType

	Applications([]queries.Application) []webhooks.Application
	// goverter:ignore Metadata CreatedAt UpdatedAt
	Application(queries.Application) webhooks.Application
	NewApplications([]webhooks.NewApplication) []queries.NewApplication
	// goverter:ignore Metadata
	NewApplication(webhooks.NewApplication) queries.NewApplication

	Endpoints([]queries.Endpoint) []webhooks.Endpoint
	// goverter:ignore Metadata UpdatedAt CreatedAt Disabled
	Endpoint(queries.Endpoint) webhooks.Endpoint
	NewEndpoints([]webhooks.NewEndpoint) []queries.NewEndpoint
	// goverter:ignore Metadata
	NewEndpoint(webhooks.NewEndpoint) queries.NewEndpoint

	Secrets([]queries.Secret) []webhooks.Secret
	Secret(queries.Secret) webhooks.Secret
	NewSecrets([]webhooks.NewSecret) []queries.NewSecret
	NewSecret(webhooks.NewSecret) queries.NewSecret

	Messages([]queries.Message) []webhooks.Message
	// goverter:ignore CreatedAt
	Message(queries.Message) webhooks.Message
	NewMessages([]webhooks.NewMessage) []queries.NewMessage
	NewMessage(webhooks.NewMessage) queries.NewMessage

	// goverter:ignore CreatedAt
	Attempt(queries.MessageAttempt) webhooks.Attempt
	Attempts([]queries.MessageAttempt) []webhooks.Attempt
	//NewAttempts([]webhooks.NewAttempt) []queries.NewMessageAttempt
	//NewAttempt(webhooks.NewAttempt) queries.NewMessageAttempt
}

func ConvertTime(t time.Time) time.Time { return t }
