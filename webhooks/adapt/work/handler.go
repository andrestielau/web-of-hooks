package work

import "woh/webhooks"

type Handler struct {
	Repo    webhooks.Repository
	Secrets webhooks.Secrets
}
