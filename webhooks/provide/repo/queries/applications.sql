

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