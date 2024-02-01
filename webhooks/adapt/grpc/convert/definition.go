package convert

import (
	"time"

	"woh/webhooks"
	webhooksv1 "woh/webhooks/adapt/grpc/v1"
)

// goverter:converter
// goverter:extend TimeToString
// goverter:extend Int32ToInt
// goverter:matchIgnoreCase
// goverter:ignoreUnexported
// goverter:output:file ./generated.go
// goverter:output:package github.com/andrestielau/web-of-hooks/webhooks/adapt/grpc/convert
type Converter interface {

	// goverter:ignore Metadata
	// goverter:useZeroValueOnPointerInconsistency
	NewApplication(*webhooksv1.App) webhooks.NewApplication
	NewApplications([]*webhooksv1.App) []webhooks.NewApplication

	// goverter:ignore Metadata
	// goverter:useZeroValueOnPointerInconsistency
	Application(webhooks.Application) *webhooksv1.App
	Applications([]webhooks.Application) []*webhooksv1.App

	// goverter:ignore CreatedAfter
	// goverter:useZeroValueOnPointerInconsistency
	ApplicationQuery(*webhooksv1.PageRequest) webhooks.ApplicationQuery

	// goverter:ignore Name ApplicationID Metadata FilterTypeIds
	// goverter:useZeroValueOnPointerInconsistency
	NewEndpoint(*webhooksv1.Endpoint) webhooks.NewEndpoint
	NewEndpoints([]*webhooksv1.Endpoint) []webhooks.NewEndpoint

	// goverter:map Uid Id
	// goverter:ignore Channels FilterTypes Secret Version Metadata
	// goverter:useZeroValueOnPointerInconsistency
	Endpoint(webhooks.Endpoint) *webhooksv1.Endpoint
	Endpoints([]webhooks.Endpoint) []*webhooksv1.Endpoint

	// goverter:ignore Offset After
	// goverter:useZeroValueOnPointerInconsistency
	EndpointQuery(*webhooksv1.PageRequest) webhooks.EndpointQuery

	// goverter:ignore ApplicationID EventTypeID EventID Payload
	// goverter:useZeroValueOnPointerInconsistency
	NewMessage(*webhooksv1.Message) webhooks.NewMessage
	NewMessages([]*webhooksv1.Message) []webhooks.NewMessage

	// goverter:map Uid Id
	// goverter:ignore EventType EventId Timestamp Payload Tags Channels
	// goverter:useZeroValueOnPointerInconsistency
	Message(webhooks.Message) *webhooksv1.Message
	Messages([]webhooks.Message) []*webhooksv1.Message

	// goverter:ignore Offset After
	// goverter:useZeroValueOnPointerInconsistency
	MessageQuery(*webhooksv1.PageRequest) webhooks.MessageQuery

	// goverter:map Uid Id
	// goverter:useZeroValueOnPointerInconsistency
	Attempt(webhooks.Attempt) *webhooksv1.Attempt
	Attempts([]webhooks.Attempt) []*webhooksv1.Attempt

	// goverter:ignore Offset After
	// goverter:useZeroValueOnPointerInconsistency
	// AttemptQuery(*webhooksv1.ListAttemptsRequest) webhooks.AttemptQuery
}

func TimeToString(t time.Time) string {
	return t.String()
}

func Int32ToInt(i int32) int {
	return int(i)
}
