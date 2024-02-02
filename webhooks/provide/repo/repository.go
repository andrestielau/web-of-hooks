package repo

import (
	"context"
	"encoding/json"

	"woh/webhooks"
	"woh/webhooks/provide/repo/convert"
	"woh/webhooks/provide/repo/queries"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

const DefaultLimit = 100

type Repository struct {
	Convert convert.Converter
	Querier queries.Querier
}

func (r *Repository) CreateApplications(ctx context.Context, applications []webhooks.NewApplication) ([]webhooks.Application, error) {
	for i, app := range applications {
		if app.RateLimit == nil {
			app.RateLimit = lo.ToPtr[int32](DefaultLimit)
		}
		applications[i] = app
	}
	res, err := r.Querier.CreateApplications(ctx, r.Convert.NewApplications(applications))
	if err != nil {
		return nil, err
	}
	return r.Convert.Applications(res), nil
}

func (r *Repository) DeleteApplications(ctx context.Context, ids []string) error {
	_, err := r.Querier.DeleteApplications(ctx, ids)
	return err
}

func (r *Repository) GetApplications(ctx context.Context, ids []string) ([]webhooks.ApplicationDetails, error) {
	res, err := r.Querier.GetApplications(ctx, ids)
	if err != nil {
		return nil, err
	}
	return r.Convert.ApplicationDetails(res), nil
}

func (r *Repository) GetApplicationsByName(ctx context.Context, names []string) ([]webhooks.Application, error) {
	res, err := r.Querier.GetApplicationsByName(ctx, names)
	if err != nil {
		return nil, err
	}
	return r.Convert.Applications(res), nil
}

func (r *Repository) ListApplications(ctx context.Context, query webhooks.ApplicationQuery) ([]webhooks.Application, error) {
	query.Limit = defaultLimit(query.Limit)
	res, err := r.Querier.ListApplications(ctx, r.Convert.ApplicationQuery(query))
	if err != nil {
		return nil, err
	}
	return r.Convert.Applications(res), nil
}
func (r *Repository) DeleteAttempts(ctx context.Context, ids []string) error {
	_, err := r.Querier.DeleteApplications(ctx, ids)
	return err
}

func (r *Repository) GetAttempts(ctx context.Context, ids []string) ([]webhooks.Attempt, error) {
	res, err := r.Querier.GetAttempts(ctx, ids)
	if err != nil {
		return nil, err
	}
	return r.Convert.Attempts(res), nil
}

func (r *Repository) ListAttempts(ctx context.Context, query webhooks.AttemptQuery) ([]webhooks.Attempt, error) {
	query.Limit = defaultLimit(query.Limit)
	res, err := r.Querier.ListAttempts(ctx, r.Convert.AttemptQuery(query))
	if err != nil {
		return nil, err
	}
	return r.Convert.Attempts(res), nil
}

func (r *Repository) CreateEndpoints(ctx context.Context, endpoints []webhooks.NewEndpoint) ([]webhooks.EndpointDetails, error) {
	for i, endpoint := range endpoints {
		res, _ := r.CreateSecrets(ctx, []webhooks.NewSecret{{
			ApplicationID: endpoint.ApplicationID,
			Value:         uuid.NewString(),
		}})
		endpoints[i].SecretId = res[0].Uid
		endpoints[i].RateLimit = lo.ToPtr[int32](10)
	}
	res, err := r.Querier.CreateEndpoints(ctx, r.Convert.NewEndpoints(endpoints))
	if err != nil {
		return nil, err
	}
	return r.Convert.EndpointDetails(res), nil
}
func (r *Repository) DeleteEndpoints(ctx context.Context, ids []string) error {
	_, err := r.Querier.DeleteEndpoints(ctx, ids)
	return err
}

func (r *Repository) GetEndpoints(ctx context.Context, ids []string) ([]webhooks.EndpointDetails, error) {
	res, err := r.Querier.GetEndpoints(ctx, ids)
	if err != nil {
		return nil, err
	}
	return r.Convert.EndpointDetails(res), nil
}

func (r *Repository) GetEndpointsByUrl(ctx context.Context, urls []string) ([]webhooks.EndpointDetails, error) {
	res, err := r.Querier.GetEndpointsByUrl(ctx, urls)
	if err != nil {
		return nil, err
	}
	return r.Convert.EndpointDetails(res), nil
}

func (r *Repository) GetEndpointsByTenantAndEventTypes(ctx context.Context, tenantId string, eventTypeKeys []string) ([]webhooks.EndpointDetails, error) {
	res, err := r.Querier.GetEndpointsByTenantAndEventTypes(ctx, tenantId, eventTypeKeys)
	if err != nil {
		return nil, err
	}
	return r.Convert.EndpointDetails(res), nil
}

func (r *Repository) ListEndpoints(ctx context.Context, query webhooks.EndpointQuery) ([]webhooks.Endpoint, error) {
	query.Limit = defaultLimit(query.Limit)
	res, err := r.Querier.ListEndpoints(ctx, r.Convert.EndpointQuery(query))
	if err != nil {
		return nil, err
	}
	return r.Convert.Endpoints(res), nil
}

func (r *Repository) ListApplicationEndpoints(ctx context.Context, application_uid string, query webhooks.EndpointQuery) ([]webhooks.Endpoint, error) {
	query.Limit = defaultLimit(query.Limit)
	q := r.Convert.ApplicationEndpointQuery(query)
	q.ApplicationUid = application_uid
	res, err := r.Querier.ListApplicationEndpoints(ctx, q)
	if err != nil {
		return nil, err
	}
	return r.Convert.Endpoints(res), nil
}

func (r *Repository) CreateEventTypes(ctx context.Context, eventTypes []webhooks.NewEventType) ([]webhooks.EventType, error) {
	res, err := r.Querier.CreateEventTypes(ctx, r.Convert.NewEventTypes(eventTypes))
	if err != nil {
		return nil, err
	}
	return r.Convert.EventTypes(res), nil
}

func (r *Repository) DeleteEventTypes(ctx context.Context, ids []string) error {
	_, err := r.Querier.DeleteEventTypes(ctx, ids)
	return err
}

func (r *Repository) GetEventTypes(ctx context.Context, ids []string) ([]webhooks.EventType, error) {
	res, err := r.Querier.GetEventTypes(ctx, ids)
	if err != nil {
		return nil, err
	}
	return r.Convert.EventTypes(res), nil
}

func (r *Repository) GetEventTypesByKeys(ctx context.Context, keys []string) ([]webhooks.EventType, error) {
	res, err := r.Querier.GetEventTypesByKeys(ctx, keys)
	if err != nil {
		return nil, err
	}
	return r.Convert.EventTypes(res), nil
}

func (r *Repository) ListEventTypes(ctx context.Context, query webhooks.EventTypeQuery) ([]webhooks.EventType, error) {
	query.Limit = defaultLimit(query.Limit)
	res, err := r.Querier.ListEventTypes(ctx, r.Convert.EventTypeQuery(query))
	if err != nil {
		return nil, err
	}
	return r.Convert.EventTypes(res), nil
}

func (r *Repository) CreateMessages(ctx context.Context, messages []webhooks.NewMessage) ([]webhooks.MessageDetails, error) {
	res, err := r.Querier.CreateMessages(ctx, r.Convert.NewMessages(messages))
	if err != nil {
		return nil, err
	}
	return r.Convert.MessageDetails(res), nil
}

func (r *Repository) DeleteMessages(ctx context.Context, ids []string) error {
	_, err := r.Querier.DeleteMessages(ctx, ids)
	return err
}

func (r *Repository) GetMessages(ctx context.Context, ids []string) ([]webhooks.MessageDetails, error) {
	res, err := r.Querier.GetMessages(ctx, ids)
	if err != nil {
		return nil, err
	}
	return r.Convert.MessageDetails(res), nil
}
func (r *Repository) ListMessages(ctx context.Context, query webhooks.MessageQuery) ([]webhooks.Message, error) {
	query.Limit = defaultLimit(query.Limit)
	res, err := r.Querier.ListMessages(ctx, r.Convert.MessageQuery(query))
	if err != nil {
		return nil, err
	}
	return r.Convert.Messages(res), nil
}
func (r *Repository) ListApplicationMessages(ctx context.Context, application_uid string, query webhooks.MessageQuery) ([]webhooks.Message, error) {
	query.Limit = defaultLimit(query.Limit)
	q := r.Convert.ApplicationMessageQuery(query)
	q.ApplicationUid = application_uid
	res, err := r.Querier.ListApplicationMessages(ctx, q)
	if err != nil {
		return nil, err
	}
	return r.Convert.Messages(res), nil
}
func (r *Repository) CreateSecrets(ctx context.Context, secrets []webhooks.NewSecret) ([]webhooks.Secret, error) {
	res, err := r.Querier.CreateSecrets(ctx, r.Convert.NewSecrets(secrets))
	if err != nil {
		return nil, err
	}
	return r.Convert.Secrets(res), nil
}
func (r *Repository) GetSecrets(ctx context.Context, ids []string) ([]webhooks.Secret, error) {
	res, err := r.Querier.GetSecrets(ctx, ids)
	if err != nil {
		return nil, err
	}
	return r.Convert.Secrets(res), nil
}
func (r *Repository) DeleteSecrets(ctx context.Context, ids []string) error {
	_, err := r.Querier.DeleteSecrets(ctx, ids)
	return err
}
func (r *Repository) ListSecrets(ctx context.Context, query webhooks.SecretQuery) ([]webhooks.Secret, error) {
	query.Limit = defaultLimit(query.Limit)
	res, err := r.Querier.ListSecrets(ctx, r.Convert.SecretQuery(query))
	if err != nil {
		return nil, err
	}
	return r.Convert.Secrets(res), nil
}

func (r *Repository) ListApplicationSecrets(ctx context.Context, application_uid string) ([]webhooks.Secret, error) {
	res, err := r.Querier.ListApplicationSecrets(ctx, application_uid)
	if err != nil {
		return nil, err
	}
	return r.Convert.Secrets(res), nil
}

func (r *Repository) EmitEvent(ctx context.Context, event webhooks.NewEvent) ([]webhooks.Message, error) {
	eventTypes, err := r.Querier.GetEventTypesByKeys(ctx, event.EventTypeKeys)
	if err != nil {
		return nil, err
	}
	messages := []queries.Message{}

	for _, eventType := range eventTypes {

		eventId := uuid.NewString()
		payload := webhooks.Payload{
			EventTypeKey: eventType.Key,
			ReferenceID:  event.ReferenceID,
		}
		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}

		// TODO: check if this is the correct way to do this
		apps, err := r.Querier.ListApplications(ctx, queries.ListApplicationsParams{
			Limit:    1000,
			TenantID: event.TenantId,
		})
		if err != nil {
			return nil, err
		}

		for _, app := range apps {
			newMessage := webhooks.NewMessage{
				ApplicationID: app.Uid,
				EventTypeID:   eventType.Uid,
				EventID:       eventId,
				Payload:       string(jsonPayload),
			}
			messageDetails, err := r.Querier.CreateMessages(ctx, r.Convert.NewMessages([]webhooks.NewMessage{newMessage}))
			if err != nil {
				return nil, err
			}

			messages = append(messages, lo.Map(messageDetails, func(detail queries.MessageDetails, _ int) queries.Message {
				return detail.Message
			})...)
		}
	}

	return r.Convert.Messages(messages), nil
}
func defaultLimit(l int) int {
	if l == 0 {
		l = DefaultLimit
	}
	return l
}
