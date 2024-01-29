

-- CreateApplications inserts applications into the database
-- name: CreateApplications :many
INSERT INTO webhooks.application (
    tenant_id,
    rate_limit,
    metadata
) 
SELECT 
    u.tenant_id,
    u.rate_limit,
    u.metadata
FROM unnest(pggen.arg('applications')::webhooks.new_application[]) u
ON CONFLICT DO NOTHING
RETURNING 
    id,
    uid,
    tenant_id,
    rate_limit,
    metadata,
    created_at;

-- DeleteApplications deletes application by uid
-- name: DeleteApplications :exec
DELETE FROM webhooks.application WHERE uid = ANY(pggen.arg('ids')::UUID[]);

-- ListApplications lists registered applications
-- name: ListApplications :many
SELECT
    id,
    uid,
    tenant_id,
    rate_limit,
    metadata,
    created_at,
    updated_at
FROM webhooks.application
ORDER BY uid
LIMIT pggen.arg('limit')
OFFSET pggen.arg('offset');
