package handle

import (
	"context"
	"fmt"
	webhooksv1 "woh/webhooks/adapt/grpc/v1"

	"woh/webhooks"

	"github.com/samber/lo"
)

// CreateApps implements webhooksv1.WebHookServiceServer.
func (h *Handler) CreateApps(ctx context.Context, request *webhooksv1.CreateAppsRequest) (*webhooksv1.CreateAppsResponse, error) {
	if res, err := h.Repo.CreateApplications(ctx, h.Convert.NewApplications(request.Data)); err != nil {
		return &webhooksv1.CreateAppsResponse{Errors: webhooks.GrpcErrors(err)}, nil
	} else {
		return &webhooksv1.CreateAppsResponse{Data: h.Convert.Applications(res)}, nil
	}
}

// DeleteApps implements webhooksv1.WebHookServiceServer.
func (h *Handler) DeleteApps(ctx context.Context, request *webhooksv1.DeleteAppsRequest) (*webhooksv1.DeleteAppsResponse, error) {
	if err := h.Repo.DeleteApplications(ctx, request.Ids); err != nil {
		return &webhooksv1.DeleteAppsResponse{Errors: webhooks.GrpcErrors(err)}, nil
	}
	return &webhooksv1.DeleteAppsResponse{}, nil
}

// GetApps implements webhooksv1.WebHookServiceServer.
func (h *Handler) GetApps(ctx context.Context, request *webhooksv1.GetAppsRequest) (*webhooksv1.GetAppsResponse, error) {
	if res, err := h.Repo.GetApplications(ctx, request.Ids); err != nil {
		return &webhooksv1.GetAppsResponse{Errors: webhooks.GrpcErrors(err)}, nil
	} else if len(res) == 0 {
		errors := make([]*webhooksv1.Error, 1)
		errors[0] = &webhooksv1.Error{
			Code:   404,
			Index:  "0",
			Reason: fmt.Sprintf("Applications with uids %s not found", request.Ids),
		}
		return &webhooksv1.GetAppsResponse{Errors: errors}, nil
	} else {

		apps := map[string]*webhooksv1.App{}
		lo.ForEach(res, func(app webhooks.ApplicationDetails, _ int) {
			apps[app.Uid] = h.Convert.Application(app.Application)
		})
		return &webhooksv1.GetAppsResponse{Data: apps}, nil
	}
}

// ListApps implements webhooksv1.WebHookServiceServer.
func (h *Handler) ListApps(ctx context.Context, request *webhooksv1.ListAppsRequest) (*webhooksv1.ListAppsResponse, error) {
	if request.Page == nil {
		request.Page = &webhooksv1.PageRequest{}
	}
	if request.Page.Limit == nil {
		request.Page.Limit = lo.ToPtr(int32(20))
	}
	if res, err := h.Repo.ListApplications(ctx, h.Convert.ApplicationQuery(request.Page)); err != nil {
		return &webhooksv1.ListAppsResponse{Errors: webhooks.GrpcErrors(err)}, nil
	} else {
		apps := lo.Map(res, func(app webhooks.Application, _ int) *webhooksv1.App {
			return h.Convert.Application(app)
		})
		return &webhooksv1.ListAppsResponse{Data: apps}, nil
	}
}
