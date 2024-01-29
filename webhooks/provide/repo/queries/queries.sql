
-- CreateEventTypes inserts event types into the database
-- name: CreateEventTypes :many
 INSERT INTO webhooks.event_type (
    key
) 
SELECT 
    u.key
FROM unnest(pggen.arg('event_types')::webhooks.new_event_type[]) u
ON CONFLICT DO NOTHING
RETURNING 
    id,
    uid,
    key,
    created_at;

-- CreateApplications inserts applications into the database
-- name: CreateApplications :many
 INSERT INTO webhooks.application (
    tenant_id,
    rate_limit,
    metadata
) 
SELECT 
    u.tenant_id,
    u.rate_limit,
    u.metadata
FROM unnest(pggen.arg('applications')::webhooks.new_application[]) u
ON CONFLICT DO NOTHING
RETURNING 
    id,
    uid,
    tenant_id,
    rate_limit,
    metadata,
    created_at;

-- CreateEndpoints inserts endpoints into the database
-- name: CreateEndpoints :many
 INSERT INTO webhooks.endpoint (
    application_id,
    url,
    name,
    rate_limit,
    metadata,
    description
) 
SELECT 
    a.id,
    u.url,
    u.name,
    u.rate_limit,
    u.metadata,
    u.description
FROM unnest(pggen.arg('endpoints')::webhooks.new_endpoint[]) u
JOIN webhooks.application a ON u.application_id = a.uid
ON CONFLICT DO NOTHING
RETURNING 
    id,
    uid,
        application_id,
    rate_limit,
    metadata,
    created_at;

-- ListEventTypes lists event-types
-- name: ListEventTypes :many
SELECT
    id,
    uid,
    key,
    created_at
FROM webhooks.event_type
ORDER BY uid
LIMIT pggen.arg('limit')
OFFSET pggen.arg('offset');

-- ListEndpoints lists endpoints
-- name: ListEndpoints :many
SELECT 
    id,
    uid,
    url,
    name,
    metadata,
    disabled,
    rate_limit,
    created_at,
    updated_at,
    description,
    application_id
FROM webhooks.endpoint
ORDER BY uid
LIMIT pggen.arg('limit')
OFFSET pggen.arg('offset');
