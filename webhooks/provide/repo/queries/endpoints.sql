-- CreateEndpoints inserts endpoints into the database
-- name: CreateEndpoints :many
WITH inserted AS (
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
    RETURNING
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
), inserted_filters AS (
    INSERT INTO webhooks.endpoint_filter (
        event_type_id,
        endpoint_id
    )
    SELECT 
        e.id,
        i.id
    FROM unnest(pggen.arg('endpoints')::webhooks.new_endpoint[]) n, 
        unnest(filter_type_ids::UUID[]) f,
        webhooks.event_type e,
        inserted i 
    WHERE e.uid = f AND i.url = n.url
) SELECT ((
        i.id,
        i.url,
        i.name,
        i.application_id,
        i.uid,
        i.rate_limit,
        i.metadata,
        i.disabled,
        i.description,
        i.created_at,
        i.updated_at
    )::webhooks.endpoint,
    u.filter_type_ids
)::webhooks.endpoint_details 
FROM inserted i
INNER JOIN unnest(pggen.arg('endpoints')::webhooks.new_endpoint[]) u
    ON i.url = u.url;


-- DeleteEndpoints deletes endpoints by uid
-- name: DeleteEndpoints :exec
DELETE FROM webhooks.endpoint WHERE uid = ANY(pggen.arg('ids')::UUID[]);

-- GetEndpoints gets endpoints by id
-- name: GetEndpoints :many
SELECT ((
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
    )::webhooks.endpoint,
    (SELECT ARRAY_AGG(e.uid::UUID)
    FROM webhooks.event_type e
    INNER JOIN webhooks.endpoint_filter f 
        ON f.event_type_id = e.id  
    WHERE f.endpoint_id = id)
)::webhooks.endpoint_details FROM webhooks.endpoint
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
WHERE created_at > pggen.arg('created_after')  
ORDER BY uid
LIMIT pggen.arg('limit')
OFFSET pggen.arg('offset');

-- ListApplicationEndpoints lists endpoints
-- name: ListApplicationEndpoints :many
SELECT (
    e.id,
    e.url,
    e.name,
    e.application_id,
    e.uid,
    e.rate_limit,
    e.metadata,
    e.disabled,
    e.description,
    e.created_at,
    e.updated_at
)::webhooks.endpoint
FROM webhooks.endpoint e
JOIN webhooks.application ON e.application_id = webhooks.application.id
WHERE e.created_at > pggen.arg('created_after') 
AND webhooks.application.uid = pggen.arg('application_uid')::uuid
ORDER BY e.uid
LIMIT pggen.arg('limit')
OFFSET pggen.arg('offset');