package handle

import (
	"net/http"
	"strings"
	"time"

	"woh/package/utils/media"
	"woh/webhooks"
	webhooksv1 "woh/webhooks/adapt/http/v1"
	"woh/webhooks/render/page"
	"woh/webhooks/render/style/theme"

	"woh/webhooks/adapt/http/convert"

	"github.com/alexedwards/scs/v2"
)

type Handler struct {
	Session *scs.SessionManager
	Repo    webhooks.Repository
	Secrets webhooks.Secrets
	Convert Converters
}
type Converters struct {
	Application *convert.ApplicationConverterImpl
	Endpoint    *convert.EndpointConverterImpl
	Secret      *convert.SecretConverterImpl
	Message     *convert.MessageConverterImpl
}

func Convert() Converters {
	return Converters{
		Application: &convert.ApplicationConverterImpl{},
		Endpoint:    &convert.EndpointConverterImpl{},
		Secret:      &convert.SecretConverterImpl{},
		Message:     &convert.MessageConverterImpl{},
	}
}

var _ webhooksv1.ServerInterface = &Handler{}

const status = "ok"

var start = time.Now()

// GetHealth implements webhooksv1.ServerInterface.
func (h *Handler) GetHealth(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.Header.Get("Accept"), "text/html") {
		currentCount := h.Session.GetInt(r.Context(), "count")
		h.Session.Put(r.Context(), "count", currentCount+1)
		t := "dark"
		if currentCount%2 == 1 {
			t = "light"
		}
		page.Health(status, start.String()).Render(theme.Set(r.Context(), t), w)
		return
	}
	media.Res(w, media.Accept(r), map[string]any{"status": status, "since": start})
}
