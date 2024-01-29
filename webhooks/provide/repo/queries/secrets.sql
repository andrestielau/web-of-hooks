

-- DeleteSecrets deletes secrets by uid
-- name: DeleteSecrets :exec
DELETE FROM webhooks.secret WHERE uid = ANY(pggen.arg('ids')::UUID[]);

-- ListSecrets lists secrets
-- name: ListMesListSecretssages :many
SELECT
    id,
    uid
FROM webhooks.secret
ORDER BY uid
LIMIT pggen.arg('limit')
OFFSET pggen.arg('offset');
