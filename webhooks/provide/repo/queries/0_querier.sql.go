// Code generated by pggen. DO NOT EDIT.

package queries

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"time"
)

// Querier is a typesafe Go interface backed by SQL queries.
type Querier interface {
	Now(ctx context.Context) (time.Time, error)

	// CreateApplications inserts applications into the database
	CreateApplications(ctx context.Context, applications []NewApplication) ([]Application, error)

	// DeleteApplications deletes application by uid
	DeleteApplications(ctx context.Context, ids []string) (pgconn.CommandTag, error)

	// GetApplications gets applications by id
	GetApplications(ctx context.Context, ids []string) ([]Application, error)

	// GetApplicationsByName gets applications by name
	GetApplicationsByName(ctx context.Context, names []string) ([]Application, error)

	// ListApplications lists registered applications
	ListApplications(ctx context.Context, params ListApplicationsParams) ([]Application, error)

	// DeleteAttempts deletes attempts by uid
	DeleteAttempts(ctx context.Context, ids []string) (pgconn.CommandTag, error)

	// GetAttempts gets attempts by id
	GetAttempts(ctx context.Context, ids []string) ([]MessageAttempt, error)

	// ListAttempts lists message attempts
	ListAttempts(ctx context.Context, params ListAttemptsParams) ([]MessageAttempt, error)

	// CreateEndpoints inserts endpoints into the database
	CreateEndpoints(ctx context.Context, endpoints []NewEndpoint) ([]EndpointDetails, error)

	// DeleteEndpoints deletes endpoints by uid
	DeleteEndpoints(ctx context.Context, ids []string) (pgconn.CommandTag, error)

	// GetEndpoints gets endpoints by id
	GetEndpoints(ctx context.Context, ids []string) ([]EndpointDetails, error)

	// GetEndpointsByUrl gets endpoints by url
	GetEndpointsByUrl(ctx context.Context, urls []string) ([]EndpointDetails, error)

	// ListEndpoints lists endpoints
	ListEndpoints(ctx context.Context, params ListEndpointsParams) ([]Endpoint, error)

	// ListApplicationEndpoints lists endpoints
	ListApplicationEndpoints(ctx context.Context, params ListApplicationEndpointsParams) ([]Endpoint, error)

	// CreateEventTypes inserts event types into the database
	CreateEventTypes(ctx context.Context, eventTypes []NewEventType) ([]EventType, error)

	// DeleteEventTypes deletes endpoints by uid
	DeleteEventTypes(ctx context.Context, keys []string) (pgconn.CommandTag, error)

	// GetEventTypes gets event-types by id
	GetEventTypes(ctx context.Context, ids []string) ([]EventType, error)

	// GetEventTypes gets event-types by id
	GetEventTypesByKeys(ctx context.Context, keys []string) ([]EventType, error)

	// ListEventTypes lists event-types
	ListEventTypes(ctx context.Context, params ListEventTypesParams) ([]EventType, error)

	// CreateMessages inserts messages into the database
	CreateMessages(ctx context.Context, messages []NewMessage) ([]MessageDetails, error)

	// DeleteMessages deletes messages by uid
	DeleteMessages(ctx context.Context, ids []string) (pgconn.CommandTag, error)

	// GetMessages gets messages by id
	GetMessages(ctx context.Context, ids []string) ([]MessageDetails, error)

	// ListMessages lists event-types
	ListMessages(ctx context.Context, params ListMessagesParams) ([]Message, error)

	// ListApplicationMessages lists event-types
	ListApplicationMessages(ctx context.Context, params ListApplicationMessagesParams) ([]Message, error)

	// CreateSecrets creates secrets
	CreateSecrets(ctx context.Context, secrets []NewSecret) ([]Secret, error)

	// GetSecrets gets secrets by id
	GetSecrets(ctx context.Context, ids []string) ([]Secret, error)

	// DeleteSecrets deletes secrets by uid
	DeleteSecrets(ctx context.Context, ids []string) (pgconn.CommandTag, error)

	// ListSecrets lists secrets
	ListSecrets(ctx context.Context, params ListSecretsParams) ([]Secret, error)

	// Update secret value by uid
	UpdateSecret(ctx context.Context, value string, uid string) (pgconn.CommandTag, error)

	// ListApplicationSecrets lists secrets of an application
	ListApplicationSecrets(ctx context.Context, applicationUid string) ([]Secret, error)
}

var _ Querier = &DBQuerier{}

type DBQuerier struct {
	conn  genericConn   // underlying Postgres transport to use
	types *typeResolver // resolve types by name
}

// genericConn is a connection like *pgx.Conn, pgx.Tx, or *pgxpool.Pool.
type genericConn interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
}

// NewQuerier creates a DBQuerier that implements Querier.
func NewQuerier(conn genericConn) *DBQuerier {
	return &DBQuerier{conn: conn, types: newTypeResolver()}
}

// Application represents the Postgres composite type "application".
type Application struct {
	ID        *int32       `json:"id"`
	Name      string       `json:"name"`
	Uid       string       `json:"uid"`
	TenantID  string       `json:"tenant_id"`
	RateLimit *int32       `json:"rate_limit"`
	Metadata  pgtype.JSONB `json:"metadata"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

// Endpoint represents the Postgres composite type "endpoint".
type Endpoint struct {
	ID            *int32       `json:"id"`
	Url           string       `json:"url"`
	Name          string       `json:"name"`
	ApplicationID *int32       `json:"application_id"`
	Uid           string       `json:"uid"`
	RateLimit     *int32       `json:"rate_limit"`
	Metadata      pgtype.JSONB `json:"metadata"`
	Disabled      *bool        `json:"disabled"`
	Description   string       `json:"description"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
}

// EndpointDetails represents the Postgres composite type "endpoint_details".
type EndpointDetails struct {
	Endpoint      Endpoint `json:"endpoint"`
	FilterTypeIds []string `json:"filter_type_ids"`
	Secret        string   `json:"secret"`
}

// EventType represents the Postgres composite type "event_type".
type EventType struct {
	ID        *int32    `json:"id"`
	Key       string    `json:"key"`
	Uid       string    `json:"uid"`
	CreatedAt time.Time `json:"created_at"`
}

// Message represents the Postgres composite type "message".
type Message struct {
	ID            *int32    `json:"id"`
	ApplicationID *int32    `json:"application_id"`
	EventTypeID   *int32    `json:"event_type_id"`
	Uid           string    `json:"uid"`
	CreatedAt     time.Time `json:"created_at"`
	EventID       string    `json:"event_id"`
	Payload       string    `json:"payload"`
}

// MessageAttempt represents the Postgres composite type "message_attempt".
type MessageAttempt struct {
	ID             *int32    `json:"id"`
	Uid            string    `json:"uid"`
	EndpointID     *int32    `json:"endpoint_id"`
	MessageID      *int32    `json:"message_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Status         *int32    `json:"status"`
	Retry          *int32    `json:"retry"`
	ResponseStatus *int32    `json:"response_status"`
	Response       string    `json:"response"`
}

// MessageDetails represents the Postgres composite type "message_details".
type MessageDetails struct {
	Message  Message          `json:"message"`
	Attempts []MessageAttempt `json:"attempts"`
}

// NewApplication represents the Postgres composite type "new_application".
type NewApplication struct {
	Name      string       `json:"name"`
	TenantID  string       `json:"tenant_id"`
	RateLimit *int32       `json:"rate_limit"`
	Metadata  pgtype.JSONB `json:"metadata"`
}

// NewEndpoint represents the Postgres composite type "new_endpoint".
type NewEndpoint struct {
	Url           string       `json:"url"`
	Name          string       `json:"name"`
	ApplicationID string       `json:"application_id"`
	SecretID      string       `json:"secret_id"`
	RateLimit     *int32       `json:"rate_limit"`
	Metadata      pgtype.JSONB `json:"metadata"`
	Description   string       `json:"description"`
	FilterTypeIds []string     `json:"filter_type_ids"`
	Channels      []string     `json:"channels"`
}

// NewEventType represents the Postgres composite type "new_event_type".
type NewEventType struct {
	Key string `json:"key"`
}

// NewMessage represents the Postgres composite type "new_message".
type NewMessage struct {
	ApplicationID string `json:"application_id"`
	EventTypeID   string `json:"event_type_id"`
	EventID       string `json:"event_id"`
	Payload       string `json:"payload"`
}

// NewSecret represents the Postgres composite type "new_secret".
type NewSecret struct {
	ApplicationID string `json:"application_id"`
	Value         string `json:"value"`
}

// Secret represents the Postgres composite type "secret".
type Secret struct {
	ID            *int32    `json:"id"`
	Uid           string    `json:"uid"`
	ApplicationID *int32    `json:"application_id"`
	Value         string    `json:"value"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// typeResolver looks up the pgtype.ValueTranscoder by Postgres type name.
type typeResolver struct {
	connInfo *pgtype.ConnInfo // types by Postgres type name
}

func newTypeResolver() *typeResolver {
	ci := pgtype.NewConnInfo()
	return &typeResolver{connInfo: ci}
}

// findValue find the OID, and pgtype.ValueTranscoder for a Postgres type name.
func (tr *typeResolver) findValue(name string) (uint32, pgtype.ValueTranscoder, bool) {
	typ, ok := tr.connInfo.DataTypeForName(name)
	if !ok {
		return 0, nil, false
	}
	v := pgtype.NewValue(typ.Value)
	return typ.OID, v.(pgtype.ValueTranscoder), true
}

// setValue sets the value of a ValueTranscoder to a value that should always
// work and panics if it fails.
func (tr *typeResolver) setValue(vt pgtype.ValueTranscoder, val interface{}) pgtype.ValueTranscoder {
	if err := vt.Set(val); err != nil {
		panic(fmt.Sprintf("set ValueTranscoder %T to %+v: %s", vt, val, err))
	}
	return vt
}

type compositeField struct {
	name       string                 // name of the field
	typeName   string                 // Postgres type name
	defaultVal pgtype.ValueTranscoder // default value to use
}

func (tr *typeResolver) newCompositeValue(name string, fields ...compositeField) pgtype.ValueTranscoder {
	if _, val, ok := tr.findValue(name); ok {
		return val
	}
	fs := make([]pgtype.CompositeTypeField, len(fields))
	vals := make([]pgtype.ValueTranscoder, len(fields))
	isBinaryOk := true
	for i, field := range fields {
		oid, val, ok := tr.findValue(field.typeName)
		if !ok {
			oid = unknownOID
			val = field.defaultVal
		}
		isBinaryOk = isBinaryOk && oid != unknownOID
		fs[i] = pgtype.CompositeTypeField{Name: field.name, OID: oid}
		vals[i] = val
	}
	// Okay to ignore error because it's only thrown when the number of field
	// names does not equal the number of ValueTranscoders.
	typ, _ := pgtype.NewCompositeTypeValues(name, fs, vals)
	if !isBinaryOk {
		return textPreferrer{ValueTranscoder: typ, typeName: name}
	}
	return typ
}

func (tr *typeResolver) newArrayValue(name, elemName string, defaultVal func() pgtype.ValueTranscoder) pgtype.ValueTranscoder {
	if _, val, ok := tr.findValue(name); ok {
		return val
	}
	elemOID, elemVal, ok := tr.findValue(elemName)
	elemValFunc := func() pgtype.ValueTranscoder {
		return pgtype.NewValue(elemVal).(pgtype.ValueTranscoder)
	}
	if !ok {
		elemOID = unknownOID
		elemValFunc = defaultVal
	}
	typ := pgtype.NewArrayType(name, elemOID, elemValFunc)
	if elemOID == unknownOID {
		return textPreferrer{ValueTranscoder: typ, typeName: name}
	}
	return typ
}

// newApplication creates a new pgtype.ValueTranscoder for the Postgres
// composite type 'application'.
func (tr *typeResolver) newApplication() pgtype.ValueTranscoder {
	return tr.newCompositeValue(
		"application",
		compositeField{name: "id", typeName: "int4", defaultVal: &pgtype.Int4{}},
		compositeField{name: "name", typeName: "text", defaultVal: &pgtype.Text{}},
		compositeField{name: "uid", typeName: "uuid", defaultVal: &pgtype.UUID{}},
		compositeField{name: "tenant_id", typeName: "text", defaultVal: &pgtype.Text{}},
		compositeField{name: "rate_limit", typeName: "int4", defaultVal: &pgtype.Int4{}},
		compositeField{name: "metadata", typeName: "jsonb", defaultVal: &pgtype.JSONB{}},
		compositeField{name: "created_at", typeName: "timestamptz", defaultVal: &pgtype.Timestamptz{}},
		compositeField{name: "updated_at", typeName: "timestamptz", defaultVal: &pgtype.Timestamptz{}},
	)
}

// newEndpointDetails creates a new pgtype.ValueTranscoder for the Postgres
// composite type 'endpoint_details'.
func (tr *typeResolver) newEndpointDetails() pgtype.ValueTranscoder {
	return tr.newCompositeValue(
		"endpoint_details",
		compositeField{name: "endpoint", typeName: "endpoint", defaultVal: tr.newEndpoint()},
		compositeField{name: "filter_type_ids", typeName: "_text", defaultVal: &pgtype.TextArray{}},
		compositeField{name: "secret", typeName: "text", defaultVal: &pgtype.Text{}},
	)
}

// newEndpoint creates a new pgtype.ValueTranscoder for the Postgres
// composite type 'endpoint'.
func (tr *typeResolver) newEndpoint() pgtype.ValueTranscoder {
	return tr.newCompositeValue(
		"endpoint",
		compositeField{name: "id", typeName: "int4", defaultVal: &pgtype.Int4{}},
		compositeField{name: "url", typeName: "text", defaultVal: &pgtype.Text{}},
		compositeField{name: "name", typeName: "text", defaultVal: &pgtype.Text{}},
		compositeField{name: "application_id", typeName: "int4", defaultVal: &pgtype.Int4{}},
		compositeField{name: "uid", typeName: "uuid", defaultVal: &pgtype.UUID{}},
		compositeField{name: "rate_limit", typeName: "int4", defaultVal: &pgtype.Int4{}},
		compositeField{name: "metadata", typeName: "jsonb", defaultVal: &pgtype.JSONB{}},
		compositeField{name: "disabled", typeName: "bool", defaultVal: &pgtype.Bool{}},
		compositeField{name: "description", typeName: "text", defaultVal: &pgtype.Text{}},
		compositeField{name: "created_at", typeName: "timestamptz", defaultVal: &pgtype.Timestamptz{}},
		compositeField{name: "updated_at", typeName: "timestamptz", defaultVal: &pgtype.Timestamptz{}},
	)
}

// newEventType creates a new pgtype.ValueTranscoder for the Postgres
// composite type 'event_type'.
func (tr *typeResolver) newEventType() pgtype.ValueTranscoder {
	return tr.newCompositeValue(
		"event_type",
		compositeField{name: "id", typeName: "int4", defaultVal: &pgtype.Int4{}},
		compositeField{name: "key", typeName: "text", defaultVal: &pgtype.Text{}},
		compositeField{name: "uid", typeName: "uuid", defaultVal: &pgtype.UUID{}},
		compositeField{name: "created_at", typeName: "timestamptz", defaultVal: &pgtype.Timestamptz{}},
	)
}

// newMessageAttempt creates a new pgtype.ValueTranscoder for the Postgres
// composite type 'message_attempt'.
func (tr *typeResolver) newMessageAttempt() pgtype.ValueTranscoder {
	return tr.newCompositeValue(
		"message_attempt",
		compositeField{name: "id", typeName: "int4", defaultVal: &pgtype.Int4{}},
		compositeField{name: "uid", typeName: "uuid", defaultVal: &pgtype.UUID{}},
		compositeField{name: "endpoint_id", typeName: "int4", defaultVal: &pgtype.Int4{}},
		compositeField{name: "message_id", typeName: "int4", defaultVal: &pgtype.Int4{}},
		compositeField{name: "created_at", typeName: "timestamptz", defaultVal: &pgtype.Timestamptz{}},
		compositeField{name: "updated_at", typeName: "timestamptz", defaultVal: &pgtype.Timestamptz{}},
		compositeField{name: "status", typeName: "int4", defaultVal: &pgtype.Int4{}},
		compositeField{name: "retry", typeName: "int4", defaultVal: &pgtype.Int4{}},
		compositeField{name: "response_status", typeName: "int4", defaultVal: &pgtype.Int4{}},
		compositeField{name: "response", typeName: "text", defaultVal: &pgtype.Text{}},
	)
}

// newMessageDetails creates a new pgtype.ValueTranscoder for the Postgres
// composite type 'message_details'.
func (tr *typeResolver) newMessageDetails() pgtype.ValueTranscoder {
	return tr.newCompositeValue(
		"message_details",
		compositeField{name: "message", typeName: "message", defaultVal: tr.newMessage()},
		compositeField{name: "attempts", typeName: "_message_attempt", defaultVal: tr.newMessageAttemptArray()},
	)
}

// newMessage creates a new pgtype.ValueTranscoder for the Postgres
// composite type 'message'.
func (tr *typeResolver) newMessage() pgtype.ValueTranscoder {
	return tr.newCompositeValue(
		"message",
		compositeField{name: "id", typeName: "int4", defaultVal: &pgtype.Int4{}},
		compositeField{name: "application_id", typeName: "int4", defaultVal: &pgtype.Int4{}},
		compositeField{name: "event_type_id", typeName: "int4", defaultVal: &pgtype.Int4{}},
		compositeField{name: "uid", typeName: "uuid", defaultVal: &pgtype.UUID{}},
		compositeField{name: "created_at", typeName: "timestamptz", defaultVal: &pgtype.Timestamptz{}},
		compositeField{name: "event_id", typeName: "text", defaultVal: &pgtype.Text{}},
		compositeField{name: "payload", typeName: "text", defaultVal: &pgtype.Text{}},
	)
}

// newNewApplication creates a new pgtype.ValueTranscoder for the Postgres
// composite type 'new_application'.
func (tr *typeResolver) newNewApplication() pgtype.ValueTranscoder {
	return tr.newCompositeValue(
		"new_application",
		compositeField{name: "name", typeName: "text", defaultVal: &pgtype.Text{}},
		compositeField{name: "tenant_id", typeName: "text", defaultVal: &pgtype.Text{}},
		compositeField{name: "rate_limit", typeName: "int4", defaultVal: &pgtype.Int4{}},
		compositeField{name: "metadata", typeName: "jsonb", defaultVal: &pgtype.JSONB{}},
	)
}

// newNewApplicationRaw returns all composite fields for the Postgres composite
// type 'new_application' as a slice of interface{} to encode query parameters.
func (tr *typeResolver) newNewApplicationRaw(v NewApplication) []interface{} {
	return []interface{}{
		v.Name,
		v.TenantID,
		v.RateLimit,
		v.Metadata,
	}
}

// newNewEndpoint creates a new pgtype.ValueTranscoder for the Postgres
// composite type 'new_endpoint'.
func (tr *typeResolver) newNewEndpoint() pgtype.ValueTranscoder {
	return tr.newCompositeValue(
		"new_endpoint",
		compositeField{name: "url", typeName: "text", defaultVal: &pgtype.Text{}},
		compositeField{name: "name", typeName: "text", defaultVal: &pgtype.Text{}},
		compositeField{name: "application_id", typeName: "uuid", defaultVal: &pgtype.UUID{}},
		compositeField{name: "secret_id", typeName: "uuid", defaultVal: &pgtype.UUID{}},
		compositeField{name: "rate_limit", typeName: "int4", defaultVal: &pgtype.Int4{}},
		compositeField{name: "metadata", typeName: "jsonb", defaultVal: &pgtype.JSONB{}},
		compositeField{name: "description", typeName: "text", defaultVal: &pgtype.Text{}},
		compositeField{name: "filter_type_ids", typeName: "_text", defaultVal: &pgtype.TextArray{}},
		compositeField{name: "channels", typeName: "_text", defaultVal: &pgtype.TextArray{}},
	)
}

// newNewEndpointRaw returns all composite fields for the Postgres composite
// type 'new_endpoint' as a slice of interface{} to encode query parameters.
func (tr *typeResolver) newNewEndpointRaw(v NewEndpoint) []interface{} {
	return []interface{}{
		v.Url,
		v.Name,
		v.ApplicationID,
		v.SecretID,
		v.RateLimit,
		v.Metadata,
		v.Description,
		v.FilterTypeIds,
		v.Channels,
	}
}

// newNewEventType creates a new pgtype.ValueTranscoder for the Postgres
// composite type 'new_event_type'.
func (tr *typeResolver) newNewEventType() pgtype.ValueTranscoder {
	return tr.newCompositeValue(
		"new_event_type",
		compositeField{name: "key", typeName: "text", defaultVal: &pgtype.Text{}},
	)
}

// newNewEventTypeRaw returns all composite fields for the Postgres composite
// type 'new_event_type' as a slice of interface{} to encode query parameters.
func (tr *typeResolver) newNewEventTypeRaw(v NewEventType) []interface{} {
	return []interface{}{
		v.Key,
	}
}

// newNewMessage creates a new pgtype.ValueTranscoder for the Postgres
// composite type 'new_message'.
func (tr *typeResolver) newNewMessage() pgtype.ValueTranscoder {
	return tr.newCompositeValue(
		"new_message",
		compositeField{name: "application_id", typeName: "uuid", defaultVal: &pgtype.UUID{}},
		compositeField{name: "event_type_id", typeName: "uuid", defaultVal: &pgtype.UUID{}},
		compositeField{name: "event_id", typeName: "text", defaultVal: &pgtype.Text{}},
		compositeField{name: "payload", typeName: "text", defaultVal: &pgtype.Text{}},
	)
}

// newNewMessageRaw returns all composite fields for the Postgres composite
// type 'new_message' as a slice of interface{} to encode query parameters.
func (tr *typeResolver) newNewMessageRaw(v NewMessage) []interface{} {
	return []interface{}{
		v.ApplicationID,
		v.EventTypeID,
		v.EventID,
		v.Payload,
	}
}

// newNewSecret creates a new pgtype.ValueTranscoder for the Postgres
// composite type 'new_secret'.
func (tr *typeResolver) newNewSecret() pgtype.ValueTranscoder {
	return tr.newCompositeValue(
		"new_secret",
		compositeField{name: "application_id", typeName: "uuid", defaultVal: &pgtype.UUID{}},
		compositeField{name: "value", typeName: "text", defaultVal: &pgtype.Text{}},
	)
}

// newNewSecretRaw returns all composite fields for the Postgres composite
// type 'new_secret' as a slice of interface{} to encode query parameters.
func (tr *typeResolver) newNewSecretRaw(v NewSecret) []interface{} {
	return []interface{}{
		v.ApplicationID,
		v.Value,
	}
}

// newSecret creates a new pgtype.ValueTranscoder for the Postgres
// composite type 'secret'.
func (tr *typeResolver) newSecret() pgtype.ValueTranscoder {
	return tr.newCompositeValue(
		"secret",
		compositeField{name: "id", typeName: "int4", defaultVal: &pgtype.Int4{}},
		compositeField{name: "uid", typeName: "uuid", defaultVal: &pgtype.UUID{}},
		compositeField{name: "application_id", typeName: "int4", defaultVal: &pgtype.Int4{}},
		compositeField{name: "value", typeName: "text", defaultVal: &pgtype.Text{}},
		compositeField{name: "created_at", typeName: "timestamptz", defaultVal: &pgtype.Timestamptz{}},
		compositeField{name: "updated_at", typeName: "timestamptz", defaultVal: &pgtype.Timestamptz{}},
	)
}

// newMessageAttemptArray creates a new pgtype.ValueTranscoder for the Postgres
// '_message_attempt' array type.
func (tr *typeResolver) newMessageAttemptArray() pgtype.ValueTranscoder {
	return tr.newArrayValue("_message_attempt", "message_attempt", tr.newMessageAttempt)
}

// newNewApplicationArray creates a new pgtype.ValueTranscoder for the Postgres
// '_new_application' array type.
func (tr *typeResolver) newNewApplicationArray() pgtype.ValueTranscoder {
	return tr.newArrayValue("_new_application", "new_application", tr.newNewApplication)
}

// newNewApplicationArrayInit creates an initialized pgtype.ValueTranscoder for the
// Postgres array type '_new_application' to encode query parameters.
func (tr *typeResolver) newNewApplicationArrayInit(ps []NewApplication) pgtype.ValueTranscoder {
	dec := tr.newNewApplicationArray()
	if err := dec.Set(tr.newNewApplicationArrayRaw(ps)); err != nil {
		panic("encode []NewApplication: " + err.Error()) // should always succeed
	}
	return textPreferrer{ValueTranscoder: dec, typeName: "_new_application"}
}

// newNewApplicationArrayRaw returns all elements for the Postgres array type '_new_application'
// as a slice of interface{} for use with the pgtype.Value Set method.
func (tr *typeResolver) newNewApplicationArrayRaw(vs []NewApplication) []interface{} {
	elems := make([]interface{}, len(vs))
	for i, v := range vs {
		elems[i] = tr.newNewApplicationRaw(v)
	}
	return elems
}

// newNewEndpointArray creates a new pgtype.ValueTranscoder for the Postgres
// '_new_endpoint' array type.
func (tr *typeResolver) newNewEndpointArray() pgtype.ValueTranscoder {
	return tr.newArrayValue("_new_endpoint", "new_endpoint", tr.newNewEndpoint)
}

// newNewEndpointArrayInit creates an initialized pgtype.ValueTranscoder for the
// Postgres array type '_new_endpoint' to encode query parameters.
func (tr *typeResolver) newNewEndpointArrayInit(ps []NewEndpoint) pgtype.ValueTranscoder {
	dec := tr.newNewEndpointArray()
	if err := dec.Set(tr.newNewEndpointArrayRaw(ps)); err != nil {
		panic("encode []NewEndpoint: " + err.Error()) // should always succeed
	}
	return textPreferrer{ValueTranscoder: dec, typeName: "_new_endpoint"}
}

// newNewEndpointArrayRaw returns all elements for the Postgres array type '_new_endpoint'
// as a slice of interface{} for use with the pgtype.Value Set method.
func (tr *typeResolver) newNewEndpointArrayRaw(vs []NewEndpoint) []interface{} {
	elems := make([]interface{}, len(vs))
	for i, v := range vs {
		elems[i] = tr.newNewEndpointRaw(v)
	}
	return elems
}

// newNewEventTypeArray creates a new pgtype.ValueTranscoder for the Postgres
// '_new_event_type' array type.
func (tr *typeResolver) newNewEventTypeArray() pgtype.ValueTranscoder {
	return tr.newArrayValue("_new_event_type", "new_event_type", tr.newNewEventType)
}

// newNewEventTypeArrayInit creates an initialized pgtype.ValueTranscoder for the
// Postgres array type '_new_event_type' to encode query parameters.
func (tr *typeResolver) newNewEventTypeArrayInit(ps []NewEventType) pgtype.ValueTranscoder {
	dec := tr.newNewEventTypeArray()
	if err := dec.Set(tr.newNewEventTypeArrayRaw(ps)); err != nil {
		panic("encode []NewEventType: " + err.Error()) // should always succeed
	}
	return textPreferrer{ValueTranscoder: dec, typeName: "_new_event_type"}
}

// newNewEventTypeArrayRaw returns all elements for the Postgres array type '_new_event_type'
// as a slice of interface{} for use with the pgtype.Value Set method.
func (tr *typeResolver) newNewEventTypeArrayRaw(vs []NewEventType) []interface{} {
	elems := make([]interface{}, len(vs))
	for i, v := range vs {
		elems[i] = tr.newNewEventTypeRaw(v)
	}
	return elems
}

// newNewMessageArray creates a new pgtype.ValueTranscoder for the Postgres
// '_new_message' array type.
func (tr *typeResolver) newNewMessageArray() pgtype.ValueTranscoder {
	return tr.newArrayValue("_new_message", "new_message", tr.newNewMessage)
}

// newNewMessageArrayInit creates an initialized pgtype.ValueTranscoder for the
// Postgres array type '_new_message' to encode query parameters.
func (tr *typeResolver) newNewMessageArrayInit(ps []NewMessage) pgtype.ValueTranscoder {
	dec := tr.newNewMessageArray()
	if err := dec.Set(tr.newNewMessageArrayRaw(ps)); err != nil {
		panic("encode []NewMessage: " + err.Error()) // should always succeed
	}
	return textPreferrer{ValueTranscoder: dec, typeName: "_new_message"}
}

// newNewMessageArrayRaw returns all elements for the Postgres array type '_new_message'
// as a slice of interface{} for use with the pgtype.Value Set method.
func (tr *typeResolver) newNewMessageArrayRaw(vs []NewMessage) []interface{} {
	elems := make([]interface{}, len(vs))
	for i, v := range vs {
		elems[i] = tr.newNewMessageRaw(v)
	}
	return elems
}

// newNewSecretArray creates a new pgtype.ValueTranscoder for the Postgres
// '_new_secret' array type.
func (tr *typeResolver) newNewSecretArray() pgtype.ValueTranscoder {
	return tr.newArrayValue("_new_secret", "new_secret", tr.newNewSecret)
}

// newNewSecretArrayInit creates an initialized pgtype.ValueTranscoder for the
// Postgres array type '_new_secret' to encode query parameters.
func (tr *typeResolver) newNewSecretArrayInit(ps []NewSecret) pgtype.ValueTranscoder {
	dec := tr.newNewSecretArray()
	if err := dec.Set(tr.newNewSecretArrayRaw(ps)); err != nil {
		panic("encode []NewSecret: " + err.Error()) // should always succeed
	}
	return textPreferrer{ValueTranscoder: dec, typeName: "_new_secret"}
}

// newNewSecretArrayRaw returns all elements for the Postgres array type '_new_secret'
// as a slice of interface{} for use with the pgtype.Value Set method.
func (tr *typeResolver) newNewSecretArrayRaw(vs []NewSecret) []interface{} {
	elems := make([]interface{}, len(vs))
	for i, v := range vs {
		elems[i] = tr.newNewSecretRaw(v)
	}
	return elems
}

const nowSQL = `SELECT now();`

// Now implements Querier.Now.
func (q *DBQuerier) Now(ctx context.Context) (time.Time, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "Now")
	row := q.conn.QueryRow(ctx, nowSQL)
	var item time.Time
	if err := row.Scan(&item); err != nil {
		return item, fmt.Errorf("query Now: %w", err)
	}
	return item, nil
}

// textPreferrer wraps a pgtype.ValueTranscoder and sets the preferred encoding
// format to text instead binary (the default). pggen uses the text format
// when the OID is unknownOID because the binary format requires the OID.
// Typically occurs for unregistered types.
type textPreferrer struct {
	pgtype.ValueTranscoder
	typeName string
}

// PreferredParamFormat implements pgtype.ParamFormatPreferrer.
func (t textPreferrer) PreferredParamFormat() int16 { return pgtype.TextFormatCode }

func (t textPreferrer) NewTypeValue() pgtype.Value {
	return textPreferrer{ValueTranscoder: pgtype.NewValue(t.ValueTranscoder).(pgtype.ValueTranscoder), typeName: t.typeName}
}

func (t textPreferrer) TypeName() string {
	return t.typeName
}

// unknownOID means we don't know the OID for a type. This is okay for decoding
// because pgx call DecodeText or DecodeBinary without requiring the OID. For
// encoding parameters, pggen uses textPreferrer if the OID is unknown.
const unknownOID = 0
