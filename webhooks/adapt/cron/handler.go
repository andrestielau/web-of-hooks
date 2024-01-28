package cron

import (
	"context"
	"log"
	webhooks "woh/webhooks"
)

type Handler struct {
	Repo    webhooks.Repository
	Secrets webhooks.Secrets
}

func (h *Handler) Send(context.Context) {
	log.Println("send")
}

func (h *Handler) Sign(context.Context) {
	log.Println("sign")
}
