package queries_test

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"testing"

	"woh/package/actor/sql/pgx"
	"woh/webhooks"
	"woh/webhooks/provide/repo"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func TestDequeue(t *testing.T) {
	a := repo.New(pgx.Options{
		URL: "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable",
	})
	ctx := context.Background()
	a.Start(ctx)
	defer a.Stop(ctx)

	r := a.Repository
	numEventTypes := 3
	newEventTypes := lo.Times(numEventTypes, func(i int) webhooks.NewEventType {
		return webhooks.NewEventType{
			Key: uuid.NewString(), // TODO: key shouldn't be uuid
		}
	})
	res, err := r.CreateEventTypes(ctx, newEventTypes)
	require.NoError(t, err)
	require.Len(t, res, len(newEventTypes))

	eventTypeKeys := make([]string, numEventTypes)
	for _, eventType := range res {
		require.NotEmpty(t, eventType.Key)
		expect, i, ok := lo.FindIndexOf(newEventTypes, func(e webhooks.NewEventType) bool {
			return e.Key == eventType.Key
		})
		require.True(t, ok)
		require.Nil(t, eventTypeKeys[i])
		require.Equal(t, eventType.Key, expect.Key)
		eventTypeKeys[i] = eventType.Key
	}

	numTenants := 2 // 20 ants (icebreaker)
	tenantIds := lo.Times(numTenants, func(i int) string {
		return fmt.Sprintf("tenant%d-%s", i, uuid.NewString())
	})

	appsPerTenant := 2
	newApplications := lo.Times(numTenants*appsPerTenant, func(i int) webhooks.NewApplication {
		return webhooks.NewApplication{
			TenantID:  tenantIds[i%numTenants],
			Name:      fmt.Sprintf("app%d-%s", i, uuid.NewString()),
			RateLimit: lo.ToPtr[int32](min(rand.Int31(), 20)),
		}
	})
	res2, err := r.CreateApplications(ctx, newApplications)
	require.NoError(t, err)
	require.Len(t, res2, len(newApplications))

	appIds := make([]int32, len(newApplications))
	appUids := make([]string, len(newApplications))
	for _, app := range res2 { // TODO: check metadata
		require.NotNil(t, app.ID)
		require.NotEmpty(t, app.Uid)
		expect, i, ok := lo.FindIndexOf(newApplications, func(a webhooks.NewApplication) bool {
			return a.Name == app.Name
		})
		require.True(t, ok)
		require.Nil(t, appIds[i])
		require.Equal(t, expect.TenantID, app.TenantID)
		require.Equal(t, expect.RateLimit, app.RateLimit)
		appIds[i] = *app.ID
		appUids[i] = app.Uid
	}

	endpointsPerApp := 3
	newEndpoints := lo.Times(endpointsPerApp*len(appIds), func(i int) webhooks.NewEndpoint {
		var filterTypes []string
		switch i % 4 {
		case 1:
			filterTypes = []string{"e1"}
		case 2:
			filterTypes = []string{"e2", "e3"}
		case 3:
			filterTypes = []string{"e1", "e3"}
		}
		return webhooks.NewEndpoint{
			Url:           fmt.Sprintf("http://app%d.com/endpoint%d", i/2+1, i+1),
			Name:          fmt.Sprintf("Endpoint%d-%s", i, uuid.NewString()),
			RateLimit:     lo.ToPtr[int32](min(rand.Int31(), 10)),
			Description:   "description " + strconv.Itoa(i),
			ApplicationID: appUids[i/endpointsPerApp],
			FilterTypes:   filterTypes,
		}
	})
	res3, err := r.CreateEndpoints(ctx, newEndpoints)
	require.NoError(t, err)
	require.Len(t, res2, len(newEndpoints))
	endpointIds := make([]int32, len(newEndpoints))
	endpointUids := make([]string, len(newEndpoints))
	for _, endpoint := range res3 {
		require.NotNil(t, endpoint.ID)
		require.NotEmpty(t, endpoint.Uid)
		require.NotNil(t, endpoint.Disabled)
		require.False(t, *endpoint.Disabled)
		expect, i, ok := lo.FindIndexOf(newEndpoints, func(e webhooks.NewEndpoint) bool {
			return e.Name == endpoint.Name
		})
		require.True(t, ok)
		require.Nil(t, endpointIds[i])
		_, appIdIdx, ok := lo.FindIndexOf(appUids, func(id string) bool {
			return id == expect.ApplicationID
		})
		require.True(t, ok)
		require.Equal(t, appIds[appIdIdx], endpoint.ApplicationID)
		require.Equal(t, expect.RateLimit, endpoint.RateLimit)
		require.Equal(t, expect.Description, endpoint.Description)
		require.Equal(t, expect.Url, endpoint.Url)
		endpointIds[i] = *endpoint.ID
		endpointUids[i] = endpoint.Uid
	}

	newMessages := lo.Times(numTenants*appsPerTenant*numEventTypes, func(i int) webhooks.NewMessage {
		return webhooks.NewMessage{
			EventTypeID:   eventTypeKeys[i%len(newEventTypes)],
			ApplicationID: appUids[i/(appsPerTenant*numTenants)],
			EventID:       fmt.Sprintf("msg%d-%s", i, uuid.NewString()),
			Payload:       "something-" + uuid.NewString(),
		}
	})
	res4, err := r.CreateMessages(ctx, newMessages)
	require.NoError(t, err)
	require.Len(t, res2, len(newMessages))
	messageIds := make([]int32, len(newMessages))
	messageUids := make([]string, len(newMessages))
	for _, msg := range res4 {
		require.NotNil(t, msg.ID)
		require.NotEmpty(t, msg.Uid)
		require.NotNil(t, msg.EventID)
		expect, i, ok := lo.FindIndexOf(newMessages, func(e webhooks.NewMessage) bool {
			return e.EventID == msg.EventID
		})
		require.True(t, ok)
		require.Nil(t, messageIds[i])
		require.Empty(t, messageUids[i])
		require.Equal(t, expect.Payload, msg.Payload)
		// TODO:Check Attempts
	}
}
