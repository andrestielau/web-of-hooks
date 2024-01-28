package grpc

import (
	"context"
	webhooks "woh/webhooks"
	webhooksv1 "woh/webhooks/adapt/grpc/v1"
)

type Handler struct {
	Repo    webhooks.Repository
	Secrets webhooks.Secrets
}

var _ webhooksv1.WebHookServiceServer = &Handler{}

// CreateApps implements webhooksv1.WebHookServiceServer.
func (*Handler) CreateApps(context.Context, *webhooksv1.CreateAppsRequest) (*webhooksv1.CreateAppsResponse, error) {
	panic("unimplemented")
}

// CreateEndpoints implements webhooksv1.WebHookServiceServer.
func (*Handler) CreateEndpoints(context.Context, *webhooksv1.CreateEndpointsRequest) (*webhooksv1.CreateEndpointsResponse, error) {
	panic("unimplemented")
}

// CreateMessages implements webhooksv1.WebHookServiceServer.
func (*Handler) CreateMessages(context.Context, *webhooksv1.CreateMessagesRequest) (*webhooksv1.CreateMessagesResponse, error) {
	panic("unimplemented")
}

// DeleteApps implements webhooksv1.WebHookServiceServer.
func (*Handler) DeleteApps(context.Context, *webhooksv1.DeleteAppsRequest) (*webhooksv1.DeleteAppsResponse, error) {
	panic("unimplemented")
}

// DeleteEndpoints implements webhooksv1.WebHookServiceServer.
func (*Handler) DeleteEndpoints(context.Context, *webhooksv1.DeleteEndpointsRequest) (*webhooksv1.DeleteEndpointsResponse, error) {
	panic("unimplemented")
}

// DeleteMessages implements webhooksv1.WebHookServiceServer.
func (*Handler) DeleteMessages(context.Context, *webhooksv1.DeleteMessagesRequest) (*webhooksv1.DeleteMessagesResponse, error) {
	panic("unimplemented")
}

// GetApps implements webhooksv1.WebHookServiceServer.
func (*Handler) GetApps(context.Context, *webhooksv1.GetAppsRequest) (*webhooksv1.GetAppsResponse, error) {
	panic("unimplemented")
}

// GetEndpoints implements webhooksv1.WebHookServiceServer.
func (*Handler) GetEndpoints(context.Context, *webhooksv1.GetEndpointsRequest) (*webhooksv1.GetEndpointsResponse, error) {
	panic("unimplemented")
}

// GetMessages implements webhooksv1.WebHookServiceServer.
func (*Handler) GetMessages(context.Context, *webhooksv1.GetMessagesRequest) (*webhooksv1.GetMessagesResponse, error) {
	panic("unimplemented")
}

// ListApps implements webhooksv1.WebHookServiceServer.
func (*Handler) ListApps(context.Context, *webhooksv1.ListAppsRequest) (*webhooksv1.ListAppsResponse, error) {
	panic("unimplemented")
}

// ListEndpoints implements webhooksv1.WebHookServiceServer.
func (*Handler) ListEndpoints(context.Context, *webhooksv1.ListEndpointsRequest) (*webhooksv1.ListEndpointsResponse, error) {
	panic("unimplemented")
}

// ListMessages implements webhooksv1.WebHookServiceServer.
func (*Handler) ListMessages(context.Context, *webhooksv1.ListMessagesRequest) (*webhooksv1.ListMessagesResponse, error) {
	panic("unimplemented")
}
