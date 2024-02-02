package handle

import (
	webhooks "woh/webhooks"
	webhooksv1 "woh/webhooks/adapt/grpc/v1"

	"woh/webhooks/adapt/grpc/convert"

	"woh/webhooks/provide/pub/publish"
)

type Handler struct {
	Repo      webhooks.Repository
	Secrets   webhooks.Secrets
	Convert   convert.Converter
	Publisher publish.Publisher
}

var _ webhooksv1.WebHookServiceServer = &Handler{}
