
-- CreateMessages inserts messages into the database
-- name: CreateMessages :many
 INSERT INTO webhooks.message (
    application_id,
    event_type_id,
    event_id,
    payload
) 
SELECT 
    a.id,
    e.id,
    u.event_id,
    u.payload
FROM unnest(pggen.arg('messages')::webhooks.new_message[]) u
JOIN webhooks.application a ON u.application_id = a.uid
JOIN webhooks.event_type e ON u.event_type_id = e.uid
ON CONFLICT DO NOTHING
RETURNING (
    id ,
    application_id ,
    event_type_id ,
    uid ,
    created_at,
    event_id,
    payload
)::webhooks.message;

-- DeleteMessages deletes messages by uid
-- name: DeleteMessages :exec
DELETE FROM webhooks.message WHERE uid = ANY(pggen.arg('ids')::UUID[]);

-- GetMessages gets messages by id
-- name: GetMessages :many
SELECT (
    id ,
    application_id ,
    event_type_id ,
    uid ,
    created_at,
    event_id,
    payload
)::webhooks.message
FROM webhooks.message
WHERE uid = ANY(pggen.arg('ids')::uuid[]);

-- ListMessages lists event-types
-- name: ListMessages :many
SELECT (
    id ,
    application_id ,
    event_type_id ,
    uid ,
    created_at,
    event_id,
    payload
)::webhooks.message
FROM webhooks.message
ORDER BY uid
LIMIT pggen.arg('limit')
OFFSET pggen.arg('offset');
