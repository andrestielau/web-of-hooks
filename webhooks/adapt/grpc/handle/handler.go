package handle

import (
	webhooks "woh/webhooks"
	webhooksv1 "woh/webhooks/adapt/grpc/v1"

	"woh/webhooks/adapt/grpc/convert"
)

type Handler struct {
	Repo    webhooks.Repository
	Secrets webhooks.Secrets
	Convert convert.Converter
}

var _ webhooksv1.WebHookServiceServer = &Handler{}
