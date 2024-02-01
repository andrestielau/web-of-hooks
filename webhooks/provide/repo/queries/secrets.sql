-- CreateSecrets creates secrets
-- name: CreateSecrets :many
 INSERT INTO webhooks.secret (
    value,
    application_id
) 
SELECT 
    u.value,
    a.id
FROM unnest(pggen.arg('secrets')::webhooks.new_secret[]) u
 JOIN webhooks.application a ON u.application_id = a.uid
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
WHERE created_at > pggen.arg('created_after') 
ORDER BY uid
LIMIT pggen.arg('limit')
OFFSET pggen.arg('offset');

-- Update secret value by uid
-- name: UpdateSecret :exec
UPDATE webhooks.secret SET value = pggen.arg('value') WHERE uid = pggen.arg('uid');

-- ListApplicationSecrets lists secrets of an application
-- name: ListApplicationSecrets :many
SELECT (
    s.id,
    s.uid,
    s.application_id,
    s.value,
    s.created_at,
    s.updated_at
)::webhooks.secret
FROM webhooks.secret s
JOIN webhooks.application ON s.application_id = webhooks.application.id
WHERE webhooks.application.uid = pggen.arg('application_uid')::uuid
ORDER BY s.uid;

