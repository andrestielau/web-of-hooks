package handle

import (
	"context"
	webhooksv1 "woh/webhooks/adapt/grpc/v1"
)

// CreateMessages implements webhooksv1.WebHookServiceServer.
func (h *Handler) CreateMessages(context.Context, *webhooksv1.CreateMessagesRequest) (*webhooksv1.CreateMessagesResponse, error) {
	panic("unimplemented")
}

// DeleteMessages implements webhooksv1.WebHookServiceServer.
func (h *Handler) DeleteMessages(context.Context, *webhooksv1.DeleteMessagesRequest) (*webhooksv1.DeleteMessagesResponse, error) {
	panic("unimplemented")
}

// GetMessages implements webhooksv1.WebHookServiceServer.
func (h *Handler) GetMessages(context.Context, *webhooksv1.GetMessagesRequest) (*webhooksv1.GetMessagesResponse, error) {
	panic("unimplemented")
}

// ListMessages implements webhooksv1.WebHookServiceServer.
func (h *Handler) ListMessages(context.Context, *webhooksv1.ListMessagesRequest) (*webhooksv1.ListMessagesResponse, error) {
	panic("unimplemented")
}
