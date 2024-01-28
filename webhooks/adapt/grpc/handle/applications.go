package handle

import (
	"context"
	webhooksv1 "woh/webhooks/adapt/grpc/v1"
)

// CreateApps implements webhooksv1.WebHookServiceServer.
func (h *Handler) CreateApps(context.Context, *webhooksv1.CreateAppsRequest) (*webhooksv1.CreateAppsResponse, error) {
	panic("unimplemented")
}

// DeleteApps implements webhooksv1.WebHookServiceServer.
func (h *Handler) DeleteApps(context.Context, *webhooksv1.DeleteAppsRequest) (*webhooksv1.DeleteAppsResponse, error) {
	panic("unimplemented")
}

// GetApps implements webhooksv1.WebHookServiceServer.
func (h *Handler) GetApps(context.Context, *webhooksv1.GetAppsRequest) (*webhooksv1.GetAppsResponse, error) {
	panic("unimplemented")
}

// ListApps implements webhooksv1.WebHookServiceServer.
func (h *Handler) ListApps(context.Context, *webhooksv1.ListAppsRequest) (*webhooksv1.ListAppsResponse, error) {
	panic("unimplemented")
}
