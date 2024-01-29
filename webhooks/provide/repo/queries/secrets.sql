
-- ListSecrets lists secrets
-- name: ListMesListSecretssages :many
SELECT
    id,
    uid
FROM webhooks.secret
ORDER BY uid
LIMIT pggen.arg('limit')
OFFSET pggen.arg('offset');
