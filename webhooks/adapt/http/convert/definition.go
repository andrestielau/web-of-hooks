package convert

import (
	"errors"
	"net/http"
	"time"
	webhooksv1 "woh/webhooks/adapt/http/v1"
	"woh/webhooks/provide/repo/queries"

	"github.com/andrestielau/web-of-hooks/package/utils"
	"github.com/samber/lo"
)

// goverter:converter
// goverter:extend TimeToString
// goverter:output:file ./generated.go
// goverter:output:package github.com/andrestielau/web-of-hooks/webhooks/adapt/http/convert
type ApplicationConverter interface {
	// goverter:ignore ID TenantID Metadata
	// goverter:useZeroValueOnPointerInconsistency
	NewItem(webhooksv1.NewApplication) queries.NewApplication
	New([]webhooksv1.NewApplication) []queries.NewApplication

	// goverter:map Uid Id
	// goverter:useZeroValueOnPointerInconsistency
	ListItem(queries.ListApplicationsRow) webhooksv1.Application
	List([]queries.ListApplicationsRow) []webhooksv1.Application

	// goverter:map Uid Id
	// goverter:useZeroValueOnPointerInconsistency
	CreatedItem(queries.CreateApplicationsRow) webhooksv1.Application
	Created([]queries.CreateApplicationsRow) []webhooksv1.Application

	// goverter:map Uid Id
	// goverter:useZeroValueOnPointerInconsistency
	GotItem(queries.GetApplicationsRow) webhooksv1.Application
	Got([]queries.GetApplicationsRow) []webhooksv1.Application
}

// goverter:converter
// goverter:extend TimeToString
// goverter:output:file ./generated.go
// goverter:output:package github.com/andrestielau/web-of-hooks/webhooks/adapt/http/convert
type EndpointConverter interface {
	// goverter:ignore ApplicationID Metadata
	// goverter:useZeroValueOnPointerInconsistency
	NewItem(webhooksv1.NewEndpoint) queries.NewEndpoint
	New([]webhooksv1.NewEndpoint) []queries.NewEndpoint

	// goverter:map Uid Id
	// goverter:ignore Channels FilterTypes
	// goverter:useZeroValueOnPointerInconsistency
	ListItem(queries.ListEndpointsRow) webhooksv1.Endpoint
	List([]queries.ListEndpointsRow) []webhooksv1.Endpoint

	// goverter:map Uid Id
	// goverter:ignore Channels FilterTypes
	// goverter:useZeroValueOnPointerInconsistency
	CreatedItem(queries.CreateEndpointsRow) webhooksv1.Endpoint
	Created([]queries.CreateEndpointsRow) []webhooksv1.Endpoint

	// goverter:map Uid Id
	// goverter:ignore Channels FilterTypes
	// goverter:useZeroValueOnPointerInconsistency
	GotItem(queries.GetEndpointsRow) webhooksv1.Endpoint
	Got([]queries.GetEndpointsRow) []webhooksv1.Endpoint
}

// goverter:converter
// goverter:extend TimeToString
// goverter:output:file ./generated.go
// goverter:output:package github.com/andrestielau/web-of-hooks/webhooks/adapt/http/convert
type SecretConverter interface {
	// goverter:ignore ID TenantID
	// goverter:useZeroValueOnPointerInconsistency
	NewItem(webhooksv1.NewSecret) queries.NewSecret
	New([]webhooksv1.NewSecret) []queries.NewSecret

	// goverter:map Uid Id
	// goverter:ignore CreatedAt Value
	// goverter:useZeroValueOnPointerInconsistency
	ListItem(queries.ListSecretsRow) webhooksv1.Secret
	List([]queries.ListSecretsRow) []webhooksv1.Secret

	// goverter:map Uid Id
	// goverter:ignore CreatedAt Value
	// goverter:useZeroValueOnPointerInconsistency
	CreatedItem(queries.CreateSecretsRow) webhooksv1.Secret
	Created([]queries.CreateSecretsRow) []webhooksv1.Secret

	// goverter:map Uid Id
	// goverter:ignore CreatedAt Value
	// goverter:useZeroValueOnPointerInconsistency
	GotItem(queries.GetSecretsRow) webhooksv1.Secret
	Got([]queries.GetSecretsRow) []webhooksv1.Secret
}

// goverter:converter
// goverter:extend TimeToString
// goverter:output:file ./generated.go
// goverter:output:package github.com/andrestielau/web-of-hooks/webhooks/adapt/http/convert
type MessageConverter interface {
	// goverter:ignore ApplicationID EventTypeID EventID
	// goverter:useZeroValueOnPointerInconsistency
	NewItem(webhooksv1.NewMessage) queries.NewMessage
	New([]webhooksv1.NewMessage) []queries.NewMessage

	// goverter:map Uid Id
	// goverter:ignore EventTypeId EventId
	// goverter:useZeroValueOnPointerInconsistency
	ListItem(queries.ListMessagesRow) webhooksv1.Message
	List([]queries.ListMessagesRow) []webhooksv1.Message

	// goverter:map Uid Id
	// goverter:ignore EventTypeId EventId
	// goverter:useZeroValueOnPointerInconsistency
	CreatedItem(queries.CreateMessagesRow) webhooksv1.Message
	Created([]queries.CreateMessagesRow) []webhooksv1.Message
	
	// goverter:map Uid Id
	// goverter:ignore EventTypeId EventId
	// goverter:useZeroValueOnPointerInconsistency
	GotItem(queries.CreateMessagesRow) webhooksv1.Message
	Got([]queries.CreateMessagesRow) []webhooksv1.Message
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
