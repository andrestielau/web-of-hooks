
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
