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
	// goverter:useZeroValueOnPointerInconsistency
	EventType(queries.EventType) webhooks.EventType
	EventTypes([]queries.EventType) []webhooks.EventType
	NewEventType(webhooks.NewEventType) queries.NewEventType
	NewEventTypes([]webhooks.NewEventType) []queries.NewEventType
	EventTypeQuery(webhooks.EventTypeQuery) queries.ListEventTypesParams

	// goverter:ignore Metadata
	Application(queries.Application) webhooks.Application
	Applications([]queries.Application) []webhooks.Application
	// goverter:ignore Metadata
	NewApplication(webhooks.NewApplication) queries.NewApplication
	NewApplications([]webhooks.NewApplication) []queries.NewApplication
	ApplicationQuery(webhooks.ApplicationQuery) queries.ListApplicationsParams

	// goverter:ignore Metadata Disabled
	Endpoint(queries.Endpoint) webhooks.Endpoint
	Endpoints([]queries.Endpoint) []webhooks.Endpoint
	// goverter:ignore Metadata
	NewEndpoint(webhooks.NewEndpoint) queries.NewEndpoint
	NewEndpoints([]webhooks.NewEndpoint) []queries.NewEndpoint
	EndpointQuery(webhooks.EndpointQuery) queries.ListEndpointsParams
	// goverter:ignore ApplicationUid
	ApplicationEndpointQuery(webhooks.EndpointQuery) queries.ListApplicationEndpointsParams
	EndpointDetail(queries.EndpointDetails) webhooks.EndpointDetails
	EndpointDetails([]queries.EndpointDetails) []webhooks.EndpointDetails

	Secret(queries.Secret) webhooks.Secret
	Secrets([]queries.Secret) []webhooks.Secret
	NewSecret(webhooks.NewSecret) queries.NewSecret
	NewSecrets([]webhooks.NewSecret) []queries.NewSecret
	SecretQuery(webhooks.SecretQuery) queries.ListSecretsParams

	Message(queries.Message) webhooks.Message
	Messages([]queries.Message) []webhooks.Message
	NewMessage(webhooks.NewMessage) queries.NewMessage
	NewMessages([]webhooks.NewMessage) []queries.NewMessage
	MessageQuery(webhooks.MessageQuery) queries.ListMessagesParams
	// goverter:ignore ApplicationUid
	ApplicationMessageQuery(webhooks.MessageQuery) queries.ListApplicationMessagesParams
	MessageDetail(queries.MessageDetails) webhooks.MessageDetails
	MessageDetails([]queries.MessageDetails) []webhooks.MessageDetails

	// goverter:map MessageID MessageId
	// goverter:map EndpointID EndpointId
	// goverter:useZeroValueOnPointerInconsistency
	Attempt(queries.MessageAttempt) webhooks.Attempt
	Attempts([]queries.MessageAttempt) []webhooks.Attempt
	AttemptQuery(webhooks.AttemptQuery) queries.ListAttemptsParams
}

func ConvertTime(t time.Time) time.Time { return t }
