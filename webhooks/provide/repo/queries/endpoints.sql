


-- CreateEndpoints inserts endpoints into the database
-- name: CreateEndpoints :many
 INSERT INTO webhooks.endpoint (
    application_id,
    url,
    name,
    rate_limit,
    metadata,
    description
) 
SELECT 
    a.id,
    u.url,
    u.name,
    u.rate_limit,
    u.metadata,
    u.description
FROM unnest(pggen.arg('endpoints')::webhooks.new_endpoint[]) u
JOIN webhooks.application a ON u.application_id = a.uid
ON CONFLICT DO NOTHING
RETURNING 
    id,
    uid,
        application_id,
    rate_limit,
    metadata,
    created_at;

-- ListEndpoints lists endpoints
-- name: ListEndpoints :many
SELECT 
    id,
    uid,
    url,
    name,
    metadata,
    disabled,
    rate_limit,
    created_at,
    updated_at,
    description,
    application_id
FROM webhooks.endpoint
ORDER BY uid
LIMIT pggen.arg('limit')
OFFSET pggen.arg('offset');
