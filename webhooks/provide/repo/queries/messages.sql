
-- DeleteMessages deletes messages by uid
-- name: DeleteMessages :exec
DELETE FROM webhooks.message WHERE uid = ANY(pggen.arg('ids')::UUID[]);

-- ListMessages lists event-types
-- name: ListMessages :many
SELECT
    id,
    uid,
    created_at
FROM webhooks.message
ORDER BY uid
LIMIT pggen.arg('limit')
OFFSET pggen.arg('offset');
