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
RETURNING (
    uid,
    id ,
    tenant_id,
    value
)::webhooks.secret;

-- GetSecrets gets secrets by id
-- name: GetSecrets :many
SELECT (
    uid,
    id ,
    tenant_id,
    value
)::webhooks.secret
FROM webhooks.secret
WHERE uid = ANY(pggen.arg('ids')::uuid[]);

-- DeleteSecrets deletes secrets by uid
-- name: DeleteSecrets :exec
DELETE FROM webhooks.secret WHERE uid = ANY(pggen.arg('ids')::UUID[]);

-- ListSecrets lists secrets
-- name: ListSecrets :many
SELECT (
    uid,
    id ,
    tenant_id,
    value
)::webhooks.secret
FROM webhooks.secret
WHERE uid > pggen.arg('after')
ORDER BY uid
LIMIT pggen.arg('limit')
OFFSET pggen.arg('offset');
