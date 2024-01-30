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
	return nil, nil
}

func (r *Repository) DeleteApplications(ctx context.Context, ids []string) error {
	return nil
}

func (r *Repository) GetApplications(ctx context.Context, ids []string) ([]webhooks.Application, error) {
	return nil, nil
}

func (r *Repository) ListApplications(ctx context.Context, limit int, offset int) ([]webhooks.Application, error) {
	return nil, nil
}
func (r *Repository) DeleteAttempts(ctx context.Context, ids []string) error {
	return nil
}

func (r *Repository) GetAttempts(ctx context.Context, ids []string) ([]webhooks.Attempt, error) {
	return nil, nil
}

func (r *Repository) ListAttempts(ctx context.Context, limit int, offset int) ([]webhooks.Attempt, error) {
	return nil, nil
}

func (r *Repository) CreateEndpoints(ctx context.Context, endpoints []webhooks.NewEndpoint) ([]webhooks.Endpoint, error) {
	return nil, nil
}
func (r *Repository) DeleteEndpoints(ctx context.Context, ids []string) error {
	return nil
}

func (r *Repository) GetEndpoints(ctx context.Context, ids []string) ([]webhooks.Endpoint, error) {
	return nil, nil
}

func (r *Repository) ListEndpoints(ctx context.Context, limit int, offset int) ([]webhooks.Endpoint, error) {
	return nil, nil
}

func (r *Repository) CreateEventTypes(ctx context.Context, eventTypes []webhooks.NewEventType) ([]webhooks.EventType, error) {
	return nil, nil
}

func (r *Repository) DeleteEventTypes(ctx context.Context, keys []string) error {
	return nil
}

func (r *Repository) GetEventTypes(ctx context.Context, ids []string) ([]webhooks.EventType, error) {
	return nil, nil
}

func (r *Repository) ListEventTypes(ctx context.Context, limit int, offset int) ([]webhooks.EventType, error) {
	return nil, nil
}

func (r *Repository) CreateMessages(ctx context.Context, messages []webhooks.NewMessage) ([]webhooks.Message, error) {
	return nil, nil
}

func (r *Repository) DeleteMessages(ctx context.Context, ids []string) error {
	return nil
}

func (r *Repository) GetMessages(ctx context.Context, ids []string) ([]webhooks.Message, error) {
	return nil, nil
}
func (r *Repository) ListMessages(ctx context.Context, limit int, offset int) ([]webhooks.Message, error) {
	return nil, nil
}
func (r *Repository) CreateSecrets(ctx context.Context, secrets []webhooks.NewSecret) ([]webhooks.Secret, error) {
	return nil, nil
}
func (r *Repository) GetSecrets(ctx context.Context, ids []string) ([]webhooks.Secret, error) {
	return nil, nil
}
func (r *Repository) DeleteSecrets(ctx context.Context, ids []string) error {
	return nil
}
func (r *Repository) ListSecrets(ctx context.Context, limit int, offset int) ([]webhooks.Secret, error) {
	return nil, nil
}
