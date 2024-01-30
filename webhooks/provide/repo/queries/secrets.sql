-- CreateSecrets creates secrets
-- name: CreateSecrets :many
 INSERT INTO webhooks.secret (
    uid,
    tenant_id
) 
SELECT 
    u.id,
    u.tenant_id
FROM unnest(pggen.arg('secrets')::webhooks.new_secret[]) u
ON CONFLICT DO NOTHING
RETURNING 
    id,
    uid,
    tenant_id;

-- GetSecrets gets secrets by id
-- name: GetSecrets :many
SELECT 
    id,
    uid
FROM webhooks.secret
WHERE uid = ANY(pggen.arg('ids')::uuid[]);

-- DeleteSecrets deletes secrets by uid
-- name: DeleteSecrets :exec
DELETE FROM webhooks.secret WHERE uid = ANY(pggen.arg('ids')::UUID[]);

-- ListSecrets lists secrets
-- name: ListSecrets :many
SELECT
    id,
    uid
FROM webhooks.secret
ORDER BY uid
LIMIT pggen.arg('limit')
OFFSET pggen.arg('offset');
