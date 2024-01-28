package handle

import (
	"context"
	webhooksv1 "woh/webhooks/adapt/grpc/v1"
)

// CreateEndpoints implements webhooksv1.WebHookServiceServer.
func (h *Handler) CreateEndpoints(context.Context, *webhooksv1.CreateEndpointsRequest) (*webhooksv1.CreateEndpointsResponse, error) {
	panic("unimplemented")
}

// DeleteEndpoints implements webhooksv1.WebHookServiceServer.
func (h *Handler) DeleteEndpoints(context.Context, *webhooksv1.DeleteEndpointsRequest) (*webhooksv1.DeleteEndpointsResponse, error) {
	panic("unimplemented")
}

// GetEndpoints implements webhooksv1.WebHookServiceServer.
func (h *Handler) GetEndpoints(context.Context, *webhooksv1.GetEndpointsRequest) (*webhooksv1.GetEndpointsResponse, error) {
	panic("unimplemented")
}

// ListEndpoints implements webhooksv1.WebHookServiceServer.
func (h *Handler) ListEndpoints(context.Context, *webhooksv1.ListEndpointsRequest) (*webhooksv1.ListEndpointsResponse, error) {
	panic("unimplemented")
}