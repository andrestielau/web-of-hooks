

-- DeleteAttempts deletes attempts by uid
-- name: DeleteAttempts :exec
DELETE FROM webhooks.message_attempt WHERE uid = ANY(pggen.arg('ids')::UUID[]);


-- GetAttempts gets attempts by id
-- name: GetAttempts :many
SELECT (
    id,
    uid,
    endpoint_id ,
    message_id ,
    created_at,
    status,
    response_status,
    response
)::webhooks.message_attempt
FROM webhooks.message_attempt
WHERE uid = ANY(pggen.arg('ids')::uuid[]);

-- ListAttempts lists message attempts
-- name: ListAttempts :many
SELECT (
    id,
    uid,
    endpoint_id ,
    message_id ,
    created_at,
    status,
    response_status,
    response
)::webhooks.message_attempt
FROM webhooks.message_attempt
WHERE uid > pggen.arg('after') 
ORDER BY uid
LIMIT pggen.arg('limit')
OFFSET pggen.arg('offset');
