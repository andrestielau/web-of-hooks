package handle

import (
	webhooks "woh/webhooks"
	webhooksv1 "woh/webhooks/adapt/grpc/v1"
)

type Handler struct {
	Repo    webhooks.Repository
	Secrets webhooks.Secrets
}

var _ webhooksv1.WebHookServiceServer = &Handler{}
