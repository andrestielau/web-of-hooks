package cron

import (
	"context"
	webhooks "woh/webhooks"
)

type Handler struct {
	Id   WorkerId
	Repo webhooks.Repository
}

func (h *Handler) Work(context.Context) {
	// Dequeue for partition and update last seen
	// for each app in partition
	// start if not started
	// for each started and not in partition

}
