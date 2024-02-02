// Code generated by pggen. DO NOT EDIT.

package queries

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"time"
)

const createEndpointsSQL = `WITH inserted AS (
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
    FROM unnest($1::webhooks.new_endpoint[]) u
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
), linked_secrets AS (
    INSERT INTO webhooks.endpoint_secret (
        secret_id,
        endpoint_id
    )
    SELECT 
        s.id,
        i.id
    FROM unnest($1::webhooks.new_endpoint[]) n,
        webhooks.secret s,
        inserted i 
    WHERE s.uid = n.secret_id AND i.url = n.url
) , inserted_filters AS (
    INSERT INTO webhooks.endpoint_filter (
        event_type_id,
        endpoint_id
    )
    SELECT 
        e.id,
        i.id
    FROM unnest($1::webhooks.new_endpoint[]) n, 
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
    u.filter_type_ids,
    '' --secret
)::webhooks.endpoint_details 
FROM inserted i
INNER JOIN unnest($1::webhooks.new_endpoint[]) u
    ON i.url = u.url;`

// CreateEndpoints implements Querier.CreateEndpoints.
func (q *DBQuerier) CreateEndpoints(ctx context.Context, endpoints []NewEndpoint) ([]EndpointDetails, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "CreateEndpoints")
	rows, err := q.conn.Query(ctx, createEndpointsSQL, q.types.newNewEndpointArrayInit(endpoints))
	if err != nil {
		return nil, fmt.Errorf("query CreateEndpoints: %w", err)
	}
	defer rows.Close()
	items := []EndpointDetails{}
	rowRow := q.types.newEndpointDetails()
	for rows.Next() {
		var item EndpointDetails
		if err := rows.Scan(rowRow); err != nil {
			return nil, fmt.Errorf("scan CreateEndpoints row: %w", err)
		}
		if err := rowRow.AssignTo(&item); err != nil {
			return nil, fmt.Errorf("assign CreateEndpoints row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close CreateEndpoints rows: %w", err)
	}
	return items, err
}

const deleteEndpointsSQL = `DELETE FROM webhooks.endpoint WHERE uid = ANY($1::UUID[]);`

// DeleteEndpoints implements Querier.DeleteEndpoints.
func (q *DBQuerier) DeleteEndpoints(ctx context.Context, ids []string) (pgconn.CommandTag, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "DeleteEndpoints")
	cmdTag, err := q.conn.Exec(ctx, deleteEndpointsSQL, ids)
	if err != nil {
		return cmdTag, fmt.Errorf("exec query DeleteEndpoints: %w", err)
	}
	return cmdTag, err
}

const getEndpointsSQL = `SELECT ((
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
    WHERE f.endpoint_id = id),
    (SELECT value FROM webhooks.secret s, webhooks.endpoint_secret es
    WHERE es.endpoint_id = id AND s.id = es.secret_id 
    LIMIT 1)
)::webhooks.endpoint_details FROM webhooks.endpoint
WHERE uid = ANY($1::uuid[]);`

// GetEndpoints implements Querier.GetEndpoints.
func (q *DBQuerier) GetEndpoints(ctx context.Context, ids []string) ([]EndpointDetails, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "GetEndpoints")
	rows, err := q.conn.Query(ctx, getEndpointsSQL, ids)
	if err != nil {
		return nil, fmt.Errorf("query GetEndpoints: %w", err)
	}
	defer rows.Close()
	items := []EndpointDetails{}
	rowRow := q.types.newEndpointDetails()
	for rows.Next() {
		var item EndpointDetails
		if err := rows.Scan(rowRow); err != nil {
			return nil, fmt.Errorf("scan GetEndpoints row: %w", err)
		}
		if err := rowRow.AssignTo(&item); err != nil {
			return nil, fmt.Errorf("assign GetEndpoints row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close GetEndpoints rows: %w", err)
	}
	return items, err
}

const getEndpointsByUrlSQL = `SELECT ((
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
    WHERE f.endpoint_id = id),
    (SELECT value FROM webhooks.secret s, webhooks.endpoint_secret es
    WHERE es.endpoint_id = id AND s.id = es.secret_id 
    LIMIT 1)
)::webhooks.endpoint_details FROM webhooks.endpoint
WHERE url = ANY($1);`

// GetEndpointsByUrl implements Querier.GetEndpointsByUrl.
func (q *DBQuerier) GetEndpointsByUrl(ctx context.Context, urls []string) ([]EndpointDetails, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "GetEndpointsByUrl")
	rows, err := q.conn.Query(ctx, getEndpointsByUrlSQL, urls)
	if err != nil {
		return nil, fmt.Errorf("query GetEndpointsByUrl: %w", err)
	}
	defer rows.Close()
	items := []EndpointDetails{}
	rowRow := q.types.newEndpointDetails()
	for rows.Next() {
		var item EndpointDetails
		if err := rows.Scan(rowRow); err != nil {
			return nil, fmt.Errorf("scan GetEndpointsByUrl row: %w", err)
		}
		if err := rowRow.AssignTo(&item); err != nil {
			return nil, fmt.Errorf("assign GetEndpointsByUrl row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close GetEndpointsByUrl rows: %w", err)
	}
	return items, err
}

const listEndpointsSQL = `SELECT (
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
WHERE created_at > $1  
ORDER BY uid
LIMIT $2
OFFSET $3;`

type ListEndpointsParams struct {
	CreatedAfter time.Time `json:"created_after"`
	Limit        int       `json:"limit"`
	Offset       int       `json:"offset"`
}

// ListEndpoints implements Querier.ListEndpoints.
func (q *DBQuerier) ListEndpoints(ctx context.Context, params ListEndpointsParams) ([]Endpoint, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "ListEndpoints")
	rows, err := q.conn.Query(ctx, listEndpointsSQL, params.CreatedAfter, params.Limit, params.Offset)
	if err != nil {
		return nil, fmt.Errorf("query ListEndpoints: %w", err)
	}
	defer rows.Close()
	items := []Endpoint{}
	rowRow := q.types.newEndpoint()
	for rows.Next() {
		var item Endpoint
		if err := rows.Scan(rowRow); err != nil {
			return nil, fmt.Errorf("scan ListEndpoints row: %w", err)
		}
		if err := rowRow.AssignTo(&item); err != nil {
			return nil, fmt.Errorf("assign ListEndpoints row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close ListEndpoints rows: %w", err)
	}
	return items, err
}

const listApplicationEndpointsSQL = `SELECT (
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
WHERE e.created_at > $1 
AND webhooks.application.uid = $2::uuid
ORDER BY e.uid
LIMIT $3
OFFSET $4;`

type ListApplicationEndpointsParams struct {
	CreatedAfter   time.Time `json:"created_after"`
	ApplicationUid string    `json:"application_uid"`
	Limit          int       `json:"limit"`
	Offset         int       `json:"offset"`
}

// ListApplicationEndpoints implements Querier.ListApplicationEndpoints.
func (q *DBQuerier) ListApplicationEndpoints(ctx context.Context, params ListApplicationEndpointsParams) ([]Endpoint, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "ListApplicationEndpoints")
	rows, err := q.conn.Query(ctx, listApplicationEndpointsSQL, params.CreatedAfter, params.ApplicationUid, params.Limit, params.Offset)
	if err != nil {
		return nil, fmt.Errorf("query ListApplicationEndpoints: %w", err)
	}
	defer rows.Close()
	items := []Endpoint{}
	rowRow := q.types.newEndpoint()
	for rows.Next() {
		var item Endpoint
		if err := rows.Scan(rowRow); err != nil {
			return nil, fmt.Errorf("scan ListApplicationEndpoints row: %w", err)
		}
		if err := rowRow.AssignTo(&item); err != nil {
			return nil, fmt.Errorf("assign ListApplicationEndpoints row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close ListApplicationEndpoints rows: %w", err)
	}
	return items, err
}
