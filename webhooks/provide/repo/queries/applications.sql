

-- CreateApplications inserts applications into the database
-- name: CreateApplications :many
INSERT INTO webhooks.application (
    name,
    tenant_id,
    rate_limit,
    metadata
) 
SELECT 
    name,
    tenant_id,
    rate_limit,
    metadata
FROM unnest(pggen.arg('applications')::webhooks.new_application[])
ON CONFLICT(tenant_id, name) DO NOTHING
RETURNING (
    id,
    name,
    uid,
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
SELECT ((
    id,
    name,
    uid,
    tenant_id,
    rate_limit,
    metadata,
    created_at,
    updated_at
)::webhooks.application, 
(SELECT ARRAY_AGG((
    id,
    url,
    name,
    application_id,
    uid,
    rate_limit,
    metadata,
    disabled,
    description,
    created_at,
    updated_at
)::webhooks.endpoint) FROM webhooks.endpoint e WHERE e.application_id = id)
)::webhooks.application_details
FROM webhooks.application
WHERE uid = ANY(pggen.arg('ids')::uuid[]);

-- GetApplicationsByName gets applications by name
-- name: GetApplicationsByName :many
SELECT (
    id,
    name,
    uid,
    tenant_id,
    rate_limit,
    metadata,
    created_at,
    updated_at
)::webhooks.application
FROM webhooks.application
WHERE name = ANY(pggen.arg('names'));

-- ListApplications lists registered applications
-- name: ListApplications :many
SELECT (
    id,
    name,
    uid,
    tenant_id,
    rate_limit,
    metadata,
    created_at,
    updated_at
)::webhooks.application
FROM webhooks.application
WHERE created_at > pggen.arg('created_after') 
AND (pggen.arg('tenant_id') = '' OR tenant_id = pggen.arg('tenant_id'))
ORDER BY created_at
LIMIT pggen.arg('limit')
OFFSET pggen.arg('offset');
