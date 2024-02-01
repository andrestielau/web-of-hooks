package handle

import (
	"context"
	"fmt"
	webhooksv1 "woh/webhooks/adapt/grpc/v1"

	"woh/webhooks"

	"github.com/samber/lo"
)

// CreateMessages implements webhooksv1.WebHookServiceServer.
func (h *Handler) CreateMessages(ctx context.Context, request *webhooksv1.CreateMessagesRequest) (*webhooksv1.CreateMessagesResponse, error) {
	if res, err := h.Repo.CreateMessages(ctx, h.Convert.NewMessages(request.Data)); err != nil {
		return &webhooksv1.CreateMessagesResponse{Errors: webhooks.GrpcErrors(err)}, nil
	} else {
		messages := lo.Map(res, func(detail webhooks.MessageDetails, _ int) webhooks.Message {
			return detail.Message
		})
		return &webhooksv1.CreateMessagesResponse{Data: h.Convert.Messages(messages)}, nil
	}
}

// DeleteMessages implements webhooksv1.WebHookServiceServer.
func (h *Handler) DeleteMessages(ctx context.Context, request *webhooksv1.DeleteMessagesRequest) (*webhooksv1.DeleteMessagesResponse, error) {
	if err := h.Repo.DeleteMessages(ctx, request.Ids); err != nil {
		return &webhooksv1.DeleteMessagesResponse{Errors: webhooks.GrpcErrors(err)}, nil
	}
	return &webhooksv1.DeleteMessagesResponse{}, nil
}

// GetMessages implements webhooksv1.WebHookServiceServer.
func (h *Handler) GetMessages(ctx context.Context, request *webhooksv1.GetMessagesRequest) (*webhooksv1.GetMessagesResponse, error) {
	if res, err := h.Repo.GetMessages(ctx, request.Ids); err != nil {
		return &webhooksv1.GetMessagesResponse{Errors: webhooks.GrpcErrors(err)}, nil
	} else if len(res) == 0 {
		errors := make([]*webhooksv1.Error, 1)
		errors[0] = &webhooksv1.Error{
			Code:   404,
			Index:  "0",
			Reason: fmt.Sprintf("Messages with uids %s not found", request.Ids),
		}
		return &webhooksv1.GetMessagesResponse{Errors: errors}, nil
	} else {

		messages := map[string]*webhooksv1.Message{}
		lo.ForEach(res, func(detail webhooks.MessageDetails, _ int) {
			messages[detail.Uid] = h.Convert.Message(detail.Message)
		})
		return &webhooksv1.GetMessagesResponse{Data: messages}, nil
	}
}

// ListMessages implements webhooksv1.WebHookServiceServer.
func (h *Handler) ListMessages(ctx context.Context, request *webhooksv1.ListMessagesRequest) (*webhooksv1.ListMessagesResponse, error) {
	if request.Page == nil {
		request.Page = &webhooksv1.PageRequest{}
	}
	if request.Page.Limit == nil {
		request.Page.Limit = lo.ToPtr(int32(20))
	}
	if res, err := h.Repo.ListMessages(ctx, h.Convert.MessageQuery(request.Page)); err != nil {
		return &webhooksv1.ListMessagesResponse{Errors: webhooks.GrpcErrors(err)}, nil
	} else {
		messages := lo.Map(res, func(detail webhooks.Message, _ int) *webhooksv1.Message {
			return h.Convert.Message(detail)
		})
		return &webhooksv1.ListMessagesResponse{Data: messages}, nil
	}
}

// EmitEvent implements webhooksv1.WebHookServiceServer.
func (h *Handler) EmitEvent(ctx context.Context, request *webhooksv1.EmitEventRequest) (*webhooksv1.EmitEventResponse, error) {
	response := webhooksv1.EmitEventResponse{}
	if res, err := h.Repo.EmitEvent(ctx, h.Convert.NewEvent(request)); err != nil {
		response.Errors = webhooks.GrpcErrors(err)
		return &response, nil
	} else {
		response.Data = h.Convert.Messages(res)
	}
	return &response, nil

}
