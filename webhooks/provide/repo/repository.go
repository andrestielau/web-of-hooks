package repo

import (
	"context"

	"woh/webhooks"
	"woh/webhooks/provide/repo/convert"
	"woh/webhooks/provide/repo/queries"
)

type Repository struct {
	Convert convert.Converter
	Querier queries.Querier
}

func (r *Repository) CreateApplications(ctx context.Context, applications []webhooks.NewApplication) ([]webhooks.Application, error) {
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

func (r *Repository) GetApplications(ctx context.Context, ids []string) ([]webhooks.Application, error) {
	res, err := r.Querier.GetApplications(ctx, ids)
	if err != nil {
		return nil, err
	}
	return r.Convert.Applications(res), nil
}

func (r *Repository) ListApplications(ctx context.Context, query webhooks.ApplicationQuery) ([]webhooks.Application, error) {
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
	res, err := r.Querier.ListAttempts(ctx, r.Convert.AttemptQuery(query))
	if err != nil {
		return nil, err
	}
	return r.Convert.Attempts(res), nil
}

func (r *Repository) CreateEndpoints(ctx context.Context, endpoints []webhooks.NewEndpoint) ([]webhooks.EndpointDetails, error) {
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

func (r *Repository) ListEndpoints(ctx context.Context, query webhooks.EndpointQuery) ([]webhooks.Endpoint, error) {
	res, err := r.Querier.ListEndpoints(ctx, r.Convert.EndpointQuery(query))
	if err != nil {
		return nil, err
	}
	return r.Convert.Endpoints(res), nil
}

func (r *Repository) ListApplicationEndpoints(ctx context.Context, application_uid string, query webhooks.EndpointQuery) ([]webhooks.Endpoint, error) {
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

func (r *Repository) ListEventTypes(ctx context.Context, query webhooks.EventTypeQuery) ([]webhooks.EventType, error) {
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
	res, err := r.Querier.ListMessages(ctx, r.Convert.MessageQuery(query))
	if err != nil {
		return nil, err
	}
	return r.Convert.Messages(res), nil
}
func (r *Repository) ListApplicationMessages(ctx context.Context, application_uid string, query webhooks.MessageQuery) ([]webhooks.Message, error) {
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
