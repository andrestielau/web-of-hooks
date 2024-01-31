
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
RETURNING (
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
)::webhooks.endpoint;

-- DeleteEndpoints deletes endpoints by uid
-- name: DeleteEndpoints :exec
DELETE FROM webhooks.endpoint WHERE uid = ANY(pggen.arg('ids')::UUID[]);

-- GetEndpoints gets endpoints by id
-- name: GetEndpoints :many
SELECT (
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
)::webhooks.endpoint
FROM webhooks.endpoint
WHERE uid = ANY(pggen.arg('ids')::uuid[]);

-- ListEndpoints lists endpoints
-- name: ListEndpoints :many
SELECT (
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
)::webhooks.endpoint
FROM webhooks.endpoint
WHERE uid > pggen.arg('after') 
ORDER BY uid
LIMIT pggen.arg('limit')
OFFSET pggen.arg('offset');
