

-- DeleteAttempts deletes attempts by uid
-- name: DeleteAttempts :exec
DELETE FROM webhooks.message_attempt WHERE uid = ANY(pggen.arg('ids')::UUID[]);


-- GetAttempts gets attempts by id
-- name: GetAttempts :many
SELECT 
    id,
    uid,
    created_at
FROM webhooks.message_attempt
WHERE uid = ANY(pggen.arg('ids')::uuid[]);

-- ListAttempts lists message attempts
-- name: ListAttempts :many
SELECT
    id,
    uid,
    created_at
FROM webhooks.message_attempt
ORDER BY uid
LIMIT pggen.arg('limit')
OFFSET pggen.arg('offset');
