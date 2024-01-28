
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
FROM unnest(pggen.arg('event_types')::webhooks.new_application[]) u
ON CONFLICT DO NOTHING
RETURNING 
    id,
    uid,
    tenant_id,
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