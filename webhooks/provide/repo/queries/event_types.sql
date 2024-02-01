
-- CreateEventTypes inserts event types into the database
-- name: CreateEventTypes :many
 INSERT INTO webhooks.event_type (
    key
) 
SELECT 
    u.key
FROM unnest(pggen.arg('event_types')::webhooks.new_event_type[]) u
ON CONFLICT DO NOTHING
RETURNING (
    id,
    key,
    uid,
    created_at
):: webhooks.event_type;

-- DeleteEventTypes deletes endpoints by uid
-- name: DeleteEventTypes :exec
DELETE FROM webhooks.event_type WHERE key = ANY(pggen.arg('keys')::TEXT[]);

-- GetEventTypes gets event-types by id
-- name: GetEventTypes :many
SELECT (
    id,
    key,
    uid,
    created_at
):: webhooks.event_type
FROM webhooks.event_type
WHERE uid = ANY(pggen.arg('ids')::uuid[]);

-- GetEventTypes gets event-types by id
-- name: GetEventTypesByKeys :many
SELECT (
    id,
    key,
    uid,
    created_at
):: webhooks.event_type
FROM webhooks.event_type
WHERE key = ANY(pggen.arg('keys'));

-- ListEventTypes lists event-types
-- name: ListEventTypes :many
SELECT (
    id,
    key,
    uid,
    created_at
):: webhooks.event_type
FROM webhooks.event_type
WHERE created_at > pggen.arg('created_after') 
ORDER BY uid
LIMIT pggen.arg('limit')
OFFSET pggen.arg('offset');
