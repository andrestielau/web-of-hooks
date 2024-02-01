package dispatcher_test

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"woh/package/actor/sql/pgx"

	"woh/webhooks"
	"woh/webhooks/adapt/subs/dispatcher"
	"woh/webhooks/provide/repo"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func TestDispatcher(t *testing.T) {
	timer := time.NewTimer(time.Second)
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		log.Println("yo", r.Header)
		timer.Stop()
		json.NewEncoder(w).Encode(map[string]any{"status": "all right"})
	}))
	r := repo.New(pgx.ProvideOptions())
	ctx := context.Background()
	defer r.Stop(ctx)
	r.Start(ctx)

	appRes, err := r.CreateApplications(ctx, []webhooks.NewApplication{{
		Name:      uuid.NewString(),
		RateLimit: lo.ToPtr[int32](10),
	}})
	require.NoError(t, err)
	require.Len(t, appRes, 1)
	log.Println(appRes)

	secretRes, err := r.CreateSecrets(ctx, []webhooks.NewSecret{{
		ApplicationID: appRes[0].Uid,
		Value:         base64.StdEncoding.EncodeToString([]byte(uuid.NewString())),
	}})
	require.NoError(t, err)
	require.Len(t, secretRes, 1)
	log.Println(secretRes)
	epRes, err := r.CreateEndpoints(ctx, []webhooks.NewEndpoint{{
		RateLimit:     lo.ToPtr[int32](10),
		SecretId:      secretRes[0].Uid,
		ApplicationID: appRes[0].Uid,
		Url:           s.URL,
	}})
	require.NoError(t, err)
	require.Len(t, epRes, 1)
	log.Println(epRes)

	h := dispatcher.Handler{Repo: r}
	res, err := h.Handle(&message.Message{
		UUID: uuid.NewString(),
		Metadata: message.Metadata{
			"EndpointID": epRes[0].Uid,
		},
		Payload: message.Payload(`{ "foo": "bar" }`),
	})
	require.NoError(t, err)
	log.Println(res)
	<-timer.C
}
