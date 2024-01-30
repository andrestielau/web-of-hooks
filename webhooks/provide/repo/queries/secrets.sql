-- CreateSecrets creates secrets
-- name: CreateSecrets :many
 INSERT INTO webhooks.secret (
    value,
    application_id
) 
SELECT 
    u.value,
    u.application_id
FROM unnest(pggen.arg('secrets')::webhooks.new_secret[]) u
ON CONFLICT DO NOTHING
RETURNING (
    id,
    uid,
    application_id,
    value,
    created_at,
    updated_at
)::webhooks.secret;

-- GetSecrets gets secrets by id
-- name: GetSecrets :many
SELECT (
    id,
    uid,
    application_id,
    value,
    created_at,
    updated_at
)::webhooks.secret
FROM webhooks.secret
WHERE uid = ANY(pggen.arg('ids')::uuid[]);

-- DeleteSecrets deletes secrets by uid
-- name: DeleteSecrets :exec
DELETE FROM webhooks.secret WHERE uid = ANY(pggen.arg('ids')::UUID[]);

-- ListSecrets lists secrets
-- name: ListSecrets :many
SELECT (
    id,
    uid,
    application_id,
    value,
    created_at,
    updated_at
)::webhooks.secret
FROM webhooks.secret
WHERE uid > pggen.arg('after')
ORDER BY uid
LIMIT pggen.arg('limit')
OFFSET pggen.arg('offset');

-- Update secret value by uid
-- name: UpdateSecret :exec
UPDATE webhooks.secret SET value = pggen.arg('value') WHERE uid = pggen.arg('uid');

