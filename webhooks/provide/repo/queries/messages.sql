
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
