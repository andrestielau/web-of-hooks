package handle

import (
	"context"
	"fmt"
	"woh/webhooks"
	webhooksv1 "woh/webhooks/adapt/grpc/v1"

	"github.com/samber/lo"
)

// CreateEndpoints implements webhooksv1.WebHookServiceServer.
func (h *Handler) CreateEndpoints(ctx context.Context, request *webhooksv1.CreateEndpointsRequest) (*webhooksv1.CreateEndpointsResponse, error) {
	if res, err := h.Repo.CreateEndpoints(ctx, h.Convert.NewEndpoints(request.Data)); err != nil {
		return &webhooksv1.CreateEndpointsResponse{Errors: webhooks.GrpcErrors(err)}, nil
	} else {
		endpoints := lo.Map(res, func(detail webhooks.EndpointDetails, _ int) webhooks.Endpoint {
			return detail.Endpoint
		})
		return &webhooksv1.CreateEndpointsResponse{Data: h.Convert.Endpoints(endpoints)}, nil
	}
}

// DeleteEndpoints implements webhooksv1.WebHookServiceServer.
func (h *Handler) DeleteEndpoints(ctx context.Context, request *webhooksv1.DeleteEndpointsRequest) (*webhooksv1.DeleteEndpointsResponse, error) {
	if err := h.Repo.DeleteEndpoints(ctx, request.Ids); err != nil {
		return &webhooksv1.DeleteEndpointsResponse{Errors: webhooks.GrpcErrors(err)}, nil
	}
	return &webhooksv1.DeleteEndpointsResponse{}, nil
}

// GetEndpoints implements webhooksv1.WebHookServiceServer.
func (h *Handler) GetEndpoints(ctx context.Context, request *webhooksv1.GetEndpointsRequest) (*webhooksv1.GetEndpointsResponse, error) {
	if res, err := h.Repo.GetEndpoints(ctx, request.Ids); err != nil {
		return &webhooksv1.GetEndpointsResponse{Errors: webhooks.GrpcErrors(err)}, nil
	} else if len(res) == 0 {
		errors := make([]*webhooksv1.Error, 1)
		errors[0] = &webhooksv1.Error{
			Code:   404,
			Index:  "0",
			Reason: fmt.Sprintf("Endpoints with uids %s not found", request.Ids),
		}
		return &webhooksv1.GetEndpointsResponse{Errors: errors}, nil
	} else {

		endpoints := map[string]*webhooksv1.Endpoint{}
		lo.ForEach(res, func(detail webhooks.EndpointDetails, _ int) {
			endpoints[detail.Uid] = h.Convert.Endpoint(detail.Endpoint)
		})
		return &webhooksv1.GetEndpointsResponse{Data: endpoints}, nil
	}
}

// ListEndpoints implements webhooksv1.WebHookServiceServer.
func (h *Handler) ListEndpoints(ctx context.Context, request *webhooksv1.ListEndpointsRequest) (*webhooksv1.ListEndpointsResponse, error) {
	if request.Page == nil {
		request.Page = &webhooksv1.PageRequest{}
	}
	if request.Page.Limit == nil {
		request.Page.Limit = lo.ToPtr(int32(20))
	}
	if res, err := h.Repo.ListEndpoints(ctx, h.Convert.EndpointQuery(request.Page)); err != nil {
		return &webhooksv1.ListEndpointsResponse{Errors: webhooks.GrpcErrors(err)}, nil
	} else {
		endpoints := lo.Map(res, func(detail webhooks.Endpoint, _ int) *webhooksv1.Endpoint {
			return h.Convert.Endpoint(detail)
		})
		return &webhooksv1.ListEndpointsResponse{Data: endpoints}, nil
	}
}
