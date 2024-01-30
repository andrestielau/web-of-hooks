

-- CreateApplications inserts applications into the database
-- name: CreateApplications :many
INSERT INTO webhooks.application (
    name,
    tenant_id,
    rate_limit,
    metadata
) 
SELECT 
    u.name,
    u.tenant_id,
    u.rate_limit,
    u.metadata
FROM unnest(pggen.arg('applications')::webhooks.new_application[]) u
ON CONFLICT DO NOTHING
RETURNING (
    id,
    uid,
    name,
    tenant_id,
    rate_limit,
    metadata,
    created_at,
    updated_at
)::webhooks.application;

-- DeleteApplications deletes application by uid
-- name: DeleteApplications :exec
DELETE FROM webhooks.application WHERE uid = ANY(pggen.arg('ids')::UUID[]);

-- GetApplications gets applications by id
-- name: GetApplications :many
SELECT (
    id,
    uid,
    name,
    tenant_id,
    rate_limit,
    metadata,
    created_at,
    updated_at
)::webhooks.application
FROM webhooks.application
WHERE uid = ANY(pggen.arg('ids')::uuid[]);

-- ListApplications lists registered applications
-- name: ListApplications :many
SELECT (
    id,
    uid,
    name,
    tenant_id,
    rate_limit,
    metadata,
    created_at,
    updated_at
)::webhooks.application
FROM webhooks.application
WHERE created_at > pggen.arg('created_after') 
ORDER BY created_at
LIMIT pggen.arg('limit')
OFFSET pggen.arg('offset');
