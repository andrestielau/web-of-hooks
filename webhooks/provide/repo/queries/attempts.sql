

-- DeleteAttempts deletes attempts by uid
-- name: DeleteAttempts :exec
DELETE FROM webhooks.attempt WHERE uid = ANY(pggen.arg('ids')::UUID[]);

-- ListMessages lists event-types
-- name: ListMessages :many
SELECT
    id,
    uid,
    created_at
FROM webhooks.message_attempt
ORDER BY uid
LIMIT pggen.arg('limit')
OFFSET pggen.arg('offset');
