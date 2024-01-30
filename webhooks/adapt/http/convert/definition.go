package convert

import (
	"errors"
	"net/http"
	"time"

	"woh/package/utils"
	"woh/webhooks"
	webhooksv1 "woh/webhooks/adapt/http/v1"

	"github.com/samber/lo"
)

// goverter:converter
// goverter:extend TimeToString
// goverter:output:file ./generated.go
// goverter:output:package github.com/andrestielau/web-of-hooks/webhooks/adapt/http/convert
type Converter interface {
	// goverter:ignore ID TenantID Metadata
	// goverter:useZeroValueOnPointerInconsistency
	NewApplication(webhooksv1.NewApplication) webhooks.NewApplication
	NewApplications([]webhooksv1.NewApplication) []webhooks.NewApplication

	// goverter:map Uid Id
	// goverter:useZeroValueOnPointerInconsistency
	Application(webhooks.Application) webhooksv1.Application
	Applications([]webhooks.Application) []webhooksv1.Application

	// goverter:ignore ApplicationID Metadata
	// goverter:useZeroValueOnPointerInconsistency
	NewEndpoint(webhooksv1.NewEndpoint) webhooks.NewEndpoint
	NewEndpoints([]webhooksv1.NewEndpoint) []webhooks.NewEndpoint

	// goverter:map Uid Id
	// goverter:ignore Channels FilterTypes
	// goverter:useZeroValueOnPointerInconsistency
	Endpoint(webhooks.Endpoint) webhooksv1.Endpoint
	Endpoints([]webhooks.Endpoint) []webhooksv1.Endpoint

	// goverter:ignore ID TenantID
	// goverter:useZeroValueOnPointerInconsistency
	NewSecret(webhooksv1.NewSecret) webhooks.NewSecret
	NewSecrets([]webhooksv1.NewSecret) []webhooks.NewSecret

	// goverter:map Uid Id
	// goverter:ignore CreatedAt Value
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
}

func TimeToString(t time.Time) string {
	return t.String()
}

func Error(w http.ResponseWriter, err error) bool {
	var e utils.Error
	if !errors.As(err, &e) {
		http.Error(w, e.Reason, e.Code)
		return true
	}
	return false
}

func Errors(w http.ResponseWriter, err error) ([]webhooksv1.ErrorItem, bool) {
	var e utils.Errors
	if !errors.As(err, &e) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, true
	}
	return lo.Map(e, func(e utils.Error, _ int) webhooksv1.ErrorItem {
		return webhooksv1.ErrorItem{
			Code:   e.Code,
			Index:  e.Index,
			Reason: e.Reason,
		}
	}), false
}
