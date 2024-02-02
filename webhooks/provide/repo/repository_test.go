package repo_test

import (
	"context"
	"encoding/base64"
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

	eventTypeIds := make([]*int32, numEventTypes)
	eventTypeKeys := make([]string, numEventTypes)
	for _, eventType := range res {
		require.NotEmpty(t, eventType.Key)
		expect, i, ok := lo.FindIndexOf(newEventTypes, func(e webhooks.NewEventType) bool {
			return e.Key == eventType.Key
		})
		require.True(t, ok)
		require.Nil(t, eventTypeIds[i])
		require.Equal(t, eventType.Key, expect.Key)
		eventTypeKeys[i] = eventType.Key
		eventTypeIds[i] = &eventType.ID
	}

	numTenants := 2 // 20 ants (icebreaker)
	tenantIds := lo.Times(numTenants, func(i int) string {
		return fmt.Sprintf("tenant%d-%s", i, uuid.NewString())
	})

	appsPerTenant := 2
	newApplications := lo.Times(numTenants*appsPerTenant, func(i int) webhooks.NewApplication {
		return webhooks.NewApplication{
			TenantID:  tenantIds[i%numTenants],
			Name:      fmt.Sprintf("app%d-%s", i+1, uuid.NewString()),
			RateLimit: lo.ToPtr[int32](rand.Int31() % 20),
		}
	})
	res2, err := r.CreateApplications(ctx, newApplications)
	require.NoError(t, err)
	require.Len(t, res2, len(newApplications))

	appIds := make([]*int32, len(newApplications))
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
		appIds[i] = app.ID
		appUids[i] = app.Uid
	}

	//CREATE SECRETS
	newSecrets := lo.Times(numTenants*appsPerTenant, func(i int) webhooks.NewSecret {
		return webhooks.NewSecret{
			Value:         uuid.NewString(),
			ApplicationID: appUids[i/(appsPerTenant*numTenants)],
		}
	})

	resSecrets, err := r.CreateSecrets(ctx, newSecrets)
	require.NoError(t, err)
	require.Len(t, resSecrets, len(newSecrets))

	secretIds := make([]*int32, numTenants*appsPerTenant)
	secretUids := make([]string, numTenants*appsPerTenant)
	for _, secret := range resSecrets {
		require.NotEmpty(t, secret.Value)
		expect, i, ok := lo.FindIndexOf(newSecrets, func(e webhooks.NewSecret) bool {
			return e.Value == secret.Value
		})
		require.True(t, ok)
		require.Nil(t, secretIds[i])
		require.Equal(t, secret.Value, expect.Value)
		secretUids[i] = secret.Uid
		secretIds[i] = secret.ID
	}

	endpointsPerApp := 3
	newEndpoints := lo.Times(endpointsPerApp*len(appIds), func(i int) webhooks.NewEndpoint {
		var filterTypes []string
		switch i % 4 {
		case 1:
			filterTypes = []string{eventTypeKeys[0]}
		case 2:
			filterTypes = []string{eventTypeKeys[1], eventTypeKeys[2]}
		case 3:
			filterTypes = []string{eventTypeKeys[0], eventTypeKeys[2]}
		}
		return webhooks.NewEndpoint{
			Url:           fmt.Sprintf("http://app%s.com/endpoint%d", appUids[i/endpointsPerApp], i+1),
			Name:          fmt.Sprintf("Endpoint%d-%s", i, uuid.NewString()),
			RateLimit:     lo.ToPtr[int32](rand.Int31() % 20),
			Description:   "description " + strconv.Itoa(i),
			ApplicationID: appUids[i/endpointsPerApp],
			FilterTypeIds: filterTypes,
			SecretId:      secretUids[i/(numTenants*appsPerTenant)],
		}
	})

	res3, err := r.CreateEndpoints(ctx, newEndpoints)
	require.NoError(t, err)
	require.Len(t, res3, len(newEndpoints))
	endpointIds := make([]*int32, len(newEndpoints))
	endpointUids := make([]string, len(newEndpoints))
	for _, endpoint := range res3 {
		require.NotNil(t, endpoint.ID)
		require.NotEmpty(t, endpoint.Uid)
		require.Nil(t, endpoint.Disabled)
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
		endpointIds[i] = endpoint.ID
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
	require.Len(t, res4, len(newMessages))
	messageIds := make([]*int32, len(newMessages))
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
		require.NotNil(t, msg.ID)
		require.NotEmpty(t, msg.Uid)
		require.Equal(t, expect.Payload, msg.Payload)
		messageIds[i] = msg.ID
		messageUids[i] = msg.Uid
		// TODO:Check Attempts
	}
}

func TestPopulate(t *testing.T) {
	a := repo.New(pgx.Options{
		URL: "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable",
	})
	ctx := context.Background()
	a.Start(ctx)
	defer a.Stop(ctx)

	r := a.Repository
	// CREATE EVENTS TYPE

	numEventTypes := 3
	newEventTypes := lo.Times(numEventTypes, func(i int) webhooks.NewEventType {
		return webhooks.NewEventType{
			Key: uuid.NewString(), // TODO: key shouldn't be uuid
		}
	})

	res, err := r.CreateEventTypes(ctx, newEventTypes)
	require.NoError(t, err)
	require.Len(t, res, len(newEventTypes))
	numEventTypes = len(newEventTypes)

	eventTypeIds := make([]*int32, numEventTypes)
	eventTypeUids := make([]string, numEventTypes)
	for _, eventType := range res {
		require.NotEmpty(t, eventType.Key)
		expect, i, ok := lo.FindIndexOf(newEventTypes, func(e webhooks.NewEventType) bool {
			return e.Key == eventType.Key
		})
		require.True(t, ok)
		require.Nil(t, eventTypeIds[i])
		require.Equal(t, eventType.Key, expect.Key)
		eventTypeUids[i] = eventType.Uid
		eventTypeIds[i] = &eventType.ID
	}

	_, err = r.CreateEventTypes(ctx, []webhooks.NewEventType{{Key: "onboarding.approved"}, {Key: "depositattribution.pending"}})
	require.NoError(t, err)

	numTenants := 2 // 20 ants (icebreaker)
	tenantIds := lo.Times(numTenants, func(i int) string {
		return fmt.Sprintf("tenant%d-%s", i, uuid.NewString())
	})

	//CREATE APPLICATIONS
	appsPerTenant := 2
	newApplications := lo.Times(numTenants*appsPerTenant, func(i int) webhooks.NewApplication {
		return webhooks.NewApplication{
			TenantID:  tenantIds[i%numTenants],
			Name:      fmt.Sprintf("app%d-%s", i+1, uuid.NewString()),
			RateLimit: lo.ToPtr[int32](rand.Int31() % 20),
		}
	})
	res2, err := r.CreateApplications(ctx, newApplications)
	require.NoError(t, err)
	require.Len(t, res2, len(newApplications))

	appIds := make([]*int32, len(newApplications))
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
		appIds[i] = app.ID
		appUids[i] = app.Uid
	}

	//CREATE SECRETS
	newSecrets := lo.Times(numTenants*appsPerTenant, func(i int) webhooks.NewSecret {
		return webhooks.NewSecret{
			Value:         uuid.NewString(),
			ApplicationID: appUids[i/(appsPerTenant*numTenants)],
		}
	})

	resSecrets, err := r.CreateSecrets(ctx, newSecrets)
	require.NoError(t, err)
	require.Len(t, resSecrets, len(newSecrets))

	secretIds := make([]*int32, numTenants*appsPerTenant)
	secretUids := make([]string, numTenants*appsPerTenant)
	for _, secret := range resSecrets {
		require.NotEmpty(t, secret.Value)
		expect, i, ok := lo.FindIndexOf(newSecrets, func(e webhooks.NewSecret) bool {
			return e.Value == secret.Value
		})
		require.True(t, ok)
		require.Nil(t, secretIds[i])
		require.Equal(t, secret.Value, expect.Value)
		secretUids[i] = secret.Uid
		secretIds[i] = secret.ID
	}

	//CREATE ENDPOINTS
	endpointsPerApp := 3
	newEndpoints := lo.Times(endpointsPerApp*len(appIds), func(i int) webhooks.NewEndpoint {
		var filterTypes []string
		switch i % 4 {
		case 1:
			filterTypes = []string{eventTypeUids[0]}
		case 2:
			filterTypes = []string{eventTypeUids[1], eventTypeUids[2]}
		case 3:
			filterTypes = []string{eventTypeUids[0], eventTypeUids[2]}
		}
		return webhooks.NewEndpoint{
			Url:           fmt.Sprintf("http://app%s.com/endpoint%d", appUids[i/endpointsPerApp], i+1),
			Name:          fmt.Sprintf("Endpoint%d-%s", i, uuid.NewString()),
			RateLimit:     lo.ToPtr[int32](rand.Int31() % 20),
			Description:   "description " + strconv.Itoa(i),
			ApplicationID: appUids[i/endpointsPerApp],
			FilterTypeIds: filterTypes,
			SecretId:      secretUids[i/(numTenants*appsPerTenant)],
		}
	})

	res3, err := r.CreateEndpoints(ctx, newEndpoints)
	require.NoError(t, err)
	require.Len(t, res3, len(newEndpoints))
	endpointIds := make([]*int32, len(newEndpoints))
	endpointUids := make([]string, len(newEndpoints))
	for _, endpoint := range res3 {
		require.NotNil(t, endpoint.ID)
		require.NotEmpty(t, endpoint.Uid)
		require.Nil(t, endpoint.Disabled)
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
		endpointIds[i] = endpoint.ID
		endpointUids[i] = endpoint.Uid
	}

	//CREATE MESSAGES
	newMessages := lo.Times(numTenants*appsPerTenant*numEventTypes, func(i int) webhooks.NewMessage {
		return webhooks.NewMessage{
			EventTypeID:   eventTypeUids[i%len(newEventTypes)],
			ApplicationID: appUids[i/(appsPerTenant*numTenants)],
			EventID:       fmt.Sprintf("msg%d-%s", i, uuid.NewString()),
			Payload:       "something-" + uuid.NewString(),
		}
	})
	res4, err := r.CreateMessages(ctx, newMessages)
	require.NoError(t, err)
	require.Len(t, res4, len(newMessages))
	messageIds := make([]*int32, len(newMessages))
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
		require.NotNil(t, msg.ID)
		require.NotEmpty(t, msg.Uid)
		require.Equal(t, expect.Payload, msg.Payload)
		messageIds[i] = msg.ID
		messageUids[i] = msg.Uid
		// TODO:Check Attempts
	}
}

func TestPopulateEventTypes(t *testing.T) {
	a := repo.New(pgx.Options{
		URL: "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable",
	})
	ctx := context.Background()
	a.Start(ctx)
	defer a.Stop(ctx)

	r := a.Repository
	// CREATE EVENTS TYPE

	numEventTypes := 1
	newEventTypes := lo.Times(numEventTypes, func(i int) webhooks.NewEventType {
		return webhooks.NewEventType{
			Key: uuid.NewString(), // TODO: key shouldn't be uuid
		}
	})

	res, err := r.CreateEventTypes(ctx, newEventTypes)
	require.NoError(t, err)
	require.Len(t, res, len(newEventTypes))
	numEventTypes = len(newEventTypes)

	eventTypeIds := make([]*int32, numEventTypes)
	eventTypeUids := make([]string, numEventTypes)
	for _, eventType := range res {
		require.NotEmpty(t, eventType.Key)
		expect, i, ok := lo.FindIndexOf(newEventTypes, func(e webhooks.NewEventType) bool {
			return e.Key == eventType.Key
		})
		require.True(t, ok)
		require.Nil(t, eventTypeIds[i])
		require.Equal(t, eventType.Key, expect.Key)
		eventTypeUids[i] = eventType.Uid
		eventTypeIds[i] = &eventType.ID
	}

	_, err = r.CreateEventTypes(ctx, []webhooks.NewEventType{{Key: "onboarding.approved"}, {Key: "depositattribution.pending"}})
	require.NoError(t, err)

}

func TestPopulateOne(t *testing.T) {
	a := repo.New(pgx.Options{
		URL: "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable",
	})
	ctx := context.Background()
	a.Start(ctx)
	defer a.Stop(ctx)

	r := a.Repository
	// VARIABLES
	tenantID := "tenant1"
	applicationName := "app1"
	endpointUrl := "https://webhook.site/ef9bc06a-b1e6-435f-a438-4f1093b47c26"
	endpointName := "endpoint1"
	eventType1 := "onboarding.approved"
	eventType2 := "depositattribution.pending"

	// CREATE EVENTS TYPE
	eventTypes, err := r.CreateEventTypes(ctx, []webhooks.NewEventType{{Key: eventType1}, {Key: eventType2}})
	require.NoError(t, err)
	if len(eventTypes) == 0 {
		eventTypes, err = r.GetEventTypesByKeys(ctx, []string{eventType1, eventType2})
		require.NoError(t, err)
	}
	require.Len(t, eventTypes, len(eventTypes))

	//CREATE APPLICATIONS
	newApplications := []webhooks.NewApplication{{
		TenantID:  tenantID,
		Name:      applicationName,
		RateLimit: lo.ToPtr[int32](rand.Int31() % 20),
	}}

	apps, err := r.CreateApplications(ctx, newApplications)
	require.NoError(t, err)
	if len(apps) == 0 {
		apps, err = r.GetApplicationsByName(ctx, []string{applicationName})
		require.NoError(t, err)
	}
	require.Len(t, apps, len(newApplications))

	//CREATE SECRETS
	newSecrets := []webhooks.NewSecret{{
		Value:         base64.StdEncoding.EncodeToString([]byte(uuid.NewString())),
		ApplicationID: apps[0].Uid,
	}}

	secrets, err := r.CreateSecrets(ctx, newSecrets)
	require.NoError(t, err)
	require.Len(t, secrets, len(newSecrets))

	//CREATE ENDPOINTS
	newEndpoints := []webhooks.NewEndpoint{{
		Url:           endpointUrl,
		Name:          endpointName,
		RateLimit:     lo.ToPtr[int32](rand.Int31() % 20),
		Description:   "description " + endpointName,
		ApplicationID: apps[0].Uid,
		FilterTypeIds: []string{eventTypes[0].Uid},
		SecretId:      secrets[0].Uid,
	}}

	endpoints, err := r.CreateEndpoints(ctx, newEndpoints)
	require.NoError(t, err)
	if len(endpoints) == 0 {
		endpoints, err = r.GetEndpointsByUrl(ctx, []string{endpointUrl})
		require.NoError(t, err)
	}
	require.Len(t, endpoints, len(newEndpoints))
}
