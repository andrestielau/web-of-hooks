package convert

import (
	"time"

	"woh/webhooks"
	webhooksv1 "woh/webhooks/adapt/http/v1"
)

// goverter:converter
// goverter:extend TimeToString
// goverter:matchIgnoreCase
// goverter:output:file ./generated.go
// goverter:output:package github.com/andrestielau/web-of-hooks/webhooks/adapt/http/convert
type Converter interface {
	// goverter:map Key Id
	// goverter:ignore CreatedAt
	// goverter:useZeroValueOnPointerInconsistency
	EventType(webhooks.EventType) webhooksv1.EventType
	EventTypes([]webhooks.EventType) []webhooksv1.EventType

	// goverter:ignore Offset After
	// goverter:useZeroValueOnPointerInconsistency
	EventTypeQuery(webhooksv1.ListEventTypesParams) webhooks.EventTypeQuery

	// goverter:ignore Metadata
	// goverter:useZeroValueOnPointerInconsistency
	NewApplication(webhooksv1.NewApplication) webhooks.NewApplication
	NewApplications([]webhooksv1.NewApplication) []webhooks.NewApplication

	// goverter:useZeroValueOnPointerInconsistency
	Application(webhooks.Application) webhooksv1.Application
	Applications([]webhooks.Application) []webhooksv1.Application

	// goverter:ignore Offset CreatedAfter
	// goverter:useZeroValueOnPointerInconsistency
	ApplicationQuery(webhooksv1.ListApplicationsParams) webhooks.ApplicationQuery

	// goverter:ignore ApplicationID Metadata
	// goverter:useZeroValueOnPointerInconsistency
	NewEndpoint(webhooksv1.NewEndpoint) webhooks.NewEndpoint
	NewEndpoints([]webhooksv1.NewEndpoint) []webhooks.NewEndpoint

	// goverter:map Uid Id
	// goverter:ignore Channels FilterTypes
	// goverter:useZeroValueOnPointerInconsistency
	Endpoint(webhooks.Endpoint) webhooksv1.Endpoint
	Endpoints([]webhooks.Endpoint) []webhooksv1.Endpoint

	// goverter:ignore Offset After
	// goverter:useZeroValueOnPointerInconsistency
	EndpointQuery(webhooksv1.ListEndpointsParams) webhooks.EndpointQuery

	// goverter:useZeroValueOnPointerInconsistency
	NewSecret(webhooksv1.NewSecret) webhooks.NewSecret
	NewSecrets([]webhooksv1.NewSecret) []webhooks.NewSecret

	// goverter:useZeroValueOnPointerInconsistency
	Secret(webhooks.Secret) webhooksv1.Secret
	Secrets([]webhooks.Secret) []webhooksv1.Secret

	// goverter:ignore ApplicationID EventTypeID EventID
	// goverter:useZeroValueOnPointerInconsistency
	NewMessage(webhooksv1.NewMessage) webhooks.NewMessage
	NewMessages([]webhooksv1.NewMessage) []webhooks.NewMessage

	// goverter:map Uid Id
	// goverter:ignore EventTypeId EventId
	// goverter:useZeroValueOnPointerInconsistency
	Message(webhooks.Message) webhooksv1.Message
	Messages([]webhooks.Message) []webhooksv1.Message

	// goverter:ignore Offset After
	// goverter:useZeroValueOnPointerInconsistency
	MessageQuery(webhooksv1.ListMessagesParams) webhooks.MessageQuery

	// goverter:map Uid Id
	// goverter:ignore EndpointId MessageId Response ResponseStatus Status
	// goverter:useZeroValueOnPointerInconsistency
	Attempt(webhooks.Attempt) webhooksv1.Attempt
	Attempts([]webhooks.Attempt) []webhooksv1.Attempt

	// goverter:ignore Offset After
	// goverter:useZeroValueOnPointerInconsistency
	AttemptQuery(webhooksv1.ListAttemptsParams) webhooks.AttemptQuery
}

func TimeToString(t time.Time) string {
	return t.String()
}
