
-- CreateMessages inserts messages into the database
-- name: CreateMessages :many
WITH new_messages AS (
    INSERT INTO webhooks.message (
        application_id,
        event_type_id,
        event_id,
        payload
    ) SELECT 
        a.id,
        e.id,
        u.event_id,
        u.payload
    FROM unnest(pggen.arg('messages')::webhooks.new_message[]) u
    JOIN webhooks.application a ON u.application_id = a.uid
    JOIN webhooks.event_type e ON u.event_type_id = e.key
    ON CONFLICT DO NOTHING
    RETURNING 
        id,
        application_id,
        event_type_id,
        uid,
        created_at,
        event_id,
        payload
), new_attempts AS (
    INSERT INTO webhooks.message_attempt (
        uid,
        endpoint_id,
        message_id
    ) SELECT
        generate_ulid(),
        e.id,
        n.id
    FROM new_messages n
    INNER JOIN webhooks.endpoint_filter f 
        ON f.event_type_id = n.event_type_id
    INNER JOIN webhooks.endpoint e 
        ON e.id = f.endpoint_id 
            AND e.application_id = n.application_id
    RETURNING 
        id,
        uid, 
        endpoint_id,
        message_id,
        created_at,
        updated_at,
        retry,
        status,
        response_status,
        response
) SELECT ((
        id,
        application_id,
        event_type_id,
        uid,
        created_at,
        event_id,
        payload
    )::webhooks.message,
    (SELECT ARRAY_AGG(a::webhooks.message_attempt) 
    FROM new_attempts a 
    WHERE id = a.message_id)
)::webhooks.message_details FROM new_messages;

-- DeleteMessages deletes messages by uid
-- name: DeleteMessages :exec
DELETE FROM webhooks.message WHERE uid = ANY(pggen.arg('ids')::UUID[]);

-- GetMessages gets messages by id
-- name: GetMessages :many
SELECT ((
        id,
        application_id,
        event_type_id,
        uid,
        created_at,
        event_id,
        payload
    )::webhooks.message,
    (SELECT ARRAY_AGG(a::webhooks.message_attempt) FROM webhooks.message_attempt a WHERE id = a.message_id)
)::webhooks.message_details FROM webhooks.message
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
WHERE created_at > pggen.arg('created_after') 
ORDER BY uid
LIMIT pggen.arg('limit')
OFFSET pggen.arg('offset');

-- ListApplicationMessages lists event-types
-- name: ListApplicationMessages :many
SELECT (
    m.id ,
    m.application_id ,
    m.event_type_id ,
    m.uid ,
    m.created_at,
    m.event_id,
    m.payload
)::webhooks.message
FROM webhooks.message m
JOIN webhooks.application ON m.application_id = webhooks.application.id
WHERE m.created_at > pggen.arg('created_after') 
AND webhooks.application.uid = pggen.arg('application_uid')::uuid
ORDER BY m.uid
LIMIT pggen.arg('limit')
OFFSET pggen.arg('offset');
